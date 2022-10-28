package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetAzure(providersRef *providers.Providers) error {

	// if nil != providersRef.Azure {
	// 	return fmt.Errorf("Azure value is null")
	// }

	if providersRef.Azure.Enabled {
		ProvidersData.Enable = "azure"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Azure.Enabled) + `
	primaryAccount: ` + providersRef.Azure.PrimaryAccount +
		GetAzureAccounts(providersRef.Azure.Accounts)
		//TODO
		//GetAzureBakeryDefaults(providersRef.Azure.Accounts)
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
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") + `
		      clientId: ` + account.ClientId + `
		      appKey: ` + account.AppKey + `
		      tenantId: ` + account.TenantId + `
		      subscriptionId: ` + account.SubscriptionId + `
		      objectId: ` + account.ObjectId + `
		      defaultResourceGroup: ` + account.DefaultResourceGroup + `
		      defaultKeyVault: ` + account.DefaultKeyVault + `
		      packerResourceGroup: ` + account.PackerResourceGroup + `
		      packerStorageAccount: ` + account.PackerStorageAccount +
				getProvidersStringArrayAppend(account.Regions, "regions", "- ") + `
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
