package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetTencentcloud(providersRef *providers.Providers) error {

	// if nil != providersRef.Tencentcloud {
	// 	return fmt.Errorf("Tencent value is null")
	// }

	if providersRef.Tencentcloud.Enabled {
		ProvidersData.Enable = "tencentcloud"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Tencentcloud.Enabled) + `
	primaryAccount: ` + providersRef.Tencentcloud.PrimaryAccount +
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
			str += `
		    - name: ` + account.Name + `
		      environment: ` + account.Environment +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) + `
		      secretId: ` + account.SecretId + `
		      secretKey: ` + account.SecretKey +
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
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefault.TemplateFile +
			GetTencentBaseImages(bakeryDefault.BaseImages) + `
		    secretId: ` + bakeryDefault.SecretId + `
		    secretKey: ` + bakeryDefault.SecretKey
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
			  - baseImage:
			    id: ` + baseImage.BaseImage.Id + `
			    shortDescription: ` + baseImage.BaseImage.ShortDescription + `
			    detailedDescription: ` + baseImage.BaseImage.DetailedDescription + `
			    packageType: ` + baseImage.BaseImage.PackageType + `
			    templateFile: ` + baseImage.BaseImage.TemplateFile
			}
			for _, virtualSetting := range baseImage.VirtualizationSettings {
				if nil != virtualSetting {
					str += `
			    virtualizationSettings:
			      - region: ` + virtualSetting.Region + `
			        zone: ` + virtualSetting.Zone + `
			        instanceType: ` + virtualSetting.InstanceType + `
			        sourceImageId: ` + virtualSetting.SourceImageId + `
			        sshUserName: ` + virtualSetting.SshUserName
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
