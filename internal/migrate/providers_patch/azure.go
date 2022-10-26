package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func (ProvidersData *Providers) SetAzure(KustomizeData structs.Kustomize) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Azure {
	// 	return fmt.Errorf("Azure value is null")
	// }

	if KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Azure.Enabled {
		ProvidersData.Enable = "azure"
	}

	str := `enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Azure.Enabled) + `
	primaryAccount: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Azure.PrimaryAccount +
		GetAzureAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Azure.Accounts)
		//TODO
		//GetAzureBakeryDefaults(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Azure.Accounts)
	// bakeryDefaults:
	//   templateFile: azure-linux.json
	//   baseImages: []`

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Azure = str

	return nil
}

func GetAzureAccounts(accounts []*providers.AzureAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += `
		    - name: ` + account.Name + `
		        environment: ` + account.Environment +
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		        clientId: ` + account.ClientId + `
		        appKey: ` + account.AppKey + `
		        tenantId: ` + account.TenantId + `
		        subscriptionId: ` + account.SubscriptionId + `
		        objectId: ` + account.ObjectId + `
		        defaultResourceGroup: ` + account.DefaultResourceGroup + `
		        defaultKeyVault: ` + account.DefaultKeyVault + `
		        packerResourceGroup: ` + account.PackerResourceGroup + `
		        packerStorageAccount: ` + account.PackerStorageAccount +
				getProvidersStringArray(account.Regions, "regions") + `
		        useSshPublicKey: ` + account.UseSshPublicKey + `
		        permission: {}` //TODO + account.Permission`
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
