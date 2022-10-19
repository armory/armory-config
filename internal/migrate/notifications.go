package migrate

import "strings"

func GetNotifications(KustomizeData Kustomize) string {

	return GetSlack(KustomizeData) +
		GetTwilio(KustomizeData) +
		GetGithubStatus(KustomizeData)
}

func GetSlack(KustomizeData Kustomize) string {
	str := `
	slack:
	  enabled: ` + "true" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Notifications.Slack.Enable*/ + `
	  botName: spinnaker_bot
	  token: XXXX`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetTwilio(KustomizeData Kustomize) string {
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

func GetGithubStatus(KustomizeData Kustomize) string {
	str := `
	github-status:
	  enabled: false
	  token: XXXX`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
