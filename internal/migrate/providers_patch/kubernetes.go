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
			str += `
		    - name: ` + account.Name +
				getProvidersStringArray(account.RequiredGroupMembership, "requiredGroupMembership") + `
		      dockerRegistries: []` /*+ account.KubeDockerRegistries +*/ + `
		      providerVersion: ` + account.ProviderVersion + `
		      configureImagePullSecrets: ` + strconv.FormatBool(account.ConfigureImagePullSecrets) + `
		      serviceAccount: ` + strconv.FormatBool(account.ServiceAccount) + `
		      cacheThreads: ` + strconv.FormatInt(int64(account.CacheThreads), 10) +
				getProvidersStringArray(account.Namespaces, "namespaces") + `
		      kubeconfigFile: ` + account.KubeconfigFile +
				getProvidersStringArray(account.OmitNamespaces, "omitNamespaces") +
				getProvidersStringArray(account.Kinds, "kinds") +
				getProvidersStringArray(account.OmitKinds, "omitKinds") +
				getProvidersStringArray(account.CustomResources, "customResources") +
				getProvidersStringArray(account.CachingPolicies, "cachingPolicies") +
				getProvidersStringArray(account.OauthScopes, "oauthScopes") + `
		      onlySpinnakerManaged: ` + strconv.FormatBool(account.OnlySpinnakerManaged) + `
		      environment: ` + account.Environment + `
		      checkPermissionsOnStartup: ` + strconv.FormatBool(account.CheckPermissionsOnStartup) + `
		      liveManifestCalls: ` + strconv.FormatBool(account.LiveManifestCalls) + `
		      rawResourcesEndpointConfig:` +
				strings.Replace(getProvidersStringArray(account.RawResourcesEndpointConfig.KindExpressions, "kindExpressions"), "\t", "     ", -1) +
				strings.Replace(getProvidersStringArray(account.RawResourcesEndpointConfig.OmitKindExpressions, "omitKindExpressions"), "\t", "     ", -1) + `
		      permission: {}` //TODO + account.Permission`
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
