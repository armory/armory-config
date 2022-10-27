package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/stats"
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
	str := `
	enabled: ` + strconv.FormatBool(statsReference.Enabled) + `
	endpoint: ` + statsReference.Endpoint + `
	instanceId: ` + statsReference.InstanceId +
		getDeploymentMethodStats(statsReference) + `
	connectionTimeoutMillis: ` + strconv.FormatInt(int64(statsReference.ConnectionTimeoutMillis), 10) + `
	readTimeoutMillis: ` + strconv.FormatInt(int64(statsReference.ReadTimeoutMillis), 10)

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func getDeploymentMethodStats(deploymentStats *stats.Stats) string {
	str := ""
	str = `
	deploymentMethod: {}`
	//TODO

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
