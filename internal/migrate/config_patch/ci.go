package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetCi(KustomizeData structs.Kustomize) string {

	return GetJenkinsCi(KustomizeData) +
		GetTravisCi(KustomizeData) +
		GetWerckerCi(KustomizeData) +
		GetConcourseCi(KustomizeData) +
		GetGcbCi(KustomizeData) +
		GetCodebuildCi(KustomizeData)
}

func GetJenkinsCi(KustomizeData structs.Kustomize) string {
	str := `
	jenkins:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Jenkins.Enable +
	GetJenkinsCiMasters(KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Jenkins) */

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

/*func GetJenkinsCiMasters(jenkins ci.Jenkins) string {
	str := ""

	if nil != jenkins.Masters {
		str += `
	  masters:
`
		for _, master := range jenkins.Masters {
			str += `
	    - name:` + master.Name + `
	    - Address:` + master.Address + `
	    - Username:` + master.Username + `
	    - Password:` + master.Password + `
	    - Csrf:` + strconv.FormatBool(master.Csrf) + `
		- Permission: {}` //+ master.Permission
		}
	} else {
		str += `	  masters: []`
	}

	str += `
		masters: []`
	return str
}*/

func GetTravisCi(KustomizeData structs.Kustomize) string {
	str := `
	travis:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Travis.Enable*/ + `
	  masters: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetWerckerCi(KustomizeData structs.Kustomize) string {
	str := `
	wercker:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Wercker.Enable*/ + `
	  masters: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetConcourseCi(KustomizeData structs.Kustomize) string {
	str := `
	concourse:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Concourse.Enable*/ + `
	  masters: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGcbCi(KustomizeData structs.Kustomize) string {
	str := `
	gcb:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Gcb.Enable*/ + `
	  masters: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetCodebuildCi(KustomizeData structs.Kustomize) string {
	str := `
	codebuild:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Ci.Codebuild.Enable*/ + `
	  masters: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
