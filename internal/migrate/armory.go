package migrate

import (
	"strings"
)

func GetArmory(KustomizeData Kustomize) string {

	return GetDinghyArmory(KustomizeData) +
		GetDiagnosticsArmory(KustomizeData) +
		GetTerraformArmory(KustomizeData) +
		GetSecretsArmory(KustomizeData)
}

func GetDinghyArmory(KustomizeData Kustomize) string {
	str := `
	dinghy:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Armory.Dinghy.Enable */ + `
	  templateOrg: cloudtools
	  templateRepo: dinghy-templates
	  githubToken: XXXX
	  githubEndpoint: https://github.dev.dou.com/api/v3
	  repositoryRawdataProcessing: false
	  dinghyFilename: dinghyfile
	  autoLockPipelines: true
	  fiatUser: svc-spinnaker-dinghy` /*+
	GetDinghyNotifiers(KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Armory.Dinghy)*/

	str = strings.Replace(str, "\t", "        ", -1)
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

func GetDiagnosticsArmory(KustomizeData Kustomize) string {
	str := `
	diagnostics:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Armory.Diagnostics.Enable */ + `
	  uuid: XXXX
	  logging:
	    enabled: false
	    endpoint: https://debug.armory.io/v1/logs`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetTerraformArmory(KustomizeData Kustomize) string {
	str := `
	terraform:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Armory.Terraform.Enable */ + `
	  uuid: XXXX
	  git:
	    enabled: false
	    accessToken: XXXX`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetSecretsArmory(KustomizeData Kustomize) string {
	str := `
	secrets:
	  vault:
	    enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Armory.Secrets.Vault.Enable */ + `
	    path: kubernetes`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
