package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetStats(KustomizeData structs.Kustomize) string {

	return GetStatsConfig(KustomizeData)
}

func GetStatsConfig(KustomizeData structs.Kustomize) string {
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
