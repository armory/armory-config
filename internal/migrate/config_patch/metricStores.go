package config_patch

import (
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

	//TODO check if api_key or api-key, also add_source_metalabels, credentials_path
	if nil != metricsReference.Datadog {
		str = `
		datadog:` +
			helpers.PrintFmtBool(`enabled: `, metricsReference.Datadog.Enabled, 5, true) +
			helpers.PrintFmtStr(`api_key: `, metricsReference.Datadog.ApiKey, 5, true) +
			helpers.PrintFmtStr(`app_key: `, metricsReference.Datadog.AppKey, 5, true)

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
		prometheus:` +
			helpers.PrintFmtBool(`enabled: `, metricsReference.Prometheus.Enabled, 5, true) +
			helpers.PrintFmtStr(`push-gateway: `, metricsReference.Prometheus.PushGateway, 5, true) +
			helpers.PrintFmtBool(`add_source_metalabels: `, metricsReference.Prometheus.AddSourceMetalabels, 5, true)
	}
	return str
}

func GetStackdriverMetrics(metricsReference *metricStores.MetricStores) string {
	str := ""

	if nil != metricsReference.Stackdriver {
		str = `
		stackdriver:` +
			helpers.PrintFmtBool(`enabled: `, metricsReference.Stackdriver.Enabled, 5, true) +
			helpers.PrintFmtStr(`project: `, metricsReference.Stackdriver.Project, 5, true) +
			helpers.PrintFmtStr(`zone: `, metricsReference.Stackdriver.Zone, 5, true) +
			helpers.PrintFmtStr(`credentials_path: `, metricsReference.Stackdriver.CredentialsPath, 5, true)
	}
	return str
}

func GetNewrelicMetrics(metricsReference *metricStores.MetricStores) string {
	str := ""

	if nil != metricsReference.Newrelic {
		str = `
		newrelic:` +
			helpers.PrintFmtBool(`enabled: `, metricsReference.Newrelic.Enabled, 5, true) +
			helpers.PrintFmtStr(`insert-key: `, metricsReference.Newrelic.InsertKey, 5, true) +
			helpers.PrintFmtStr(`host: `, metricsReference.Newrelic.Host, 5, true)

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
	return helpers.PrintFmtInt(`period: `, metricsReference.Period, 4, true) +
		helpers.PrintFmtBool(`enabled: `, metricsReference.Enabled, 4, true)
}
