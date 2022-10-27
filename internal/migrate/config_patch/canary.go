package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/canary"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetCanary(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Canary {
		str = GetCanaryConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Canary)
	}
	return str
}

func GetCanaryConfig(canaryReference *canary.Canary) string {
	str := `
		enabled: ` + strconv.FormatBool(canaryReference.Enabled)

	if nil != canaryReference.ServiceIntegrations {
		str += GetCanaryServiceIntegrations(canaryReference)
	}
	str += `
		reduxLogger: ` + strconv.FormatBool(canaryReference.ReduxLogger) + `
		defaultJudge: ` + canaryReference.DefaultJudge + `
		stagesEnabled: ` + strconv.FormatBool(canaryReference.StagesEnabled) + `
		templatesEnabled: ` + strconv.FormatBool(canaryReference.TemplatesEnabled) + `
		showAllConfigsEnabled: ` + strconv.FormatBool(canaryReference.ShowAllConfigsEnabled)

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}

func GetCanaryServiceIntegrations(canary *canary.Canary) string {
	str := ""

	// Accounts           []*CanaryAccounts `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
	if nil != canary.ServiceIntegrations {
		str += `
		serviceIntegrations:`

		for _, serviceInteg := range canary.ServiceIntegrations {
			str += `
		  - name: ` + serviceInteg.Name + `
		    enabled: ` + strconv.FormatBool(serviceInteg.Enabled) + `
		    gcsEnabledd: ` + strconv.FormatBool(serviceInteg.GcsEnabled) +
				getCanaryAccounts(serviceInteg.Accounts) + `
		    stackdriverEnabled: ` + strconv.FormatBool(serviceInteg.StackdriverEnabled) + `
		    s3Enabled: ` + strconv.FormatBool(serviceInteg.S3Enabled)
		}
	} else {
		str += `
		serviceIntegrations: []`
	}

	return str
}

func getCanaryAccounts(accounts []*canary.CanaryAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		    accounts:`
		for _, account := range accounts {
			str += `
		      - name: ` + account.Name + `
		        project: ` + account.Project + `
		        jsonPath: ` + account.JsonPath + `
		        bucket: ` + account.Bucket + `
		        bucketLocation: ` + account.BucketLocation + `
		        rootFolder: ` + account.RootFolder +
				getSupportedTypes(account.SupportedTypes) + `
		        endpoint:
		          baseUrl: ` + account.Endpoint.BaseUrl + `
		        username: ` + account.Username + `
		        password: ` + account.Password + `
		        usernamePasswordFile: ` + account.UsernamePasswordFile + `
		        apiKey: ` + account.ApiKey + `
		        applicationKey: ` + account.ApplicationKey + `
		        accessToken: ` + account.AccessToken + `
		        defaultScopeKey: ` + account.DefaultScopeKey + `
		        defaultLocationKey: ` + account.DefaultLocationKey + `
		        region: ` + account.Region + `
		        profileName: ` + account.ProfileName + `
		        accessKeyId: ` + account.AccessKeyId + `
		        secretAccessKey: ` + account.SecretAccessKey
		}
	} else {
		str += `
		    accounts: []`
	}

	return str
}

func getSupportedTypes(supTypes []string) string {
	str := ""

	if nil != supTypes {
		str += `
		          supportedTypes:`
		for _, supType := range supTypes {
			str += `
		            ` + supType
		}
	} else {
		str += `
		          supportedTypes: []`
	}

	return str
}
