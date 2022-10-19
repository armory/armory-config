package migrate

import "strings"

func GetMetricStores(KustomizeData Kustomize) string {

	return GetDatadog(KustomizeData) +
		GetPrometheus(KustomizeData) +
		GetStackdriver(KustomizeData) +
		GetNewrelic(KustomizeData) +
		GetMetricStoresGen(KustomizeData)
}

func GetDatadog(KustomizeData Kustomize) string {
	str := `
	datadog:
	  enabled: false
	  api_key: encrypted:k8s!n:spin-secrets!k:datadog-apikey
	  app_key: encrypted:k8s!n:spin-secrets!k:datadog-appkey
`

	/*if nil != KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog.Tags {
		str += `	  tags: #  Your datadog custom tags. Please delimit the KVP with colons i.e. app:test env:dev
		- env:dev`
	} else {
		str += `	  tags: []`
	}*/

	str += `	  tags: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetPrometheus(KustomizeData Kustomize) string {
	str := `
	prometheus:
	  enabled: false
	  add_source_metalabels: true`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.Prometheus
}

func GetStackdriver(KustomizeData Kustomize) string {
	str := `
	stackdriver:
	  enabled: false
	  project: google-project-id                                     # (Required). The project Spinnaker’s metrics should be published to.
	  zone: us-east1-b                                               # (Required). The zone Spinnaker’s metrics should be associated with.
	  credentials_path: encryptedFile:k8s!n:spin-secrets!k:gcp-json  # (Secret). A path to a Google JSON service account that has permission to publish metrics.`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver
}

func GetNewrelic(KustomizeData Kustomize) string {
	str := `
	newrelic:
	  tags: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic
}

func GetMetricStoresGen(KustomizeData Kustomize) string {
	str := `
	period: 30` + /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.Period + */ `
	enabled: false` /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.enabled + */

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
