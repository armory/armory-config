package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetAppEngineData(providersRef *providers.Providers) error {

	if nil != providersRef.Appengine {
		if providersRef.Appengine.Enabled {
			ProvidersData.Enable = "appengine"
		}

		str := helpers.PrintFmtBool(`enabled: `, providersRef.Appengine.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.Appengine.PrimaryAccount, 5, true) +
			GetAppEngineAccounts(providersRef)

		ProvidersData.AppEngine = str
	} else {
		ProvidersData.AppEngine = " {}"
		// 	return fmt.Errorf("Appengine value is null")
	}

	return nil
}

func GetAppEngineAccounts(provider *providers.Providers) string {
	str := ""

	if nil != provider.Appengine.Accounts {
		str += `
		  accounts:`
		for _, account := range provider.Appengine.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 5, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 6, true) +
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) +
				helpers.PrintFmtStr(`providerVersion: `, account.ProviderVersion, 6, true) +
				getPermissions(account.Permissions) +
				helpers.PrintFmtStr(`project: `, account.Project, 6, true) +
				helpers.PrintFmtStr(`jsonPath: `, account.JsonPath, 6, true) +
				helpers.PrintFmtStr(`localRepositoryDirectory: `, account.LocalRepositoryDirectory, 6, true) +
				helpers.PrintFmtStr(`gitHttpsUsername: `, account.GitHttpsUsername, 6, true) +
				helpers.PrintFmtStr(`gitHttpsPassword: `, account.GitHttpsPassword, 6, true) +
				helpers.PrintFmtStr(`githubOAuthAccessToken: `, account.GithubOAuthAccessToken, 6, true) +
				helpers.PrintFmtStr(`sshPrivateKeyFilePath: `, account.SshPrivateKeyFilePath, 6, true) +
				helpers.PrintFmtStr(`sshPrivateKeyPassphrase: `, account.SshPrivateKeyPassphrase, 6, true) +
				helpers.PrintFmtStr(`sshKnownHostsFilePath: `, account.SshKnownHostsFilePath, 6, true) +
				helpers.PrintFmtStr(`sshTrustUnknownHosts: `, strconv.FormatBool(account.SshTrustUnknownHosts), 6, true) +
				helpers.PrintFmtStr(`gcloudReleaseTrack: `, account.GcloudReleaseTrack, 6, true) +
				strings.Replace(getProvidersStringArrayAppend(account.Services, "services", "- "), "\t", "   ", -1) +
				strings.Replace(getProvidersStringArrayAppend(account.Versions, "versions", "- "), "\t", "   ", -1) +
				strings.Replace(getProvidersStringArrayAppend(account.OmitServices, "omitServices", "- "), "\t", "   ", -1) +
				strings.Replace(getProvidersStringArrayAppend(account.OmitVersions, "omitVersions", "- "), "\t", "   ", -1) +
				helpers.PrintFmtInt(`cachingIntervalSeconds: `, account.CachingIntervalSeconds, 6, true)
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}
