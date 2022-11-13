package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/telemetry"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetTelemetry(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Telemetry {
		str = GetTelemetryConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Telemetry)
	} else {
		str += ` {}`
	}
	return str
}

func GetTelemetryConfig(telemetryReference *telemetry.Telemetry) string {

	str := helpers.PrintFmtBool(`enabled: `, telemetryReference.Enabled, 4, true) +
		helpers.PrintFmtStr(`endpoint: `, telemetryReference.Endpoint, 4, true) +
		helpers.PrintFmtStr(`instanceId: `, telemetryReference.InstanceId, 4, true) +
		helpers.PrintFmtInt(`connectionTimeoutMillis: `, telemetryReference.ConnectionTimeoutMillis, 4, true) +
		helpers.PrintFmtInt(`readTimeoutMillis: `, telemetryReference.ReadTimeoutMillis, 4, true)

	return str
}
