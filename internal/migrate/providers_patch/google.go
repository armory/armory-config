package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetGoogle(providersRef *providers.Providers) error {

	// if nil != providersRef.Google {
	// 	return fmt.Errorf("Google value is null")
	// }

	if providersRef.Google.Enabled {
		ProvidersData.Enable = "google"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Google.Enabled) + `
	primaryAccount: ` + providersRef.Google.PrimaryAccount +
		GetGoogleAccounts(providersRef.Google.Accounts) +
		GetGoogleBakeryDefaultsAccounts(providersRef.Google.BakeryDefaults)

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
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) + `
		      project: ` + account.Project + `
		      jsonPath: ` + account.JsonPath + `
		      alphaListed: ` + strconv.FormatBool(account.AlphaListed) +
				getProvidersStringArrayAppend(account.ImageProjects, "imageProjects", "- ") + `
		      environment: ` + account.Environment + `
		      consul:` + account.Environment + `
		        enabled: ` + strconv.FormatBool(account.Consul.Enabled) + `
		        agentEndpoint: ` + account.Consul.AgentEndpoint + `
		        agentPort: ` + helpers.IntToString(account.Consul.AgentPort) +
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
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefault.TemplateFile +
			GetGoogleBaseImages(bakeryDefault.BaseImages) + `
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

func GetGoogleBaseImages(baseImages []*providers.GoogleBaseImages) string {
	str := ""

	if nil != baseImages {
		str += `
		    baseImages:`
		if nil != baseImages {
			for _, baseImage := range baseImages {
				if nil != baseImage.BaseImage {
					str += `
			- baseImage:
			    id: ` + baseImage.BaseImage.Id + `
			    shortDescription: ` + baseImage.BaseImage.ShortDescription + `
			    detailedDescription: ` + baseImage.BaseImage.DetailedDescription + `
			    packageType: ` + baseImage.BaseImage.PackageType + `
			    templateFile: ` + baseImage.BaseImage.TemplateFile + `
			    imageFamily: ` + strconv.FormatBool(baseImage.BaseImage.ImageFamily)
				}
				if nil != baseImage.VirtualizationSettings {
					str += `
			  virtualizationSettings:`
					virtualSettings := baseImage.VirtualizationSettings
					//for _, virtualSettings := range baseImage.VirtualizationSettings {
					str += `
			    sourceImage: ` + virtualSettings.SourceImage + `
			    sourceImageFamily: ` + virtualSettings.SourceImageFamily
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
