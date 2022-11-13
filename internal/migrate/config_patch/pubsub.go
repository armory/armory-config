package config_patch

import (
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
	str := helpers.PrintFmtBool(`enabled: `, pubsubReference.Enabled, 4, true)

	if nil != pubsubReference.Google {
		str += `
		google:` +
			helpers.PrintFmtBool(`enabled: `, pubsubReference.Google.Enabled, 5, true) +
			helpers.PrintFmtStr(`pubsubType: `, pubsubReference.Google.PubsubType, 5, true) +
			GetGoogleSubscriptions(pubsubReference.Google) +
			GetGooglePublishers(pubsubReference.Google)
	}
	return str
}

func GetGoogleSubscriptions(google *pubsub.Google) string {
	str := ""

	if nil != google.Subscriptions {
		str += `
		  subscriptions:`
		for _, account := range google.Subscriptions {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 5, true) +
				helpers.PrintFmtStr(`project: `, account.Project, 6, true) +
				helpers.PrintFmtStr(`subscriptionName: `, account.SubscriptionName, 6, true) +
				helpers.PrintFmtStr(`jsonPath: `, account.JsonPath, 6, true) +
				helpers.PrintFmtStr(`templatePath: `, account.TemplatePath, 6, true) +
				helpers.PrintFmtInt(`ackDeadlineSeconds: `, account.AckDeadlineSeconds, 6, true) +
				helpers.PrintFmtStr(`messageFormat: `, account.MessageFormat, 6, true)
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
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`project: `, account.Project, 7, true) +
				helpers.PrintFmtStr(`topicName: `, account.TopicName, 7, true) +
				helpers.PrintFmtStr(`jsonPath: `, account.JsonPath, 7, true) +
				helpers.PrintFmtStr(`content: `, account.Content, 7, true)
		}
	} else {
		str += `
		  publishers: []`
	}

	return str
}
