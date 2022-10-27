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
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
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
				getProvidersStringArray(account.Services, "services") +
				getProvidersStringArray(account.Versions, "versions") +
				getProvidersStringArray(account.OmitServices, "omitServices") +
				getProvidersStringArray(account.OmitVersions, "omitVersions") + `
		      cachingIntervalSeconds   : ` + strconv.FormatInt(int64(account.CachingIntervalSeconds), 10) + `
		      - Permission: {}` //TODO + master.Permission
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "      ", -1)

	return str
}
