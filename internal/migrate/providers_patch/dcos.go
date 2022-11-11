package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetDcos(providersRef *providers.Providers) error {

	if nil != providersRef.Dcos {
		if providersRef.Dcos.Enabled {
			ProvidersData.Enable = "dcos"
		}

		str := helpers.PrintFmtBool(`enabled: `, providersRef.Dcos.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.Dcos.PrimaryAccount, 5, true) +
			GetDcosAccounts(providersRef.Dcos.Accounts) +
			GetDcosMainClusters(providersRef.Dcos.Clusters)

		str = strings.Replace(str, "\t", "          ", -1)
		ProvidersData.Dcos = str
	} else {
		ProvidersData.Dcos = " {}"
		// 	return fmt.Errorf("Dcos value is null")
	}

	return nil
}

func GetDcosAccounts(accounts []*providers.DcosAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
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
			str += helpers.PrintFmtStr(`- name: `, Cluster.Name, 8, true) +
				helpers.PrintFmtStrApostrophe(`uid: `, Cluster.Uid, 9, true) +
				helpers.PrintFmtStr(`password: `, Cluster.Password, 9, true) +
				helpers.PrintFmtStr(`serviceKeyFile: `, Cluster.ServiceKeyFile, 9, true)
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

			str += helpers.PrintFmtStr(`- name: `, Cluster.Name, 5, true) +
				helpers.PrintFmtStr(`dcosUrl: `, Cluster.DcosUrl, 6, true) +
				helpers.PrintFmtStr(`caCertFile: `, Cluster.CaCertFile, 6, true) +
				helpers.PrintFmtBool(`insecureSkipTlsVerify: `, Cluster.InsecureSkipTlsVerify, 6, true) + `
		    loadBalancer:` +
				helpers.PrintFmtStr(`image: `, clusterLbImage, 7, true) +
				helpers.PrintFmtStr(`serviceAccountSecret: `, clusterLbSss, 7, true)
		}
	} else {
		str += `
		  clusters: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
