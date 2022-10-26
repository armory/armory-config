package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func (ProvidersData *Providers) SetGoogle(KustomizeData structs.Kustomize) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Google {
	// 	return fmt.Errorf("Google value is null")
	// }

	if KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Google.Enabled {
		ProvidersData.Enable = "google"
	}

	str := `enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Google.Enabled) + `
	primaryAccount: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Google.PrimaryAccount +
		GetGoogleAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Google.Accounts) +
		GetGoogleBakeryDefaultsAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Google.BakeryDefault)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Google = str

	return nil
}

func GetGoogleAccounts(accounts []*providers.GoogleAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += `
		    - name: ` + account.Name +
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      project: ` + account.Project + `
		      jsonPath: ` + account.JsonPath + `
		      alphaListed: ` + strconv.FormatBool(account.AlphaListed) +
				getProvidersStringArray(account.ImageProjects, "imageProjects") + `
		      environment: ` + account.Environment + `
		      consul: ` + account.Environment + `
		        enabled: ` + strconv.FormatBool(account.Consul.Enabled) + `
		        agentEndpoint: ` + account.Consul.AgentEndpoint + `
		        agentPort: ` + strconv.FormatInt(int64(account.Consul.AgentPort), 10) +
				getProvidersStringArray(account.Consul.Datacenters, "datacenters") +
				getProvidersStringArray(account.Consul.Regions, "regions") + `
		      permission: {}` //TODO + account.Permission`
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetGoogleBakeryDefaultsAccounts(bakeryDefault *providers.GoogleBakeryDefaults) string {
	str := ""

	if nil != bakeryDefault {
		str += `
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefault.TemplateFile + `
		    baseImages: []` + /*TODO bakeryDefault.BaseImages + */ `
		    zone: ` + bakeryDefault.Zone + `
		    network: ` + bakeryDefault.Network + `
		    useInternalIp: ` + strconv.FormatBool(bakeryDefault.UseInternalIp)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
