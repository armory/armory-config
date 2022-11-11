package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetGoogle(providersRef *providers.Providers) error {

	if nil != providersRef.Google {
		if providersRef.Google.Enabled {
			ProvidersData.Enable = "google"
		}

		str := helpers.PrintFmtBool(`enabled: `, providersRef.Google.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.Google.PrimaryAccount, 5, true) +
			GetGoogleAccounts(providersRef.Google.Accounts) +
			GetGoogleBakeryDefaultsAccounts(providersRef.Google.BakeryDefaults)

		str = strings.Replace(str, "\t", "          ", -1)
		ProvidersData.Google = str
	} else {
		ProvidersData.Google = " {}"
		// 	return fmt.Errorf("Google value is null")
	}

	return nil
}

func GetGoogleAccounts(accounts []*providers.GoogleAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`project: `, account.Project, 7, true) +
				helpers.PrintFmtStr(`jsonPath: `, account.JsonPath, 7, true) +
				helpers.PrintFmtBool(`alphaListed: `, account.AlphaListed, 7, true) +
				getProvidersStringArrayAppend(account.ImageProjects, "imageProjects", "- ") +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) + `
		      consul:` +
				//TODO Add check to avoid null pointer
				helpers.PrintFmtBool(`enabled: `, account.Consul.Enabled, 8, true) +
				helpers.PrintFmtStr(`agentEndpoint: `, account.Consul.AgentEndpoint, 8, true) +
				helpers.PrintFmtInt(`agentPort: `, account.Consul.AgentPort, 8, true) +
				strings.Replace(getProvidersStringArray(account.Consul.Datacenters, "datacenters"), "\t", "     ", -1) +
				getProvidersStringArrayAppend(account.Regions, "regions", "- ")
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
		  bakeryDefaults:` +
			helpers.PrintFmtStr(`templateFile: `, bakeryDefault.TemplateFile, 6, true) +
			GetGoogleBaseImages(bakeryDefault.BaseImages) +
			helpers.PrintFmtStr(`zone: `, bakeryDefault.Zone, 6, true) +
			helpers.PrintFmtStr(`network: `, bakeryDefault.Network, 6, true) +
			helpers.PrintFmtBool(`useInternalIp: `, bakeryDefault.UseInternalIp, 6, true)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetGoogleBaseImages(baseImages []*providers.GoogleBaseImages) string {
	str := ""

	if nil != baseImages {
		str += `
		    baseImages:`
		if nil != baseImages {
			for _, baseImage := range baseImages {
				if nil != baseImage.BaseImage {
					str += `
			- baseImage:` +
						helpers.PrintFmtStr(`id: `, baseImage.BaseImage.Id, 8, true) +
						helpers.PrintFmtStr(`shortDescription: `, baseImage.BaseImage.ShortDescription, 8, true) +
						helpers.PrintFmtStr(`detailedDescription: `, baseImage.BaseImage.DetailedDescription, 8, true) +
						helpers.PrintFmtStr(`packageType: `, baseImage.BaseImage.PackageType, 8, true) +
						helpers.PrintFmtStr(`templateFile: `, baseImage.BaseImage.TemplateFile, 8, true) +
						helpers.PrintFmtBool(`imageFamily: `, baseImage.BaseImage.ImageFamily, 8, true)
				}
				if nil != baseImage.VirtualizationSettings {
					str += `
			  virtualizationSettings:`
					virtualSettings := baseImage.VirtualizationSettings
					//for _, virtualSettings := range baseImage.VirtualizationSettings {
					str += helpers.PrintFmtStr(`sourceImage: `, virtualSettings.SourceImage, 8, true) +
						helpers.PrintFmtStr(`sourceImageFamily: `, virtualSettings.SourceImageFamily, 8, true)
					// }
				} else {
					str += `
					virtualizationSettings: []`
				}
			}
		}
	} else {
		str += `
		    baseImages: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
