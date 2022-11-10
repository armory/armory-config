package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetCloudfoundry(providersRef *providers.Providers) error {

	// if nil != providersRef.Cloudfoundry {
	// 	return fmt.Errorf("Cloudfoundry value is null")
	// }

	if providersRef.Cloudfoundry.Enabled {
		ProvidersData.Enable = "cloudfoundry"
	}

	str := helpers.PrintFmtBool(`enabled: `, providersRef.Oracle.Enabled, 5, true) +
		helpers.PrintFmtStr(`primaryAccount: `, providersRef.Oracle.PrimaryAccount, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`user: `, account.User, 7, true) +
				helpers.PrintFmtBool(`skipSslValidation: `, account.SkipSslValidation, 7, true) +
				helpers.PrintFmtStr(`api: `, account.Api, 7, true) +
				helpers.PrintFmtStr(`appsManagerUri: `, account.AppsManagerUri, 7, true) +
				helpers.PrintFmtStr(`metricsUri: `, account.MetricsUri, 7, true)
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
