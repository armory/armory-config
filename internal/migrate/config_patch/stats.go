package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/stats"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetStats(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Canary {
		str = GetStatsConfig(KustomizeData)
	}
	return str
}

func GetStatsConfig(KustomizeData structs.Kustomize) string {
	str := `
	enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats.Enabled) + `
	endpoint: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats.Endpoint + `
	instanceId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats.InstanceId +
		getDeploymentMethodStats(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats) + `
	connectionTimeoutMillis: ` + strconv.FormatInt(int64(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats.ConnectionTimeoutMillis), 10) + `
	readTimeoutMillis: ` + strconv.FormatInt(int64(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Stats.ReadTimeoutMillis), 10)

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
