package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/ci"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci {
		str = GetJenkinsCi(KustomizeData) +
			GetTravisCi(KustomizeData) +
			GetWerckerCi(KustomizeData) +
			GetConcourseCi(KustomizeData) +
			GetGcbCi(KustomizeData) +
			GetCodebuildCi(KustomizeData)
	}
	return str
}

func GetJenkinsCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Jenkins {
		str = `
		jenkins:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Jenkins.Enabled) +
			GetJenkinsCiMasters(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Jenkins)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetJenkinsCiMasters(jenkins *ci.Jenkins) string {
	str := ""

	if nil != jenkins.Masters {
		str += `
		  masters:`
		for _, master := range jenkins.Masters {
			str += `
		    - name: ` + master.Name + `
		    - Address: ` + master.Address + `
		    - Username: ` + master.Username + `
		    - Password: ` + master.Password + `
		    - Csrf: ` + strconv.FormatBool(master.Csrf) + `
		    - Permission: {}` //TODO + master.Permission
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetTravisCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Travis {
		str = `
		travis:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Travis.Enabled) +
			GetTravisCiMasters(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Travis)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetTravisCiMasters(travis *ci.Travis) string {
	str := ""

	if nil != travis.Masters {
		str += `
		  masters:`
		for _, master := range travis.Masters {
			str += `
		    - name: ` + master.Name + `
		    - Address: ` + master.Address + `
		    - BaseUrl: ` + master.BaseUrl + `
		    - githubToken: ` + master.GithubToken + `
		    - numberOfRepositories: ` + strconv.FormatInt(int64(master.NumberOfRepositories), 10) + `
		    - filteredRepositories: []` /* TODO ARRAY master.FilteredRepositories */ + `
		    - Permission: {}` //TODO + master.Permission
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetWerckerCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Wercker {
		str = `
		wercker:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Wercker.Enabled) +
			GetWerckerCiMasters(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Wercker)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetWerckerCiMasters(wercker *ci.Wercker) string {
	str := ""

	if nil != wercker.Masters {
		str += `
		  masters:`
		for _, master := range wercker.Masters {
			str += `
		    - name: ` + master.Name + `
		    - Address: ` + master.Address + `
		    - user: ` + master.User + `
		    - token: ` + master.Token + `
		    - Permission: {}` //TODO + master.Permission
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetConcourseCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Concourse {
		str = `
		concourse:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Concourse.Enabled) +
			GetConcourseCiMasters(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Concourse)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetConcourseCiMasters(concourse *ci.Concourse) string {
	str := ""

	if nil != concourse.Masters {
		str += `
		  masters:`
		for _, master := range concourse.Masters {
			str += `
		    - name: ` + master.Name + `
		    - url: ` + master.Url + `
		    - username: ` + master.Username + `
		    - password: ` + master.Password + `
		    - Permission: {}` //TODO + master.Permission
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetGcbCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Gcb {
		str = `
		gcb:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Gcb.Enabled) +
			GetGcbCiAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Gcb)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetGcbCiAccounts(gcb *ci.Gcb) string {
	str := ""

	if nil != gcb.Accounts {
		str += `
		  accounts:`
		for _, account := range gcb.Accounts {
			str += `
		    - name: ` + account.Name + `
		      permission: ` + account.Permission + `
		      project: ` + account.Project + `
		      subscriptionName: ` + account.SubscriptionName + `
		      jsonKey: ` + account.JsonKey
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetCodebuildCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Codebuild {
		str = `
		codebuild:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Codebuild.Enabled) +
			GetCodebuildCiAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci.Codebuild)

		str = strings.Replace(str, "\t", "    ", -1)
	}
	return str
}

func GetCodebuildCiAccounts(codebuild *ci.Codebuild) string {
	str := ""

	if nil != codebuild.Accounts {
		str += `
		  accounts:`
		for _, account := range codebuild.Accounts {
			str += `
		    - name: ` + account.Name + `
		      permission: ` + account.Permission + `
		      accountId: ` + account.AccountId + `
		      assumeRole: ` + account.AssumeRole + `
		      region: ` + account.Region
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}
