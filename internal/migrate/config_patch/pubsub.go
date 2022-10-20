package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetPubsub(KustomizeData structs.Kustomize) string {

	return GetPubsubConfig(KustomizeData)
}

func GetPubsubConfig(KustomizeData structs.Kustomize) string {
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
