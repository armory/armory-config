package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetDockerRegistry(providersRef *providers.Providers) error {

	if nil != providersRef.DockerRegistry {
		if providersRef.DockerRegistry.Enabled {
			ProvidersData.Enable = "dockerRegistry"
		}

		str := helpers.PrintFmtBool(`enabled: `, providersRef.DockerRegistry.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.DockerRegistry.PrimaryAccount, 5, true) +
			GetDockerRegistryAccounts(providersRef.DockerRegistry.Accounts)

		str = strings.Replace(str, "\t", "          ", -1)
		ProvidersData.DockerRegistry = str
	} else {
		ProvidersData.DockerRegistry = " {}"
		// 	return fmt.Errorf("DockerRegistry value is null")
	}
	return nil
}

func GetDockerRegistryAccounts(accounts []*providers.DockerRegistryAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 5, true) +
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) +
				getPermissions(account.Permissions) +
				helpers.PrintFmtStr(`address: `, account.Address, 6, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 6, true) +
				helpers.PrintFmtStr(`email: `, account.Email, 6, true) +
				helpers.PrintFmtInt(`cacheIntervalSeconds: `, account.CacheIntervalSeconds, 6, true) +
				helpers.PrintFmtInt(`clientTimeoutMillis: `, account.ClientTimeoutMillis, 6, true) +
				helpers.PrintFmtInt(`cacheThreads: `, account.CacheThreads, 6, true) +
				helpers.PrintFmtInt(`paginateSize: `, account.PaginateSize, 6, true) +
				helpers.PrintFmtBool(`sortTagsByDate: `, account.SortTagsByDate, 6, true) +
				helpers.PrintFmtBool(`trackDigests: `, account.TrackDigests, 6, true) +
				helpers.PrintFmtBool(`insecureRegistry: `, account.InsecureRegistry, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 6, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 6, true) +
				helpers.PrintFmtStr(`passwordCommand: `, account.PasswordCommand, 6, true) +
				strings.Replace(getProvidersStringArrayAppend(account.Repositories, "repositories", "- "), "\t", "   ", -1) +
				helpers.PrintFmtStr(`passwordFile: `, account.PasswordFile, 6, true)
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
