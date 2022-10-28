package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetOracle(providersRef *providers.Providers) error {

	// if nil != providersRef.Oracle {
	// 	return fmt.Errorf("Oracle value is null")
	// }

	if providersRef.Oracle.Enabled {
		ProvidersData.Enable = "oracle"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Oracle.Enabled) + `
	primaryAccount: ` + providersRef.Oracle.PrimaryAccount +
		GetOracleAccounts(providersRef.Oracle.Accounts) +
		GetOracleBakeryDefaultsAccounts(providersRef.Oracle.BakeryDefaults)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Oracle = str

	return nil
}

func GetOracleAccounts(accounts []*providers.OracleAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += `
		    - name: ` + account.Name + `
		      environment: ` + account.Environment +
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      compartmentId: ` + account.CompartmentId + `
		      userId: ` + account.UserId + `
		      fingerprint: ` + account.Fingerprint + `
		      sshPrivateKeyFilePath: ` + account.SshPrivateKeyFilePath + `
		      privateKeyPassphrase: ` + account.PrivateKeyPassphrase + `
		      tenancyId: ` + account.TenancyId + `
		      region: ` + account.Region + `
		      permission: {}` //TODO + account.Permission`
			//   getProvidersStringArray(account.Regions, "regions") + `
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetOracleBakeryDefaultsAccounts(bakeryDefault *providers.OracleBakery) string {
	str := ""

	if nil != bakeryDefault {
		str += `
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefault.TemplateFile + `
		    baseImages: []` + /*TODO bakeryDefault.BaseImages + */ `
		    availabilityDomain: ` + bakeryDefault.AvailabilityDomain + `
		    subnetId: ` + bakeryDefault.SubnetId + `
		    instanceShape: ` + bakeryDefault.InstanceShape
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
