package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/pubsub"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetPubsub(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub {
		str = GetPubsubConfig(KustomizeData)
	}
	return str
}

func GetPubsubConfig(KustomizeData structs.Kustomize) string {
	str := `
	    enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Enabled)

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google {
		str += `
		google:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google.Enabled) + `
		  pubsubType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google.PubsubType +
			GetGoogleSubscriptions(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google) +
			GetGooglePublishers(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetGoogleSubscriptions(google *pubsub.Google) string {
	str := ""

	if nil != google.Subscriptions {
		str += `
		  subscriptions: `
		for _, account := range google.Subscriptions {
			str += `
		    - name: ` + account.Name + `
		      project: ` + account.Project + `
		      subscriptionName: ` + account.SubscriptionName + `
		      jsonPath: ` + account.JsonPath + `
		      templatePath: ` + account.TemplatePath + `
		      ackDeadlineSeconds: ` + strconv.FormatInt(int64(account.AckDeadlineSeconds), 10) + `
		      messageFormat: ` + account.MessageFormat
		}
	} else {
		str += `
		  subscriptions: []`
	}

	return str
}

func GetGooglePublishers(google *pubsub.Google) string {
	str := ""

	if nil != google.Publishers {
		str += `
		  publishers:`
		for _, account := range google.Publishers {
			str += `
		    - name: ` + account.Name + `
		      project: ` + account.Project + `
		      topicName: ` + account.TopicName + `
		      jsonPath: ` + account.JsonPath + `
		      content: ` + account.Content
		}
	} else {
		str += `
		  publishers: []`
	}

	return str
}
