package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetCloudfoundry(providersRef *providers.Providers) error {

	// if nil != providersRef.Cloudfoundry {
	// 	return fmt.Errorf("Cloudfoundry value is null")
	// }

	if providersRef.Cloudfoundry.Enabled {
		ProvidersData.Enable = "cloudfoundry"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Cloudfoundry.Enabled) + `
	primaryAccount: ` + providersRef.Cloudfoundry.PrimaryAccount +
		GetCloudfoundryAccounts(providersRef.Cloudfoundry.Accounts)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Cloudfoundry = str

	return nil
}

func GetCloudfoundryAccounts(accounts []*providers.CFAccounts) string {
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
		      password: ` + account.Password + `
		      user: ` + account.User + `
		      skipSslValidation: ` + strconv.FormatBool(account.SkipSslValidation) + `
		      api: ` + account.Api + `
		      appsManagerUri: ` + account.AppsManagerUri + `
		      metricsUri: ` + account.MetricsUri
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
