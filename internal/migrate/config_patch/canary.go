package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/canary"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
	str := helpers.PrintFmtBool(`enabled: `, canaryReference.Enabled, 4, true)

	if nil != canaryReference.ServiceIntegrations {
		str += GetCanaryServiceIntegrations(canaryReference)
	}
	str += helpers.PrintFmtBool(`reduxLoggerEnabled: `, canaryReference.ReduxLogger, 4, true) +
		helpers.PrintFmtStr(`defaultJudge: `, canaryReference.DefaultJudge, 4, true) +
		helpers.PrintFmtBool(`stagesEnabled: `, canaryReference.StagesEnabled, 4, true) +
		helpers.PrintFmtBool(`templatesEnabled: `, canaryReference.TemplatesEnabled, 4, true) +
		helpers.PrintFmtBool(`showAllConfigsEnabled: `, canaryReference.ShowAllConfigsEnabled, 4, true)

	return str
}

func GetCanaryServiceIntegrations(canary *canary.Canary) string {
	str := ""

	// Accounts           []*CanaryAccounts `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
	if nil != canary.ServiceIntegrations {
		str += `
		serviceIntegrations:`

		for _, serviceInteg := range canary.ServiceIntegrations {
			str += helpers.PrintFmtStr(`- name: `, serviceInteg.Name, 5, true) +
				helpers.PrintFmtBool(`enabled: `, serviceInteg.Enabled, 6, true) +
				getCanaryAccounts(serviceInteg.Accounts)

			if serviceInteg.Name == `aws` {
				str += helpers.PrintFmtBool(`s3Enabled: `, serviceInteg.S3Enabled, 6, true)

			}
			if serviceInteg.Name == `google` {
				str += helpers.PrintFmtBool(`gcsEnabled: `, serviceInteg.GcsEnabled, 6, true) +
					helpers.PrintFmtBool(`stackdriverEnabled: `, serviceInteg.StackdriverEnabled, 6, true)

			}
			// helpers.PrintFmtBool(`s3Enabled: `, serviceInteg.S3Enabled, 6, true) +
			// helpers.PrintFmtBool(`gcsEnabled: `, serviceInteg.GcsEnabled, 6, true)
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

			//To avoid null pointers
			baseUrlEndpoint := ""
			if nil != account.Endpoint {
				baseUrlEndpoint = account.Endpoint.BaseUrl
			}

			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`project: `, account.Project, 7, true) +
				helpers.PrintFmtStr(`jsonPath: `, account.JsonPath, 7, true) +
				helpers.PrintFmtStr(`bucket: `, account.Bucket, 7, true) +
				helpers.PrintFmtStr(`bucketLocation: `, account.BucketLocation, 7, true) +
				helpers.PrintFmtStr(`rootFolder: `, account.RootFolder, 7, true) +
				getSupportedTypes(account.SupportedTypes) + `
		      endpoint:` +
				helpers.PrintFmtStr(`baseUrl: `, baseUrlEndpoint, 8, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`usernamePasswordFile: `, account.UsernamePasswordFile, 7, true) +
				helpers.PrintFmtStr(`apiKey: `, account.ApiKey, 7, true) +
				helpers.PrintFmtStr(`applicationKey: `, account.ApplicationKey, 7, true) +
				helpers.PrintFmtStr(`accessToken: `, account.AccessToken, 7, true) +
				helpers.PrintFmtStr(`defaultScopeKey: `, account.DefaultScopeKey, 7, true) +
				helpers.PrintFmtStr(`defaultLocationKey: `, account.DefaultLocationKey, 7, true) +
				helpers.PrintFmtStr(`region: `, account.Region, 7, true) +
				helpers.PrintFmtStr(`profileName: `, account.ProfileName, 7, true) +
				helpers.PrintFmtStr(`accessKeyId: `, account.AccessKeyId, 7, true) +
				helpers.PrintFmtStr(`secretAccessKey: `, account.SecretAccessKey, 7, true)
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
		        - ` + supType
		}
	} else {
		str += `
		      supportedTypes: []`
	}

	return str
}
