package providers_patch

import (
	"fmt"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

type Providers struct {
	AppEngine      string
	Aws            string
	Ecs            string
	Azure          string
	Dcos           string
	DockerRegistry string
	Google         string
	Huaweicloud    string
	Kubernetes     string
	Tencentcloud   string
	Oracle         string
	Cloudfoundry   string
	Enable         string //Which provider is currently enable
}

func GetProvidersData(KustomizeData structs.Kustomize) string {

	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers {
		prob := Providers{}
		prob.SetProvidersData(KustomizeData)

		str = `
  validation:
    providers:
      ` + prob.Enable + `:
        enabled: true
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
          ` + prob.Cloudfoundry
	}

	return str
}

// This function fills the Providers struct in a valid format
func (ProvidersData *Providers) SetProvidersData(KustomizeData structs.Kustomize) error {

	// Slice of function calls to fill all the Providers Data
	functionCalls := []func(providersRef *providers.Providers) error{
		ProvidersData.SetAppEngineData,
		ProvidersData.SetAwsData,
		ProvidersData.SetEcs,
		ProvidersData.SetAzure,
		ProvidersData.SetDcos,
		ProvidersData.SetDockerRegistry,
		ProvidersData.SetGoogle,
		ProvidersData.SetHuaweicloud,
		ProvidersData.SetKubernetes,
		ProvidersData.SetTencentcloud,
		ProvidersData.SetOracle,
		ProvidersData.SetCloudfoundry,
	}

	for _, function := range functionCalls {
		err := function(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers)
		if err != nil {
			fmt.Println("Error while executing the function:%s Error:(%s)", function, err)
			return err
			// return fmt.Errorf("Error while executing the function:%s Error:(%s)", function, err)
		}
	}

	return nil
}

func (ProvidersData *Providers) SetDcos(providersRef *providers.Providers) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Dcos {
	// 	return fmt.Errorf("Dcos value is null")
	// }

	// if KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Dcos.Enable {
	// 	ProvidersData.Enable = "dcos"
	// }

	str := `enabled: false
	accounts: []
	clusters: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Dcos.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Dcos = str

	return nil
}

func (ProvidersData *Providers) SetHuaweicloud(providersRef *providers.Providers) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Huaweicloud {
	// 	return fmt.Errorf("Huaweicloud value is null")
	// }

	// if KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Huaweicloud.Enabled {
	// 	ProvidersData.Enable = "huaweicloud"
	// }

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Huaweicloud.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Huaweicloud = str

	return nil
}

func (ProvidersData *Providers) SetTencentcloud(providersRef *providers.Providers) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Tencent {
	// 	return fmt.Errorf("Tencent value is null")
	// }

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Tencentcloud.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Tencentcloud = str

	return nil
}

func (ProvidersData *Providers) SetOracle(providersRef *providers.Providers) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Oracle {
	// 	return fmt.Errorf("Oracle value is null")
	// }

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  templateFile: oci.json
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Oracle.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Oracle = str

	return nil
}

func (ProvidersData *Providers) SetCloudfoundry(providersRef *providers.Providers) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Cloudfoundry {
	// 	return fmt.Errorf("Cloudfoundry value is null")
	// }

	str := `enabled: false
	accounts: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Cloudfoundry.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Cloudfoundry = str

	return nil
}

func getProvidersStringArray(stringArray []string, fieldName string) string {
	str := ""

	if nil != stringArray {
		str += `
		      ` + fieldName + `:`
		for _, stringValue := range stringArray {
			str += `
		        ` + stringValue
		}
	} else {
		str += `
		      ` + fieldName + `: []`
	}

	return str
}
