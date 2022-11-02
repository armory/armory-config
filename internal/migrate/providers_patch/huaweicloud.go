package providers_patch

import (
	"strconv"
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

	str := `enabled: ` + strconv.FormatBool(providersRef.Huaweicloud.Enabled) + `
	primaryAccount: ` + providersRef.Huaweicloud.PrimaryAccount +
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
			str += `
		    - name: ` + account.Name + `
		      environment: ` + account.Environment +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) + `
		      accountType: ` + account.AccountType + `
		      authUrl: ` + account.AuthUrl + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      projectName: ` + account.ProjectName + `
		      domainName: ` + account.DomainName + `
		      insecure: ` + strconv.FormatBool(account.Insecure) +
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
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefault.TemplateFile +
			GetHuaweiBaseImages(bakeryDefault.BaseImages) + `
		    authUrl: ` + bakeryDefault.AuthUrl + `
		    username: ` + bakeryDefault.Username + `
		    password: ` + bakeryDefault.Password + `
		    projectName: ` + bakeryDefault.ProjectName + `
		    domainName: ` + bakeryDefault.DomainName + `
		    insecure: ` + strconv.FormatBool(bakeryDefault.Insecure) + `
		    vpcId: ` + bakeryDefault.VpcId + `
		    subnetId: ` + bakeryDefault.SubnetId + `
		    securityGroup: ` + bakeryDefault.SecurityGroup + `
		    eipBandwidthSize: ` + helpers.IntToString(bakeryDefault.EipBandwidthSize)
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
			        instanceType: ` + virtualSetting.InstanceType + `
			        sourceImageId: ` + virtualSetting.SourceImageId + `
			        sshUserName: ` + virtualSetting.SshUserName + `
			        eiptype: ` + virtualSetting.Eiptype //TODO Missing Proto (typo)
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
