package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetEcs(providersRef *providers.Providers) error {

	// if nil != providersRef.Ecs {
	// 	return fmt.Errorf("Ecs value is null")
	// }

	if providersRef.Ecs.Enabled {
		ProvidersData.Enable = "ecs"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Ecs.Enabled) + `
	primaryAccount: ` + providersRef.Ecs.PrimaryAccount +
		GetEcsAccounts(providersRef.Ecs.Accounts)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Ecs = str

	return nil
}

func GetEcsAccounts(accounts []*providers.EcsAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += `
		    - name: ` + account.Name +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) + `
		      awsAccount: ` + account.AwsAccount + `
		      environment: ` + account.Environment
			//TODO assumeRole Missing proto
			//TODO providerVersion Missing proto
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
