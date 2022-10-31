package config_patch

import (
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetTelemetry(KustomizeData structs.Kustomize) string {
	str := ""

	//TODO AAA missing proto
	// if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Telemetry {
	// 	str = GetTelemetryConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Telemetry)
	// }
	return str
}

// func GetTelemetryConfig(telemetryReference *telemetry.Telemetry) string {
// 	str := `
// 	enabled: ` + strconv.FormatBool(telemetryReference.Enabled) + `
// 	endpoint: ` + telemetryReference.Endpoint + `
// 	instanceId: ` + telemetryReference.InstanceId /*+
// 	getDeploymentMethodStats(telemetryReference) */+`
// 	connectionTimeoutMillis: ` + strconv.FormatInt(int64(telemetryReference.ConnectionTimeoutMillis), 10) + `
// 	readTimeoutMillis: ` + strconv.FormatInt(int64(telemetryReference.ReadTimeoutMillis), 10)

// 	str = strings.Replace(str, "\t", "        ", -1)
// 	return str
// }
