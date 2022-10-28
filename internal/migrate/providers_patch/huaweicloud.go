package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
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
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      accountType: ` + account.AccountType + `
		      authUrl: ` + account.AuthUrl + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      projectName: ` + account.ProjectName + `
		      domainName: ` + account.DomainName + `
		      insecure: ` + strconv.FormatBool(account.Insecure) +
				getProvidersStringArray(account.Regions, "regions") + `
		      permission: {}` //TODO + account.Permission`
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
		    templateFile: ` + bakeryDefault.TemplateFile + `
		    baseImages: []` + /*TODO bakeryDefault.BaseImages + */ `
		    authUrl: ` + bakeryDefault.AuthUrl + `
		    username: ` + bakeryDefault.Username + `
		    password: ` + bakeryDefault.Password + `
		    projectName: ` + bakeryDefault.ProjectName + `
		    domainName: ` + bakeryDefault.DomainName + `
		    insecure: ` + strconv.FormatBool(bakeryDefault.Insecure) + `
		    vpcId: ` + bakeryDefault.VpcId + `
		    subnetId: ` + bakeryDefault.SubnetId + `
		    securityGroup: ` + bakeryDefault.SecurityGroup + `
		    eipBandwidthSize: ` + strconv.FormatInt(int64(bakeryDefault.EipBandwidthSize), 10)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
