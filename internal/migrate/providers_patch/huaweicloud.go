package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetHuaweicloud(providersRef *providers.Providers) error {

	// if nil != providersRef.Huaweicloud {
	// 	return fmt.Errorf("Huaweicloud value is null")
	// }

	if providersRef.Huaweicloud.Enabled {
		ProvidersData.Enable = "huaweicloud"
	}

	str := helpers.PrintFmtBool(`enabled: `, providersRef.Huaweicloud.Enabled, 5, true) +
		helpers.PrintFmtStr(`primaryAccount: `, providersRef.Huaweicloud.PrimaryAccount, 5, true) +
		GetHuaweiCloudAccounts(providersRef.Huaweicloud.Accounts) +
		GetHuaweiBakeryDefaultsAccounts(providersRef.Huaweicloud.BakeryDefaults)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Huaweicloud = str

	return nil
}

func GetHuaweiCloudAccounts(accounts []*providers.HuaweiAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`accountType: `, account.AccountType, 7, true) +
				helpers.PrintFmtStr(`authUrl: `, account.AuthUrl, 7, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`projectName: `, account.ProjectName, 7, true) +
				helpers.PrintFmtStr(`domainName: `, account.DomainName, 7, true) +
				helpers.PrintFmtBool(`insecure: `, account.Insecure, 7, true) +
				getProvidersStringArrayAppend(account.Regions, "regions", "- ")
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetHuaweiBakeryDefaultsAccounts(bakeryDefault *providers.HuaweiBakery) string {
	str := ""

	if nil != bakeryDefault {
		str += `
		  bakeryDefaults:` +
			helpers.PrintFmtStr(`templateFile: `, bakeryDefault.TemplateFile, 6, true) +
			GetHuaweiBaseImages(bakeryDefault.BaseImages) +
			helpers.PrintFmtStr(`authUrl: `, bakeryDefault.AuthUrl, 6, true) +
			helpers.PrintFmtStr(`username: `, bakeryDefault.Username, 6, true) +
			helpers.PrintFmtStr(`password: `, bakeryDefault.Password, 6, true) +
			helpers.PrintFmtStr(`projectName: `, bakeryDefault.ProjectName, 6, true) +
			helpers.PrintFmtStr(`domainName: `, bakeryDefault.DomainName, 6, true) +
			helpers.PrintFmtBool(`insecure: `, bakeryDefault.Insecure, 6, true) +
			helpers.PrintFmtStr(`vpcId: `, bakeryDefault.VpcId, 6, true) +
			helpers.PrintFmtStr(`subnetId: `, bakeryDefault.SubnetId, 6, true) +
			helpers.PrintFmtStr(`securityGroup: `, bakeryDefault.SecurityGroup, 6, true) +
			helpers.PrintFmtInt(`eipBandwidthSize: `, bakeryDefault.EipBandwidthSize, 6, true)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetHuaweiBaseImages(baseImages []*providers.HuaweiBaseImages) string {
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
						helpers.PrintFmtStr(`instanceType: `, virtualSetting.InstanceType, 10, true) +
						helpers.PrintFmtStr(`sourceImageId: `, virtualSetting.SourceImageId, 10, true) +
						helpers.PrintFmtStr(`sshUserName: `, virtualSetting.SshUserName, 10, true) +
						helpers.PrintFmtStr(`eiptype: `, virtualSetting.EipType, 10, true)
				}
			}
		}
	} else {
		str += `
		    baseImages: []`
	}

	return str
}
