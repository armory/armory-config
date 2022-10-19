// Package migrate is to create and insert the data inside kustomize files
package migrate

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations"
	"github.com/austinthao5/golang_proto_test/internal/fileio"
)

type Kustomize struct {
	Spin_flavor          string
	Output_dir           string
	Halyard              *deploymentConfigurations.HalFile
	CurrentDeploymentPos int
	PatchSizing          bool
}

// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.AppEngine.Enable
func (KustomizeData Kustomize) CreateKustomization() error {
	str := `apiVersion: kustomize.config.k8s.io/v1beta1
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
	- config-providers.yml          #Contains the providers configuration
`

	if "ARMORY" == KustomizeData.Spin_flavor {
		str += `	- armory-patch.yml              #Contains Specific Armory config
`
	}

	//Todo Check if patch-sizing not empty
	KustomizeData.PatchSizing = true

	if KustomizeData.PatchSizing {
		str += `	- patch-sizing.yml              #Contains Halyard DeploymentEnvironment
`
	}

	str = strings.Replace(str, "\t", "", -1)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/kustomization.yml", str)
	if err != nil {
		return err
	}

	header := KustomizeData.GetHeaders()
	KustomizeData.CurrentDeploymentPos = KustomizeData.GetCurrentDeploymentPosition()

	//This should never ocurre because of the validation validateCurrentDeploymentExists
	// if -1 == KustomizeData.CurrentDeploymentPos {
	// 	return fmt.Errorf("Error while getting current deployment position")
	// }

	// Slice of function calls to generate all the kustomize files
	functionCalls := []func(string) error{
		KustomizeData.CreateSpinnakerService,
		KustomizeData.CreateConfigPatch,
		KustomizeData.CreateProfilesPatch,
		KustomizeData.CreateFilesPatch,
		KustomizeData.CreateServiceSettingsPatch,
		KustomizeData.CreateConfigProviders,
	}

	if "ARMORY" == KustomizeData.Spin_flavor {
		functionCalls = append(functionCalls, KustomizeData.CreateArmoryPatch)
	}

	if KustomizeData.PatchSizing {
		functionCalls = append(functionCalls, KustomizeData.CreatePatchSizing)
	}

	for _, function := range functionCalls {
		err = function(header)
		if err != nil {
			return err
			// return fmt.Errorf("Error while executing the function:%s Error:(%s)", function, err)
		}
	}

	return nil
}

func (KustomizeData Kustomize) GetHeaders() string {
	apiVersion := ""

	if "OSS" == KustomizeData.Spin_flavor {
		apiVersion = "apiVersion: spinnaker.io/v1alpha2"
	} else {
		apiVersion = "apiVersion: spinnaker.armory.io/v1alpha2"
	}

	header := apiVersion + `
kind: SpinnakerService
metadata:
  name: spinnaker
spec:`

	return header
}

func (KustomizeData Kustomize) GetCurrentDeploymentPosition() int {

	for i, a := range KustomizeData.Halyard.GetDeploymentConfiguration() {
		if KustomizeData.Halyard.GetCurrentDeployment() == a.Name {
			return i
		}
	}

	return -1
}

