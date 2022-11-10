package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/notifications"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetNotifications(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications {
		str = GetSlack(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications) +
			GetTwilio(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications) +
			GetGithubStatus(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications)
	}
	return str
}

func GetSlack(notificationsReference *notifications.Notifications) string {
	str := ""

	if nil != notificationsReference.Slack {
		str = `
		slack:` +
			helpers.PrintFmtBool(`enabled: `, notificationsReference.Slack.Enabled, 5, true) +
			helpers.PrintFmtStr(`botName: `, notificationsReference.Slack.BotName, 5, true) +
			helpers.PrintFmtStr(`token: `, notificationsReference.Slack.Token, 5, true) +
			helpers.PrintFmtStr(`baseUrl: `, notificationsReference.Slack.BaseUrl, 5, true) +
			helpers.PrintFmtBool(`forceUseIncomingWebhook: `, notificationsReference.Slack.ForceUseIncomingWebhook, 5, true)
	}
	return str
}

func GetTwilio(notificationsReference *notifications.Notifications) string {
	str := ""

	if nil != notificationsReference.Twilio {
		str = `
		twilio:` +
			helpers.PrintFmtBool(`enabled: `, notificationsReference.Twilio.Enabled, 5, true) +
			helpers.PrintFmtStr(`account: `, notificationsReference.Twilio.Account, 5, true) +
			helpers.PrintFmtStr(`baseUrl: `, notificationsReference.Twilio.BaseUrl, 5, true) +
			helpers.PrintFmtStrApostrophe(`from: `, notificationsReference.Twilio.From, 5, true) +
			helpers.PrintFmtBool(`token: `, notificationsReference.Twilio.Token, 5, true)
	}
	return str
}

func GetGithubStatus(notificationsReference *notifications.Notifications) string {
	str := ""

	if nil != notificationsReference.GithubStatus {
		str = `
		github-status:` +
			helpers.PrintFmtBool(`enabled: `, notificationsReference.GithubStatus.Enabled, 5, true) +
			helpers.PrintFmtStr(`token: `, notificationsReference.GithubStatus.Token, 5, true)

	}
	return str
}
