package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/notifications"
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
		slack:
		  enabled: ` + strconv.FormatBool(notificationsReference.Slack.Enabled) + `
		  botName: ` + notificationsReference.Slack.BotName + `
		  token: ` + notificationsReference.Slack.Token + `
		  baseUrl: ` + notificationsReference.Slack.BaseUrl + `
		  forceUseIncomingWebhook: ` + strconv.FormatBool(notificationsReference.Slack.ForceUseIncomingWebhook)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetTwilio(notificationsReference *notifications.Notifications) string {
	str := ""

	if nil != notificationsReference.Twilio {
		str = `
		twilio:
		  enabled: ` + strconv.FormatBool(notificationsReference.Twilio.Enabled) + `
		  account: ` + notificationsReference.Twilio.Account + `
		  baseUrl: ` + notificationsReference.Twilio.BaseUrl + `
		  from: '` + notificationsReference.Twilio.From + `'
		  token: ` + strconv.FormatBool(notificationsReference.Twilio.Token)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetGithubStatus(notificationsReference *notifications.Notifications) string {
	str := ""

	if nil != notificationsReference.GithubStatus {
		str = `
		github-status:
		  enabled: ` + strconv.FormatBool(notificationsReference.GithubStatus.Enabled) + `
		  token: ` + notificationsReference.GithubStatus.Token

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}
