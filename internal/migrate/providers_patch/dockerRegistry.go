package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetDockerRegistry(providersRef *providers.Providers) error {

	// if nil != providersRef.DockerRegistry {
	// 	return fmt.Errorf("DockerRegistry value is null")
	// }

	if providersRef.DockerRegistry.Enabled {
		ProvidersData.Enable = "dockerRegistry"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.DockerRegistry.Enabled) + `
	primaryAccount: ` + providersRef.DockerRegistry.PrimaryAccount +
		GetDockerRegistryAccounts(providersRef.DockerRegistry.Accounts)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.DockerRegistry = str

	return nil
}

func GetDockerRegistryAccounts(accounts []*providers.DockerRegistryAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += `
		    - name: ` + account.Name +
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      address: ` + account.Address + `
		      username: ` + account.Username + `
		      email: ` + account.Email + `
		      cacheIntervalSeconds: ` + strconv.FormatInt(int64(account.CacheIntervalSeconds), 10) + `
		      clientTimeoutMillis: ` + strconv.FormatInt(int64(account.ClientTimeoutMillis), 10) + `
		      cacheThreads: ` + strconv.FormatInt(int64(account.CacheThreads), 10) + `
		      paginateSize: ` + strconv.FormatInt(int64(account.PaginateSize), 10) + `
		      sortTagsByDate: ` + strconv.FormatBool(account.SortTagsByDate) + `
		      trackDigests: ` + strconv.FormatBool(account.TrackDigests) + `
		      insecureRegistry: ` + strconv.FormatBool(account.InsecureRegistry) + `
		      environment: ` + account.Environment + `
		      password: ` + account.Password + `
		      passwordCommand: ` + account.PasswordCommand +
				getProvidersStringArray(account.Repositories, "repositories") + `
		      passwordFile: ` + account.PasswordFile + `
		      permission: {}` //TODO + account.Permission`
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
