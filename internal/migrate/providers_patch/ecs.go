package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetEcs(providersRef *providers.Providers) error {

	if nil != providersRef.Ecs {
		if providersRef.Ecs.Enabled {
			ProvidersData.Enable = "ecs"
		}
		str := helpers.PrintFmtBool(`enabled: `, providersRef.Ecs.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.Ecs.PrimaryAccount, 5, true) +
			GetEcsAccounts(providersRef.Ecs.Accounts)

		str = strings.Replace(str, "\t", "          ", -1)
		ProvidersData.Ecs = str
	} else {
		ProvidersData.Ecs = " {}"
		// 	return fmt.Errorf("Ecs value is null")
	}

	return nil
}

func GetEcsAccounts(accounts []*providers.EcsAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`awsAccount: `, account.AwsAccount, 7, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true)
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
