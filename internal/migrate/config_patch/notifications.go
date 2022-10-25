package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetNotifications(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications {
		str = GetSlack(KustomizeData) +
			GetTwilio(KustomizeData) +
			GetGithubStatus(KustomizeData)
	}
	return str
}

func GetSlack(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Slack {
		str = `
		slack:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Slack.Enabled) + `
		  botName: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Slack.BotName + `
		  token: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Slack.Token + `
		  baseUrl: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Slack.BaseUrl + `
		  forceUseIncomingWebhook: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Slack.ForceUseIncomingWebhook)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetTwilio(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Twilio {
		str = `
		twilio:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Twilio.Enabled) + `
		  account: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Twilio.Account + `
		  baseUrl: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Twilio.BaseUrl + `
		  from: '` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Twilio.From + `'
		  token: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.Twilio.Token)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetGithubStatus(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.GithubStatus {
		str = `
		github-status:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.GithubStatus.Enabled) + `
		  token: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Notifications.GithubStatus.Token

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}
