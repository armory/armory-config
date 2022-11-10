package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/stats"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetStats(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats {
		str = GetStatsConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats)
	}
	return str
}

func GetStatsConfig(statsReference *stats.Stats) string {
	str := helpers.PrintFmtBool(`enabled: `, statsReference.Enabled, 4, true) +
		helpers.PrintFmtStr(`endpoint: `, statsReference.Endpoint, 4, true) +
		helpers.PrintFmtStr(`instanceId: `, statsReference.InstanceId, 4, true) +
		getDeploymentMethodStats(statsReference) +
		helpers.PrintFmtInt(`connectionTimeoutMillis: `, statsReference.ConnectionTimeoutMillis, 4, true) +
		helpers.PrintFmtInt(`readTimeoutMillis: `, statsReference.ReadTimeoutMillis, 4, true)

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func getDeploymentMethodStats(deploymentStats *stats.Stats) string {
	str := ""
	str = `
	deploymentMethod: {}`
	//TODO Missing proto

	/*if nil != deploymentStats.DeploymentMethod {
		str += `
	deploymentMethod:` +
		deploymentStats.
	} else {
		str += `
	deploymentMethod: {}`
	}
	return str*/
	return str
}
