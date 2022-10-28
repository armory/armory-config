package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
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
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") + `
		      dockerRegistries: []` /*+ TODO account.KubeDockerRegistries +*/ + `
		      providerVersion: ` + account.ProviderVersion + `
		      configureImagePullSecrets: ` + strconv.FormatBool(account.ConfigureImagePullSecrets) + `
		      serviceAccount: ` + strconv.FormatBool(account.ServiceAccount) + `
		      cacheThreads: ` + strconv.FormatInt(int64(account.CacheThreads), 10) +
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
				strings.Replace(getProvidersStringArray(ResourceEndpointKindOmitExpressions, "omitKindExpressions"), "\t", "     ", -1) + `
		      permission: {}` //TODO + account.Permission`
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}