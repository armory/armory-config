package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetDcos(providersRef *providers.Providers) error {

	// if nil != providersRef.Dcos {
	// 	return fmt.Errorf("Dcos value is null")
	// }

	if providersRef.Dcos.Enabled {
		ProvidersData.Enable = "dcos"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Dcos.Enabled) + `
	primaryAccount: ` + providersRef.Dcos.PrimaryAccount +
		GetDcosAccounts(providersRef.Dcos.Accounts) +
		GetDcosMainClusters(providersRef.Dcos.Clusters)

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Dcos = str

	return nil
}

func GetDcosAccounts(accounts []*providers.DcosAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += `
		    - name: ` + account.Name + `
		      environment: ` + account.Environment +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				getDcosDockerRegistries(account.DockerRegistries) +
				getDcosAccountsClusters(account.Clusters)
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getDcosDockerRegistries(dockerRegistries []*providers.DcosDocker) string {
	str := ""

	if nil != dockerRegistries {
		str += `
		      dockerRegistries:`
		for _, dockerRegistry := range dockerRegistries {
			str += `
		        - accountName: ` + dockerRegistry.AccountName +
				getDcosNamespaces(dockerRegistry.Namespaces, "namespaces")
		}
	} else {
		str += `
		      dockerRegistries: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getDcosAccountsClusters(Clusters []*providers.DcosAccCluster) string {
	str := ""

	if nil != Clusters {
		str += `
		      clusters:`
		for _, Cluster := range Clusters {
			str += `
		        - name: ` + Cluster.Name + `
		          uid: ` + Cluster.Uid + `
		          password: ` + Cluster.Password + `
		          serviceKeyFile: ` + Cluster.ServiceKeyFile
		}
	} else {
		str += `
		      clusters: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getDcosNamespaces(stringArray []string, fieldName string) string {
	str := ""

	if nil != stringArray {
		str += `
		          ` + fieldName + `:`
		for _, stringValue := range stringArray {
			str += `
		            ` + stringValue
		}
	} else {
		str += `
		          ` + fieldName + `: []`
	}

	return str
}

func GetDcosMainClusters(Clusters []*providers.DcosCluster) string {
	str := ""

	if nil != Clusters {
		str += `
		  clusters:`
		for _, Cluster := range Clusters {
			clusterLbImage := ""
			clusterLbSss := ""
			if nil != Cluster.LoadBalancer {
				clusterLbImage = Cluster.LoadBalancer.Image
				clusterLbSss = Cluster.LoadBalancer.ServiceAccountSecret
			}

			str += `
		  - name: ` + Cluster.Name + `
		    dcosUrl: ` + Cluster.DcosUrl + `
		    caCertFile: ` + Cluster.CaCertFile + `
		    insecureSkipTlsVerify: ` + strconv.FormatBool(Cluster.InsecureSkipTlsVerify) + `
		    loadBalancer:
		      image: ` + clusterLbImage + `
		      serviceAccountSecret: ` + clusterLbSss
		}
	} else {
		str += `
		  clusters: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
