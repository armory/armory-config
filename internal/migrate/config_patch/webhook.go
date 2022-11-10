package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/webhook"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
		trust:` +
			helpers.PrintFmtBool(`enabled: `, webhookReference.Trust.Enabled, 5, true) +
			helpers.PrintFmtStr(`trustStore: `, webhookReference.Trust.TrustStore, 5, true) +
			helpers.PrintFmtStrApostrophe(`trustStorePassword: `, webhookReference.Trust.TrustStorePassword, 5, true)
	} else {
		str += `
		trust: []`
	}

	return str
}
