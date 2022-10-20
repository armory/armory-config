package migrate

import (
	"strings"
)

func GetCanary(KustomizeData Kustomize) string {

	return GetCanaryConfig(KustomizeData)
}

func GetCanaryConfig(KustomizeData Kustomize) string {
	str := `
	enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].pubsub.Enable */ + `
	serviceIntegrations:
	- name: google
	  enabled: false
	  accounts: []
	  gcsEnabled: false
	  stackdriverEnabled: false
	- name: prometheus
	  enabled: false
	  accounts: []
	- name: datadog
	  enabled: false
	  accounts: []
	- name: signalfx
	  enabled: false
	  accounts: []
	- name: aws
	  enabled: false
	  accounts: []
	  s3Enabled: false
	- name: newrelic
	  enabled: false
	  accounts: []
	reduxLoggerEnabled: true
	defaultJudge: NetflixACAJudge-v1.0
	stagesEnabled: true
	templatesEnabled: true
	showAllConfigsEnabled: true`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
