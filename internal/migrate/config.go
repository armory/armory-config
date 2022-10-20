package migrate

import (
	"strings"
)

func GetConfigData(KustomizeData Kustomize) string {

	return GetGeneralConfig(KustomizeData) +
		GetSpecificConfig(KustomizeData)
}

func GetGeneralConfig(KustomizeData Kustomize) string {
	str := `
# === General Config ===
	  version: ` + "2.0" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Version*/ + `  # the version of Spinnaker to be deployed
	  timezone: ` + "America/Los_Angeles" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Timezone*/ + `
`

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetSpecificConfig(KustomizeData Kustomize) string {
	str := `
# === Persistent Storage ===
	  persistentStorage:` +
		GetPersistentStorage(KustomizeData) + `

# === Metric Stores ===
	  metricStores:` +
		GetMetricStores(KustomizeData) + `

# === Notifications ===
	  notifications:` +
		GetNotifications(KustomizeData) + `

# === Ci ===
	  ci:` +
		GetCi(KustomizeData) + `

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

# === Stats ===
	  stats:` +
		GetStats(KustomizeData)

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
