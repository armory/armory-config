package armory_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/armory"
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
		dinghy:
		  enabled: ` + strconv.FormatBool(armoryReference.Dinghy.Enabled) + `
		  templateOrg: ` + armoryReference.Dinghy.TemplateOrg + `
		  templateRepo: ` + armoryReference.Dinghy.TemplateRepo + `
		  githubToken: ` + armoryReference.Dinghy.GithubToken + `
		  githubEndpoint: ` + armoryReference.Dinghy.GithubEndpoint + `
		  repositoryRawdataProcessing: ` + strconv.FormatBool(armoryReference.Dinghy.RepositoryRawdataProcessing) + `
		  dinghyFilename: ` + armoryReference.Dinghy.DinghyFilename + `
		  autoLockPipelines: ` + strconv.FormatBool(armoryReference.Dinghy.AutoLockPipelines) + `
		  fiatUser: ` + armoryReference.Dinghy.FiatUser + `
		  notifiers:
		    slack:
		      enabled: ` + notifiersSlack
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
		diagnostics:
		  enabled: ` + strconv.FormatBool(armoryReference.Diagnostics.Enabled) + `
		  uuid: ` + armoryReference.Diagnostics.Uuid + `
		  logging:
		    enabled: ` + loggingEnabled + `
		    endpoint: ` + loggingEndpoint
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
		terraform:
		  enabled: ` + strconv.FormatBool(armoryReference.Terraform.Enabled) + `
		  git:
		    enabled: ` + gitEnabled + `
		    endpoint: ` + gitEndpoint
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
