package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetKubernetes(providersRef *providers.Providers) error {

	// if nil != providersRef.Kubernetes {
	// 	return fmt.Errorf("Kubernetes value is null")
	// }

	if providersRef.Kubernetes.Enabled {
		ProvidersData.Enable = "kubernetes"
	}

	str := helpers.PrintFmtBool(`enabled: `, providersRef.Kubernetes.Enabled, 5, true) +
		helpers.PrintFmtStr(`primaryAccount: `, providersRef.Kubernetes.PrimaryAccount, 5, true) +
		GetKubernetesAccounts(providersRef.Kubernetes.Accounts)

	ProvidersData.Kubernetes = str

	return nil
}

func GetKubernetesAccounts(accounts []*providers.KubernetesAcc) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			//To avoid null pointers
			var ResourceEndpointKindExpressions []string
			var ResourceEndpointKindOmitExpressions []string
			if nil != account.RawResourcesEndpointConfig {
				ResourceEndpointKindExpressions = account.RawResourcesEndpointConfig.KindExpressions
				ResourceEndpointKindOmitExpressions = account.RawResourcesEndpointConfig.OmitKindExpressions
			}

			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				getDockerRegistries(account.DockerRegistries) +
				helpers.PrintFmtStr(`context: `, account.Context, 7, true) +
				helpers.PrintFmtStr(`providerVersion: `, account.ProviderVersion, 7, true) +
				helpers.PrintFmtBool(`configureImagePullSecrets: `, account.ConfigureImagePullSecrets, 7, true) +
				helpers.PrintFmtBool(`serviceAccount: `, account.ServiceAccount, 7, true) +
				helpers.PrintFmtInt(`cacheThreads: `, account.CacheThreads, 7, true) +
				getProvidersStringArrayAppend(account.Namespaces, "namespaces", "- ") +
				helpers.PrintFmtStr(`kubeconfigFile: `, account.KubeconfigFile, 7, true) +
				getProvidersStringArrayAppend(account.OmitNamespaces, "omitNamespaces", "- ") +
				getProvidersStringArrayAppend(account.Kinds, "kinds", "- ") +
				getProvidersStringArrayAppend(account.OmitKinds, "omitKinds", "- ") +
				getProvidersStringArrayAppend(account.CustomResources, "customResources", "- ") +
				getProvidersStringArrayAppend(account.CachingPolicies, "cachingPolicies", "- ") +
				getProvidersStringArrayAppend(account.OauthScopes, "oauthScopes", "- ") +
				helpers.PrintFmtBool(`onlySpinnakerManaged: `, account.OnlySpinnakerManaged, 7, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
				helpers.PrintFmtBool(`checkPermissionsOnStartup: `, account.CheckPermissionsOnStartup, 7, true) +
				helpers.PrintFmtBool(`liveManifestCalls: `, account.LiveManifestCalls, 7, true) + `
		      rawResourcesEndpointConfig:` +
				strings.Replace(getProvidersStringArray(ResourceEndpointKindExpressions, "kindExpressions"), "\t", "     ", -1) +
				strings.Replace(getProvidersStringArray(ResourceEndpointKindOmitExpressions, "omitKindExpressions"), "\t", "     ", -1)
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getDockerRegistries(registries []*providers.KubeDockerRegistries) string {
	str := ""

	if nil != registries {
		str += `
		      dockerRegistries:`
		for _, account := range registries {
			str += helpers.PrintFmtStr(`- accountName: `, account.AccountName, 8, true) +
				strings.Replace(getProvidersStringArrayAppend(account.Namespaces, "Namespaces", "- "), "\t", "      ", -1)
		}
	} else {
		str += `
		      dockerRegistries: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
