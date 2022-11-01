package providers_patch

import (
	"fmt"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
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

func getProvidersStringArrayAppend(stringArray []string, fieldName string, valueToAppend string) string {
	str := ""

	if nil != stringArray {
		str += `
		      ` + fieldName + `:`
		for _, stringValue := range stringArray {
			str += `
		        ` + valueToAppend + stringValue
		}
	} else {
		str += `
		      ` + fieldName + `: []`
	}

	return str
}

func getPermissions(permissionsRef *permissions.Permissions) string {
	str := ""

	if nil != permissionsRef {
		str += `
		    permissions:
		      READ:`
		for _, permissionRead := range permissionsRef.READ {
			str += `
		        - ` + permissionRead
		}
		str += `
		      WRITE:`
		for _, permissionWrite := range permissionsRef.WRITE {
			str += `
		        - ` + permissionWrite
		}

		//If both are 0 this will override the permissions instead of having both write and read fields empty
		if 0 == len(permissionsRef.READ) && 0 == len(permissionsRef.WRITE) {
			str = `
		    permissions: {}`
		}
	} else {
		str = `
		    permissions: {}`
	}

	return str
}
