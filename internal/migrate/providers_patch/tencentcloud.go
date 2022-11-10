package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetTencentcloud(providersRef *providers.Providers) error {

	// if nil != providersRef.Tencentcloud {
	// 	return fmt.Errorf("Tencent value is null")
	// }

	if providersRef.Tencentcloud.Enabled {
		ProvidersData.Enable = "tencentcloud"
	}

	str := helpers.PrintFmtBool(`enabled: `, providersRef.Tencentcloud.Enabled, 5, true) +
		helpers.PrintFmtStr(`primaryAccount: `, providersRef.Tencentcloud.PrimaryAccount, 5, true) +
		GetTencentcloudAccounts(providersRef.Tencentcloud.Accounts) +
		GetTencentBakeryDefaultsAccounts(providersRef.Tencentcloud.BakeryDefaults)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Tencentcloud = str

	return nil
}

func GetTencentcloudAccounts(accounts []*providers.TencentAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`secretId: `, account.SecretId, 7, true) +
				helpers.PrintFmtStr(`secretKey: `, account.SecretKey, 7, true) +
				getProvidersStringArrayAppend(account.Regions, "regions", "- ")
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetTencentBakeryDefaultsAccounts(bakeryDefault *providers.TencentBakery) string {
	str := ""

	if nil != bakeryDefault {
		str += `
		  bakeryDefaults:` +
			helpers.PrintFmtStr(`templateFile: `, bakeryDefault.TemplateFile, 6, true) +
			GetTencentBaseImages(bakeryDefault.BaseImages) +
			helpers.PrintFmtStr(`secretId: `, bakeryDefault.SecretId, 6, true) +
			helpers.PrintFmtStr(`secretKey: `, bakeryDefault.SecretKey, 6, true)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetTencentBaseImages(baseImages []*providers.TencentBaseImages) string {
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
					helpers.PrintFmtStr(`templateFile: `, baseImage.BaseImage.TemplateFile, 8, true)
			}
			for _, virtualSetting := range baseImage.VirtualizationSettings {
				if nil != virtualSetting {
					str += `
			    virtualizationSettings:` +
						helpers.PrintFmtStr(`- region: `, virtualSetting.Region, 9, true) +
						helpers.PrintFmtStr(`zone: `, virtualSetting.Zone, 10, true) +
						helpers.PrintFmtStr(`instanceType: `, virtualSetting.InstanceType, 10, true) +
						helpers.PrintFmtStr(`sourceImageId: `, virtualSetting.SourceImageId, 10, true) +
						helpers.PrintFmtStr(`sshUserName: `, virtualSetting.SshUserName, 10, true)
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
