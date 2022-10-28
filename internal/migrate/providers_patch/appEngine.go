package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetAppEngineData(providersRef *providers.Providers) error {

	// if nil != providersRef.Appengine {
	// 	return fmt.Errorf("Appengine value is null")
	// }

	if providersRef.Appengine.Enabled {
		ProvidersData.Enable = "appengine"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Appengine.Enabled) + `
	primaryAccount: ` + providersRef.Appengine.PrimaryAccount +
		GetAppEngineAccounts(providersRef)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.AppEngine = str

	return nil
}

func GetAppEngineAccounts(provider *providers.Providers) string {
	str := ""

	if nil != provider.Appengine.Accounts {
		str += `
		  accounts:`
		for _, account := range provider.Appengine.Accounts {
			str += `
		  - name: ` + account.Name + `
		    environment              : ` + account.Environment +
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) + `
		    project                  : ` + account.Project + `
		    jsonPath                 : ` + account.JsonPath + `
		    localRepositoryDirectory : ` + account.LocalRepositoryDirectory + `
		    gitHttpUsername          : ` + account.GitHttpUsername + `
		    gitHttpsPassword         : ` + account.GitHttpsPassword + `
		    githubOAuthAccessToken   : ` + account.GithubOAuthAccessToken + `
		    sshPrivateKeyFilePath    : ` + account.SshPrivateKeyFilePath + `
		    sshPrivateKeyPassphrase  : ` + account.SshPrivateKeyPassphrase + `
		    sshKnownHostsFilePath    : ` + account.SshKnownHostsFilePath + `
		    sshTrustUnknownHosts     : ` + strconv.FormatBool(account.SshTrustUnknownHosts) + `
		    gcloudReleaseTrack       : ` + account.GcloudReleaseTrack +
				strings.Replace(getProvidersStringArrayAppend(account.Services, "services", "- "), "\t", "   ", -1) +
				strings.Replace(getProvidersStringArrayAppend(account.Versions, "versions", "- "), "\t", "   ", -1) +
				strings.Replace(getProvidersStringArrayAppend(account.OmitServices, "omitServices", "- "), "\t", "   ", -1) +
				strings.Replace(getProvidersStringArrayAppend(account.OmitVersions, "omitVersions", "- "), "\t", "   ", -1) + `
		    cachingIntervalSeconds   : ` + strconv.FormatInt(int64(account.CachingIntervalSeconds), 10) + `
		    Permission: {}` //TODO + master.Permission
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}