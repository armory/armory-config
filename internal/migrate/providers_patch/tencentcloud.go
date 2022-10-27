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
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      secretId: ` + account.SecretId + `
		      secretKey: ` + account.SecretKey +
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

func GetTencentBakeryDefaultsAccounts(bakeryDefault *providers.TencentBakery) string {
	str := ""

	if nil != bakeryDefault {
		str += `
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefault.TemplateFile + `
		    baseImages: []` /* + TODO bakeryDefault.BaseImages + */
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
