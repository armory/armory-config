// Package migrate is to create and insert the data inside kustomize files
package migrate

import (
	"fmt"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/migrate/armory_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/config_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/files_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/profile_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/providers_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/service_settings_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/sizing_patch"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.AppEngine.Enable
var writeConfigFiles bool

func CreateKustomization(KustomizeData *structs.Kustomize, writefiles bool) error {
	// KustomizeData := structs.Kustomize{}

	writeConfigFiles = writefiles
	SetHeaders(KustomizeData)
	str := GetKustomization(KustomizeData)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/kustomization.yml", str)
	if err != nil {
		return err
	}

	KustomizeData.CurrentDeploymentPos = GetCurrentDeploymentPosition(*KustomizeData)

	//This should never ocurre because of the validation validateCurrentDeploymentExists
	if -1 == KustomizeData.CurrentDeploymentPos {
		return fmt.Errorf("Error while getting current deployment position")
	}

	// Slice of function calls to generate all the kustomize files
	functionCalls := []func(structs.Kustomize) error{
		CreateSpinnakerService,
		CreateConfigPatch,
		CreateProfilesPatch,
		CreateFilesPatch,
		CreateServiceSettingsPatch,
		CreateConfigProviders,
	}

	if "ARMORY" == KustomizeData.Spin_flavor {
		functionCalls = append(functionCalls, CreateArmoryPatch)
	}

	if !sizing_patch.IsDeploymentEnvironmentEmpty(*KustomizeData) {
		functionCalls = append(functionCalls, CreatePatchSizing)
	}

	for _, function := range functionCalls {
		err = function(*KustomizeData)
		if err != nil {
			return err
			// return fmt.Errorf("Error while executing the function:%s Error:(%s)", function, err)
		}
	}

	return nil
}

func GetKustomization(KustomizeData *structs.Kustomize) string {
	str := ""

	if nil != KustomizeData {
		str = `apiVersion: kustomize.config.k8s.io/v1beta1
	kind: Kustomization

	namespace: spinnaker` + /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].deploymentEnvironment.Location +*/ `

	resources:
	- SpinnakerService.yml

	# Apply the patches top down order
	patchesStrategicMerge:
	- config-patch.yml              #Contains Spinnaker configuration (contents of a deployment)
	- profiles-patch.yml            #Contains the YAML of each service's profile
	- files-patch.yml               #Contains any other raw string files not handle in spinnakerConfig
	- service-settings-patch.yml    #Contains the config for each service's service-setting
	- config-providers-patch.yml    #Contains the providers configuration
`

		if "ARMORY" == KustomizeData.Spin_flavor {
			str += `	- armory-patch.yml              #Contains Specific Armory config
`
		}

		if !sizing_patch.IsDeploymentEnvironmentEmpty(*KustomizeData) {
			str += `	- patch-sizing.yml              #Contains Halyard DeploymentEnvironment
`
		}

		//TODO Missing catch of the error
		_, mergesFileNotEmpty := CreatePatchMerges(*KustomizeData)
		if mergesFileNotEmpty {
			str += `	- patch-merges.yml              #Contains Halyard initContainers and sidecars
`
		}

		str = strings.Replace(str, "\t", "", -1)
	}

	return str
}

func SetHeaders(KustomizeData *structs.Kustomize) {
	apiVersion := ""

	if "OSS" == KustomizeData.Spin_flavor {
		apiVersion = "apiVersion: spinnaker.io/v1alpha2"
	} else {
		apiVersion = "apiVersion: spinnaker.armory.io/v1alpha2"
	}

	KustomizeData.Header = apiVersion + `
kind: SpinnakerService
metadata:
  name: spinnaker
spec:`

}

func GetCurrentDeploymentPosition(KustomizeData structs.Kustomize) int {

	for i, a := range KustomizeData.Halyard.GetDeploymentConfigurations() {
		if KustomizeData.Halyard.GetCurrentDeployment() == a.Name {
			return i
		}
	}

	return -1
}

