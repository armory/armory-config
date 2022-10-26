package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func (ProvidersData *Providers) SetEcs(KustomizeData structs.Kustomize) error {

	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Ecs {
	// 	return fmt.Errorf("Ecs value is null")
	// }

	if KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Ecs.Enabled {
		ProvidersData.Enable = "ecs"
	}

	str := `enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Ecs.Enabled) + `
	primaryAccount: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Ecs.PrimaryAccount +
		GetEcsAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Ecs.Accounts)

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
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      awsAccount: ` + account.AwsAccount + `
		      environment: ` + account.Environment + `
		      Permission: {}` //TODO + account.Permission`
			//TODO assumeRole
			//TODO providerVersion
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
