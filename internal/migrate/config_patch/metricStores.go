package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetMetricStores(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores {
		str = GetDatadogMetrics(KustomizeData) +
			GetPrometheusMetrics(KustomizeData) +
			GetStackdriverMetrics(KustomizeData) +
			GetNewrelicMetrics(KustomizeData) +
			GetMetricStoresGenMetrics(KustomizeData)
	}
	return str
}

func GetDatadogMetrics(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog {
		str = `
		datadog:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog.Enabled) + `
		  api_key: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog.ApiKey + `
		  app_key: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog.AppKey

		if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog.Tags {
			str += `
		  tags: #  Your datadog custom tags. Please delimit the KVP with colons i.e. app:test env:dev`
			for _, b := range KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Datadog.Tags {
				str += "\n" + `		    - ` + b
				// - env:dev    # Example
			}
		} else {
			str += `
		  tags: []`
		}

		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetPrometheusMetrics(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Prometheus {
		str = `
		prometheus:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Prometheus.Enabled) + `
		  push-gateway: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Prometheus.PushGateway + `
		  add_source_metalabels: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Prometheus.AddSourceMetalabels)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetStackdriverMetrics(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver {
		str = `
		stackdriver:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver.Enabled) + `
		  project: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver.Project + `                              # (Required). The project Spinnaker’s metrics should be published to.
		  zone: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver.Zone + `                                    # (Required). The zone Spinnaker’s metrics should be associated with.
		  credentials_path: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver.CredentialsPath + `             # (Secret). A path to a Google JSON service account that has permission to publish metrics.`

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].MetricStores.Stackdriver
}

func GetNewrelicMetrics(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic {
		str = `
		newrelic:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic.Enabled) + `
		  insert-key: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic.InsertKey + `
		  host: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic.Host

		if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic &&
			nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic.Tags {
			str += `
		  tags: #  Your Newrelic custom tags.`
			for _, b := range KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Newrelic.Tags {
				str += "\n" + `		    - ` + b
			}
		} else {
			str += `
		  tags: []`
		}
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetMetricStoresGenMetrics(KustomizeData structs.Kustomize) string {

	str := `
	period: ` + strconv.FormatInt(int64(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Period), 10) + `
	enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores.Enabled)

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
