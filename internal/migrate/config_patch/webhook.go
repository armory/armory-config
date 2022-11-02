package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/webhook"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetWebhook(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Webhook {
		str = GetWebhookData(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Webhook)
	}
	return str
}

func GetWebhookData(webhookReference *webhook.Webhook) string {
	str := ""

	if nil != webhookReference.Trust {
		str += `
		trust:
		  enabled: ` + strconv.FormatBool(webhookReference.Trust.Enabled) + `
		  trustStore: ` + webhookReference.Trust.TrustStore + `
		  trustStorePassword: ` + webhookReference.Trust.TrustStorePassword
	} else {
		str += `
		trust: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)

	return str
}
