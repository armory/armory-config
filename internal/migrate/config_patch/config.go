package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetConfigData(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos] {
		str = GetGeneralConfig(KustomizeData) +
			GetSpecificConfig(KustomizeData)
	}
	return str
}

func GetGeneralConfig(KustomizeData structs.Kustomize) string {
	str := `
# === General Config ===
	  version: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Version + `  # the version of Spinnaker to be deployed
	  timezone: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Timezone + `
`

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetSpecificConfig(KustomizeData structs.Kustomize) string {
	str := `
# === Persistent Storage ===
	  persistentStorage:` +
		GetPersistentStorage(KustomizeData) + `

# === Features ===
	  features:` +
		GetFeatures(KustomizeData) + `

# === Metric Stores ===
	  metricStores:` +
		GetMetricStores(KustomizeData) + `

# === Notifications ===
	  notifications:` +
		GetNotifications(KustomizeData) + `

# === Ci ===
	  ci:` +
		GetCi(KustomizeData) + `

# === Repository ===
	  repository:` +
		GetRepository(KustomizeData) + `

# === Security ===
	  security:` +
		GetSecurity(KustomizeData) + `

# === Artifacts ===
	  artifacts:` +
		GetArtifacts(KustomizeData) + `

# === Pubsub ===
	  pubsub:` +
		GetPubsub(KustomizeData) + `

# === Canary ===
	  canary:` +
		GetCanary(KustomizeData) + `

# === Spinnaker ===
	  spinnaker:` +
		GetSpinnaker(KustomizeData) + `

# === Webhook ===
	  webhook:` +
		GetWebhook(KustomizeData) + `

# === Stats ===
	  stats:` +
		GetStats(KustomizeData) + `

# === Telemetry ===
	  telemetry:` +
		GetTelemetry(KustomizeData)

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
