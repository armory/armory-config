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
		GetAzureAccounts(providersRef.Azure.Accounts) +
		GetAzureBakeryDefaults(providersRef.Azure.BakeryDefaults)

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
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) + `
		      clientId: ` + account.ClientId + `
		      appKey: '` + account.AppKey + `'
		      tenantId: '` + account.TenantId + `'
		      subscriptionId: '` + account.SubscriptionId + `'
		      objectId: '` + account.ObjectId + `'
		      defaultResourceGroup: ` + account.DefaultResourceGroup + `
		      defaultKeyVault: ` + account.DefaultKeyVault + `
		      packerResourceGroup: ` + account.PackerResourceGroup + `
		      packerStorageAccount: ` + account.PackerStorageAccount +
				getProvidersStringArrayAppend(account.Regions, "regions", "- ") + `
		      useSshPublicKey: ` + account.UseSshPublicKey
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetAzureBakeryDefaults(bakeryDefaults *providers.AzureBakeryDefaults) string {
	str := ""

	if nil != bakeryDefaults {
		str += `
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefaults.TemplateFile +
			GetAzureBaseImages(bakeryDefaults.BaseImages)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetAzureBaseImages(baseImages []*providers.AzureBaseImages) string {
	str := ""

	if nil != baseImages {
		str += `
		    baseImages:`
		for _, baseImage := range baseImages {
			if nil != baseImage.BaseImage {
				str += `
			  - baseImage:
			    id: ` + baseImage.BaseImage.Id + `
			    shortDescription: ` + baseImage.BaseImage.ShortDescription + /*`
					  detailedDescription: ` + baseImage.BaseImage.DetailedDescription +*/`
			    packageType: ` + baseImage.BaseImage.PackageType + `
			    templateFile: ` + baseImage.BaseImage.TemplateFile + `
			    publisher: ` + baseImage.BaseImage.Publisher + `
			    offer: ` + baseImage.BaseImage.Offer + `
			    sku: '` + baseImage.BaseImage.Sku + `'
			    version: '` + baseImage.BaseImage.Version + `'`
			}
		}
	} else {
		str += `
		    baseImages: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
