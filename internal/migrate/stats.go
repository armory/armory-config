package migrate

import (
	"strings"
)

func GetStats(KustomizeData Kustomize) string {

	return GetStatsConfig(KustomizeData)
}

func GetStatsConfig(KustomizeData Kustomize) string {
	str := `
	enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].pubsub.Enable */ + `
	endpoint: https://stats.spinnaker.io
	instanceId: XXX
	deploymentMethod: {}
	connectionTimeoutMillis: 3000
	readTimeoutMillis: 5000`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
