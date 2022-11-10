package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetAzure(providersRef *providers.Providers) error {

	// if nil != providersRef.Azure {
	// 	return fmt.Errorf("Azure value is null")
	// }

	if providersRef.Azure.Enabled {
		ProvidersData.Enable = "azure"
	}

	str := helpers.PrintFmtBool(`enabled: `, providersRef.Azure.Enabled, 5, true) +
		helpers.PrintFmtStr(`primaryAccount: `, providersRef.Azure.PrimaryAccount, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`clientId: `, account.ClientId, 7, true) +
				helpers.PrintFmtStrApostrophe(`appKey: `, account.AppKey, 7, true) +
				helpers.PrintFmtStrApostrophe(`tenantId: `, account.TenantId, 7, true) +
				helpers.PrintFmtStrApostrophe(`subscriptionId: `, account.SubscriptionId, 7, true) +
				helpers.PrintFmtStrApostrophe(`objectId: `, account.ObjectId, 7, true) +
				helpers.PrintFmtStr(`defaultResourceGroup: `, account.DefaultResourceGroup, 7, true) +
				helpers.PrintFmtStr(`defaultKeyVault: `, account.DefaultKeyVault, 7, true) +
				helpers.PrintFmtStr(`packerResourceGroup: `, account.PackerResourceGroup, 7, true) +
				helpers.PrintFmtStr(`packerStorageAccount: `, account.PackerStorageAccount, 7, true) +
				getProvidersStringArrayAppend(account.Regions, "regions", "- ") +
				helpers.PrintFmtStrApostrophe(`useSshPublicKey: `, account.UseSshPublicKey, 7, true)
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
		  bakeryDefaults:` +
			helpers.PrintFmtStr(`templateFile: `, bakeryDefaults.TemplateFile, 6, true) +
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
			  - baseImage:` +
					helpers.PrintFmtStr(`id: `, baseImage.BaseImage.Id, 8, true) +
					helpers.PrintFmtStr(`shortDescription: `, baseImage.BaseImage.ShortDescription, 8, true) +
					helpers.PrintFmtStr(`detailedDescription: `, baseImage.BaseImage.DetailedDescription, 8, true) +
					helpers.PrintFmtStr(`packageType: `, baseImage.BaseImage.PackageType, 8, true) +
					helpers.PrintFmtStr(`templateFile: `, baseImage.BaseImage.TemplateFile, 8, true) +
					helpers.PrintFmtStr(`publisher: `, baseImage.BaseImage.Publisher, 8, true) +
					helpers.PrintFmtStr(`offer: `, baseImage.BaseImage.Offer, 8, true) +
					helpers.PrintFmtStrApostrophe(`sku: `, baseImage.BaseImage.Sku, 8, true) +
					helpers.PrintFmtStr(`version: `, baseImage.BaseImage.Version, 8, true)
			}
		}
	} else {
		str += `
		    baseImages: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
