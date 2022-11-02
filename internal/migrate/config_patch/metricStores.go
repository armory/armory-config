package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/metricStores"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetMetricStores(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores {
		str = GetDatadogMetrics(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores) +
			GetPrometheusMetrics(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores) +
			GetStackdriverMetrics(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores) +
			GetNewrelicMetrics(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores) +
			GetMetricStoresGenMetrics(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].MetricStores)
	}
	return str
}

func GetDatadogMetrics(metricsReference *metricStores.MetricStores) string {
	str := ""

	if nil != metricsReference.Datadog {
		str = `
		datadog:
		  enabled: ` + strconv.FormatBool(metricsReference.Datadog.Enabled) + `
		  api_key: ` + metricsReference.Datadog.ApiKey + `
		  app_key: ` + metricsReference.Datadog.AppKey

		if nil != metricsReference.Datadog.Tags {
			str += `
		  tags: #  Your datadog custom tags. Please delimit the KVP with colons i.e. app:test env:dev`
			for _, b := range metricsReference.Datadog.Tags {
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

func GetPrometheusMetrics(metricsReference *metricStores.MetricStores) string {
	str := ""

	if nil != metricsReference.Prometheus {
		str = `
		prometheus:
		  enabled: ` + strconv.FormatBool(metricsReference.Prometheus.Enabled) + `
		  push-gateway: ` + metricsReference.Prometheus.PushGateway + `
		  add_source_metalabels: ` + strconv.FormatBool(metricsReference.Prometheus.AddSourceMetalabels)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetStackdriverMetrics(metricsReference *metricStores.MetricStores) string {
	str := ""

	if nil != metricsReference.Stackdriver {
		str = `
		stackdriver:
		  enabled: ` + strconv.FormatBool(metricsReference.Stackdriver.Enabled) + `
		  project: ` + metricsReference.Stackdriver.Project + `                              # (Required). The project Spinnaker’s metricsReference should be published to.
		  zone: ` + metricsReference.Stackdriver.Zone + `                                    # (Required). The zone Spinnaker’s metricsReference should be associated with.
		  credentials_path: ` + metricsReference.Stackdriver.CredentialsPath + `             # (Secret). A path to a Google JSON service account that has permission to publish metricsReference.`

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetNewrelicMetrics(metricsReference *metricStores.MetricStores) string {
	str := ""

	if nil != metricsReference.Newrelic {
		str = `
		newrelic:
		  enabled: ` + strconv.FormatBool(metricsReference.Newrelic.Enabled) + `
		  insert-key: ` + metricsReference.Newrelic.InsertKey + `
		  host: ` + metricsReference.Newrelic.Host

		if nil != metricsReference.Newrelic &&
			nil != metricsReference.Newrelic.Tags {
			str += `
		  tags: #  Your Newrelic custom tags.`
			for _, b := range metricsReference.Newrelic.Tags {
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

func GetMetricStoresGenMetrics(metricsReference *metricStores.MetricStores) string {

	str := `
	period: ` + helpers.IntToString(metricsReference.Period) + `
	enabled: ` + strconv.FormatBool(metricsReference.Enabled)

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
