package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetNotifications(KustomizeData structs.Kustomize) string {

	return GetSlack(KustomizeData) +
		GetTwilio(KustomizeData) +
		GetGithubStatus(KustomizeData)
}

func GetSlack(KustomizeData structs.Kustomize) string {
	str := `
	slack:
	  enabled: ` + "true" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Notifications.Slack.Enable*/ + `
	  botName: spinnaker_bot
	  token: XXXX`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetTwilio(KustomizeData structs.Kustomize) string {
	str := `
	twilio:
	  enabled: true
	  account: XXXX
	  baseUrl: https://api.twilio.com/
	  from: '1234-567-8910'
	  token: encrypted:k8s!n:spin-secrets!k:twilio-token`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGithubStatus(KustomizeData structs.Kustomize) string {
	str := `
	github-status:
	  enabled: false
	  token: XXXX`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
