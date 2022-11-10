package armory_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/armory"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetArmory(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Armory {
		str = GetDinghyArmory(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Armory) +
			GetDiagnosticsArmory(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Armory) +
			GetTerraformArmory(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Armory) +
			GetSecretsArmory(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Armory)
	}

	return str
}

func GetDinghyArmory(armoryReference *armory.Armory) string {
	str := ""

	if nil != armoryReference.Dinghy {
		//To avoid null pointers
		var notifiersSlack string
		if nil != armoryReference.Dinghy.Notifiers && nil != armoryReference.Dinghy.Notifiers.Slack {
			notifiersSlack = strconv.FormatBool(armoryReference.Dinghy.Notifiers.Slack.Enabled)
		}

		str = `
		dinghy:` +
			helpers.PrintFmtBool(`enabled: `, armoryReference.Dinghy.Enabled, 5, true) +
			helpers.PrintFmtStr(`templateOrg: `, armoryReference.Dinghy.TemplateOrg, 5, true) +
			helpers.PrintFmtStr(`templateRepo: `, armoryReference.Dinghy.TemplateRepo, 5, true) +
			helpers.PrintFmtStr(`githubToken: `, armoryReference.Dinghy.GithubToken, 5, true) +
			helpers.PrintFmtStr(`githubEndpoint: `, armoryReference.Dinghy.GithubEndpoint, 5, true) +
			helpers.PrintFmtBool(`repositoryRawdataProcessing: `, armoryReference.Dinghy.RepositoryRawdataProcessing, 5, true) +
			helpers.PrintFmtStr(`dinghyFilename: `, armoryReference.Dinghy.DinghyFilename, 5, true) +
			helpers.PrintFmtBool(`autoLockPipelines: `, armoryReference.Dinghy.AutoLockPipelines, 5, true) +
			helpers.PrintFmtStr(`fiatUser: `, armoryReference.Dinghy.FiatUser, 5, true) + `
		  notifiers:
		    slack:` +
			helpers.PrintFmtStr(`enabled: `, notifiersSlack, 7, true)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

/*func GetDinghyArmoryNotifiers(dinghy armory.Dinghy) string {
	str := ""

	if nil != dinghy.Notifiers {
		str += `
	  notifiers:
`
		for _, master := range jenkins.Notifiers {
			str += `
	    - slack:
	    enabled: true
	    channel: pbd-spinnaker-slack-testing``
		}
	} else {
		str += `	  notifiers: []`
	}

	str += `
		notifiers: []`
	return str
}*/

func GetDiagnosticsArmory(armoryReference *armory.Armory) string {
	str := ""

	if nil != armoryReference.Diagnostics {
		//To avoid null pointers
		var loggingEnabled string
		var loggingEndpoint string
		if nil != armoryReference.Diagnostics.Logging {
			loggingEnabled = strconv.FormatBool(armoryReference.Diagnostics.Logging.Enabled)
			loggingEndpoint = armoryReference.Diagnostics.Logging.Endpoint
		}

		str = `
		diagnostics:` +
			helpers.PrintFmtBool(`enabled: `, armoryReference.Diagnostics.Enabled, 5, true) +
			helpers.PrintFmtStr(`uuid: `, armoryReference.Diagnostics.Uuid, 5, true) + `
		  logging:` +
			helpers.PrintFmtStr(`enabled: `, loggingEnabled, 6, true) +
			helpers.PrintFmtStr(`endpoint: `, loggingEndpoint, 6, true)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetTerraformArmory(armoryReference *armory.Armory) string {
	str := ""

	if nil != armoryReference.Terraform {
		//To avoid null pointers
		var gitEnabled string
		var gitEndpoint string
		if nil != armoryReference.Terraform.Git {
			gitEnabled = strconv.FormatBool(armoryReference.Terraform.Git.Enabled)
			gitEndpoint = armoryReference.Terraform.Git.AccessToken
		}

		str = `
		terraform:` +
			helpers.PrintFmtBool(`enabled: `, armoryReference.Terraform.Enabled, 5, true) + `
		  git:` +
			helpers.PrintFmtStr(`enabled: `, gitEnabled, 6, true) +
			helpers.PrintFmtStr(`endpoint: `, gitEndpoint, 6, true)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetSecretsArmory(armoryReference *armory.Armory) string {
	str := ""

	//TODO Missing proto
	// if nil != armoryReference.Secrets {
	// 	str = `
	// 	secrets:
	// 	  enabled: ` + strconv.FormatBool(armoryReference.Secrets.Enabled) + `
	// 	  git: ` + armoryReference.Secrets.Path
	// 	str = strings.Replace(str, "\t", "    ", -1)
	// }

	str = `
	secrets:
	  vault:
	    enabled: ` + "false" /*armoryReference.Secrets.Vault.Enable */ + `
	    path: kubernetes`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