func (KustomizeData Kustomize) CreateSpinnakerService(header string) error {
	SpinnakerServiceStr := KustomizeData.GetSpinnakerService(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/SpinnakerService.yml", SpinnakerServiceStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateConfigPatch(header string) error {
	ConfigPatchStr := KustomizeData.GetConfigPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/config-patch.yml", ConfigPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateProfilesPatch(header string) error {
	ProfilesPatchStr := KustomizeData.GetProfilesPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/profiles-patch.yml", ProfilesPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateFilesPatch(header string) error {
	FilesPatchStr := KustomizeData.GetFilesPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/files-patch.yml", FilesPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateServiceSettingsPatch(header string) error {
	ServiceSettingsPatchStr := KustomizeData.GetServiceSettingsPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/service-settings-patch.yml", ServiceSettingsPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateConfigProviders(header string) error {
	ConfigProvidersStr := KustomizeData.GetConfigProviders(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/config-providers.yml", ConfigProvidersStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreatePatchSizing(header string) error {
	PatchSizingStr := KustomizeData.GetPatchSizing(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/patch-sizing.yml", PatchSizingStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) CreateArmoryPatch(header string) error {
	ArmoryPatchStr := KustomizeData.GetArmoryPatch(header)

	err := fileio.WriteConfigsTmp(KustomizeData.Output_dir+"/armory-patch.yml", ArmoryPatchStr)
	if err != nil {
		return err
	}
	return nil
}

func (KustomizeData Kustomize) GetSpinnakerService(header string) string {

	str := header + `
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
    type: service
    service:
      type: LoadBalancer
      overrides: {}
`

	return str
}

func (KustomizeData Kustomize) GetConfigPatch(header string) string {

	str := header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.config - This section contains the contents of a deployment found in a halconfig .deploymentConfigurations[0]
    config:

# === General Config ===
      version: ` + "2.0" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Version*/ + `  # the version of Spinnaker to be deployed
      timezone: ` + "America/Los_Angeles" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Timezone*/ + `

# === Persistent Storage ===
      persistentStorage:` +
		GetPersistentStorage(KustomizeData) + `

# === Metric Stores ===
      metricStores:` +
		GetMetricStores(KustomizeData) + `

# === Notifications ===
      notifications:` +
		GetNotifications(KustomizeData) + `

# === Ci ===
      ci:` +
		GetCi(KustomizeData) + `

# === Security ===
      security:` +
		GetSecurity(KustomizeData) + `

# === Artifacts ===
      artifacts:` +
		GetArtifacts(KustomizeData) + `

# === Pubsub ===
      pubsub:` +
		GetPubsub(KustomizeData) + `

# === Canary ===
      canary:` +
		GetCanary(KustomizeData) + `

# === Stats ===
      stats:` +
		GetStats(KustomizeData) + `
`

	return str
}

func (KustomizeData Kustomize) GetProfilesPatch(header string) string {

	str := header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.profiles - This section contains the YAML of each service's profile
    profiles:
      clouddriver: {} # is the contents of ~/.hal/default/profiles/clouddriver.yml
      # deck has a special key "settings-local.js" for the contents of settings-local.js
      deck:
        # settings-local.js - contents of ~/.hal/default/profiles/settings-local.js
        # Use the | YAML symbol to indicate a block-style multiline string
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
`

	return str
}

func (KustomizeData Kustomize) GetFilesPatch(header string) string {

	str := header + `
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
    files:
    #  profiles__rosco__packer__example-packer-config.json: |
    #    {
    #      "packerSetting": "someValue"
    #    }
    #  profiles__rosco__packer__my_custom_script.sh: |
    #    #!/bin/bash -e
    #    echo "hello world!"
`

	return str
}

func (KustomizeData Kustomize) GetServiceSettingsPatch(header string) string {

	str := header + `
  # spec.spinnakerConfig - This section is how to specify configuration spinnaker
  spinnakerConfig:
    # spec.spinnakerConfig.service-settings - This section contains the YAML of the service's service-setting
    # see https://www.spinnaker.io/reference/halyard/custom/#tweakable-service-settings for available settings
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
`

	return str
}

func (KustomizeData Kustomize) GetConfigProviders(header string) string {

	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Aws.GetAwsAcc()
	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.AppEngine

	prob := Providers{}
	prob.SetProvidersData(KustomizeData)

	str := header + `
  validation:
    providers:
      ` + prob.Enable + `:
        enabled: true    # Default: true. Indicate if operator should do connectivity checks to configured kubernetes accounts before applying the manifest
  spinnakerConfig:
    config:
      providers:
        appengine:
          ` + prob.AppEngine + `
        aws:
          ` + prob.Aws + `
        ecs:
          ` + prob.Ecs + `
        azure:
          ` + prob.Azure + `
        dcos:
          ` + prob.Dcos + `
        dockerRegistry:
          ` + prob.DockerRegistry + `
        google:
          ` + prob.Google + `
        huaweicloud:
          ` + prob.Huaweicloud + `
        kubernetes:
          ` + prob.Kubernetes + `
        tencentcloud:
          ` + prob.Tencentcloud + `
        oracle:
          ` + prob.Oracle + `
        cloudfoundry:
          ` + prob.Cloudfoundry + `
`

	str = strings.Replace(str, "\t", "  ", -1)

	return str
}

func (KustomizeData Kustomize) GetPatchSizing(header string) string {
	str := header + `
  spinnakerConfig:
    config:
      deploymentEnvironment:
        customSizing:
          # This applies sizings to the clouddriver container as well as any sidecar 
          # containers running with clouddriver. (Use without spin- to only include the clouddriver container)
          spin-clouddriver:
            replicas: 1     # Set the number of replicas (pods) to be created for this service.
            # limits:
            #   cpu: 2        # Set the kubernetes CPU limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 4 cpu.
            #   memory: 4Gi   # Set the kubernetes memory limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 16Gi of memory.
            # requests:
            #   cpu: 500m
            #   memory: 2Gi
          # This applies sizings to only the service container and not to any sidecar 
          # containers running with service. (Use spin-<service> to include sidecars)
          clouddriver:
            limits:
              cpu: 2        # Set the kubernetes CPU limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 4 cpu.
              memory: 4Gi   # Set the kubernetes memory limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 16Gi of memory.
            requests:
              cpu: 500m
              memory: 2Gi
          deck:
`

	return str
}

func (KustomizeData Kustomize) GetArmoryPatch(header string) string {
	str := header + `
  spinnakerConfig:
    config:
      armory:
        dinghy:
          enabled: true
          templateOrg: cloudtools
          templateRepo: dinghy-templates
          githubToken: XXXX
          githubEndpoint: https://github.dev.dou.com/api/v3
          repositoryRawdataProcessing: false
          dinghyFilename: dinghyfile
          autoLockPipelines: true
          fiatUser: svc-spinnaker-dinghy
          notifiers:
            slack:
              enabled: true
              channel: pbd-spinnaker-slack-testing
        diagnostics:
          enabled: false
          uuid: XXXX
          logging:
            enabled: false
            endpoint: https://debug.armory.io/v1/logs
        terraform:
          enabled: true
          git:
            enabled: true
            accessToken: XXXX
        secrets:
          vault:
            enabled: false
            path: kubernetes
`

	return str
}
