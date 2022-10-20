package migrate

import (
	"strings"
)

func GetPubsub(KustomizeData Kustomize) string {

	return GetPubsubConfig(KustomizeData)
}

func GetPubsubConfig(KustomizeData Kustomize) string {
	str := `
	enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].pubsub.Enable */ + `
	google:
	  enabled: false
	  pubsubType: GOOGLE
	  subscriptions: []
	  publishers: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
