package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/pubsub"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetPubsub(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub {
		str = GetPubsubConfig(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub)
	}
	return str
}

func GetPubsubConfig(pubsubReference *pubsub.Pubsub) string {
	str := `
	    enabled: ` + strconv.FormatBool(pubsubReference.Enabled)

	if nil != pubsubReference.Google {
		str += `
		google:
		  enabled: ` + strconv.FormatBool(pubsubReference.Google.Enabled) + `
		  pubsubType: ` + pubsubReference.Google.PubsubType +
			GetGoogleSubscriptions(pubsubReference.Google) +
			GetGooglePublishers(pubsubReference.Google)

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
		      ackDeadlineSeconds: ` + helpers.IntToString(account.AckDeadlineSeconds) + `
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