func CreateSpinnakerService(KustomizeData structs.Kustomize) error {
	SpinnakerServiceStr := GetSpinnakerServicePatch(KustomizeData)

	if "" != SpinnakerServiceStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/SpinnakerService.yml", SpinnakerServiceStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateConfigPatch(KustomizeData structs.Kustomize) error {
	ConfigPatchStr := GetConfigPatch(KustomizeData)

	if "" != ConfigPatchStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/config-patch.yml", ConfigPatchStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateProfilesPatch(KustomizeData structs.Kustomize) error {
	ProfilesPatchStr := GetProfilesPatch(KustomizeData)

	if "" != ProfilesPatchStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/profiles-patch.yml", ProfilesPatchStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFilesPatch(KustomizeData structs.Kustomize) error {
	FilesPatchStr := GetFilesPatch(KustomizeData)

	if "" != FilesPatchStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/files-patch.yml", FilesPatchStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateServiceSettingsPatch(KustomizeData structs.Kustomize) error {
	ServiceSettingsPatchStr := GetServiceSettingsPatch(KustomizeData)

	if "" != ServiceSettingsPatchStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/service-settings-patch.yml", ServiceSettingsPatchStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateConfigProviders(KustomizeData structs.Kustomize) error {
	ConfigProvidersStr := GetConfigProvidersPatch(KustomizeData)

	if "" != ConfigProvidersStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/config-providers-patch.yml", ConfigProvidersStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreatePatchSizing(KustomizeData structs.Kustomize) error {
	PatchSizingStr := GetPatchSizing(KustomizeData)

	if "" != PatchSizingStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/patch-sizing.yml", PatchSizingStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateArmoryPatch(KustomizeData structs.Kustomize) error {
	ArmoryPatchStr := GetArmoryPatch(KustomizeData)

	if "" != ArmoryPatchStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/armory-patch.yml", ArmoryPatchStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetSpinnakerServicePatch(KustomizeData structs.Kustomize) string {

	str := KustomizeData.Header + `
  spinnakerConfig:
    config:
      version: 0.0.0   # the version of Spinnaker to be deployed, see config-patch.yml
      persistentStorage:
        persistentStoreType: s3
        s3:
          bucket: mybucket
          rootFolder: front50
    profiles:
      clouddriver:
      artifacts:
        http:
        enabled: true
        accounts: []
      deck:
        settings-local.js: |
          window.spinnakerSettings.feature.kustomizeEnabled = true;
          window.spinnakerSettings.feature.artifactsRewrite = true;
      echo: {}    # is the contents of ~/.hal/default/profiles/echo.yml
      fiat: {}    # is the contents of ~/.hal/default/profiles/fiat.yml
      front50: {} # is the contents of ~/.hal/default/profiles/front50.yml
      gate: {}    # is the contents of ~/.hal/default/profiles/gate.yml
      igor: {}    # is the contents of ~/.hal/default/profiles/igor.yml
      kayenta: {} # is the contents of ~/.hal/default/profiles/kayenta.yml
      orca: {}    # is the contents of ~/.hal/default/profiles/orca.yml
      rosco: {}   # is the contents of ~/.hal/default/profiles/rosco.yml

    service-settings:
      clouddriver: {}
      deck: {}
      echo: {}
      fiat: {}
      front50: {}
      gate: {}
      igor: {}
      kayenta: {}
      orca: {}
      rosco: {}
    files: {}

  expose:
    service:
      overrides: {}
      type: LoadBalancer
    type: service
`

	return str
}

func GetConfigPatch(KustomizeData structs.Kustomize) string {

	str := KustomizeData.Header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.config - This section contains the contents of a deployment found in a halconfig .deploymentConfigurations[0]
    config:` +
		config_patch.GetConfigData(KustomizeData) + `
`

	return str
}

func GetProfilesPatch(KustomizeData structs.Kustomize) string {

	str := KustomizeData.Header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.profiles - This section contains the YAML of each service's profile
    profiles:` +
		profile_patch.GetProfiles(KustomizeData) + `
`

	return str
}

func GetFilesPatch(KustomizeData structs.Kustomize) string {

	str := KustomizeData.Header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.files - This section allows you to include any other raw string files not handle above.
    # The KEY is the filepath and filename of where it should be placed
    #  - Files here will be placed into ~/.hal/default/ on halyard
    #  - __ is used in place of / for the path separator
    # The VALUE is the contents of the file.
    #  - Use the | YAML symbol to indicate a block-style multiline string
    #  - We currently only support string files
    #  - NOTE: Kubernetes has a manifest size limitation of 1MB
    files:`

	s := files_patch.GetFiles(KustomizeData)

	if s != `` {
		str += s

	} else {
		str += ` {}`
	}
	if writeConfigFiles {
		str += files_patch.WriteConfigFiles(KustomizeData)
	}

	return str
}

func GetServiceSettingsPatch(KustomizeData structs.Kustomize) string {

	str := KustomizeData.Header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.service-settings - This section contains the YAML of the service's service-setting
    # see https://www.spinnaker.io/reference/halyard/custom/#tweakable-service-settings for available settings
    service-settings:` +
		service_settings_patch.GetServiceSettings(KustomizeData) + `
`

	return str
}

func GetConfigProvidersPatch(KustomizeData structs.Kustomize) string {
	str := KustomizeData.Header +
		providers_patch.GetProvidersData(KustomizeData) + `
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}

func GetPatchSizing(KustomizeData structs.Kustomize) string {
	str := KustomizeData.Header + `
  spinnakerConfig:
    config:
      deploymentEnvironment:` +
		sizing_patch.GetDeploymentEnvironment(KustomizeData) + `
`

	return str
}

func GetArmoryPatch(KustomizeData structs.Kustomize) string {
	str := KustomizeData.Header + `
  spinnakerConfig:
    config:
      armory:` +
		armory_patch.GetArmory(KustomizeData) + `
`

	return str
}

func GetPatchMerges(KustomizeData structs.Kustomize) string {
	str := KustomizeData.Header + `
  spinnakerConfig:
    profiles: {}
  kustomize:` +
		sizing_patch.GetDeploymentEnvInitContainers(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].DeploymentEnvironment.InitContainers) + `
`

	if strings.Contains(str, "--|||EMPTY|||--") {
		str = ""
	}

	return str
}

// Returns true when patch-merges is Not Empty
func CreatePatchMerges(KustomizeData structs.Kustomize) (error, bool) {
	isNotEmpty := false
	PatchMergesStr := GetPatchMerges(KustomizeData)

	if "" != PatchMergesStr {
		err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/patch-merges.yml", PatchMergesStr)
		if err != nil {
			return err, isNotEmpty
		}
		isNotEmpty = true
	}
	return nil, isNotEmpty
}
