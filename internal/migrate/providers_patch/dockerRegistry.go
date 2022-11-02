package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) +
				getPermissions(account.Permissions) + `
		    address: ` + account.Address + `
		    username: ` + account.Username + `
		    email: ` + account.Email + `
		    cacheIntervalSeconds: ` + helpers.IntToString(account.CacheIntervalSeconds) + `
		    clientTimeoutMillis: ` + helpers.IntToString(account.ClientTimeoutMillis) + `
		    cacheThreads: ` + helpers.IntToString(account.CacheThreads) + `
		    paginateSize: ` + helpers.IntToString(account.PaginateSize) + `
		    sortTagsByDate: ` + strconv.FormatBool(account.SortTagsByDate) + `
		    trackDigests: ` + strconv.FormatBool(account.TrackDigests) + `
		    insecureRegistry: ` + strconv.FormatBool(account.InsecureRegistry) + `
		    environment: ` + account.Environment + `
		    password: ` + account.Password + `
		    passwordCommand: ` + account.PasswordCommand +
				strings.Replace(getProvidersStringArrayAppend(account.Repositories, "repositories", "- "), "\t", "   ", -1) + `
		    passwordFile: ` + account.PasswordFile
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
