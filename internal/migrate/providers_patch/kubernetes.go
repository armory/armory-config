package providers_patch

import (
	"strconv"
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

	str := `enabled: ` + strconv.FormatBool(providersRef.Kubernetes.Enabled) + `
	primaryAccount: ` + providersRef.Kubernetes.PrimaryAccount +
		GetKubernetesAccounts(providersRef.Kubernetes.Accounts)

	str = strings.Replace(str, "\t", "          ", -1)
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

			str += `
		    - name: ` + account.Name +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				getDockerRegistries(account.DockerRegistries) + `
		      providerVersion: ` + account.ProviderVersion + `
		      configureImagePullSecrets: ` + strconv.FormatBool(account.ConfigureImagePullSecrets) + `
		      serviceAccount: ` + strconv.FormatBool(account.ServiceAccount) + `
		      cacheThreads: ` + helpers.IntToString(account.CacheThreads) +
				getProvidersStringArrayAppend(account.Namespaces, "namespaces", "- ") + `
		      kubeconfigFile: ` + account.KubeconfigFile +
				getProvidersStringArrayAppend(account.OmitNamespaces, "omitNamespaces", "- ") +
				getProvidersStringArrayAppend(account.Kinds, "kinds", "- ") +
				getProvidersStringArrayAppend(account.OmitKinds, "omitKinds", "- ") +
				getProvidersStringArrayAppend(account.CustomResources, "customResources", "- ") +
				getProvidersStringArrayAppend(account.CachingPolicies, "cachingPolicies", "- ") +
				getProvidersStringArrayAppend(account.OauthScopes, "oauthScopes", "- ") + `
		      onlySpinnakerManaged: ` + strconv.FormatBool(account.OnlySpinnakerManaged) + `
		      environment: ` + account.Environment + `
		      checkPermissionsOnStartup: ` + strconv.FormatBool(account.CheckPermissionsOnStartup) + `
		      liveManifestCalls: ` + strconv.FormatBool(account.LiveManifestCalls) + `
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
			str += `
		        - accountName: ` + account.AccountName +
				strings.Replace(getProvidersStringArrayAppend(account.Namespaces, "Namespaces", "- "), "\t", "      ", -1)
		}
	} else {
		str += `
		      dockerRegistries: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
