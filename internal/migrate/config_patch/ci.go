package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/ci"
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetCi(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci {
		str = GetJenkinsCi(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci) +
			GetTravisCi(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci) +
			GetWerckerCi(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci) +
			GetConcourseCi(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci) +
			GetGcbCi(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci) +
			GetCodebuildCi(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Ci)
	}
	return str
}

func GetJenkinsCi(ciReference *ci.Ci) string {
	str := ""

	if nil != ciReference.Jenkins {
		str = `
		jenkins:
		  enabled: ` + strconv.FormatBool(ciReference.Jenkins.Enabled) +
			GetJenkinsCiMasters(ciReference.Jenkins)
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
		    - name: ` + master.Name +
				getCiPermissions(master.Permission) + `
		      Address: ` + master.Address + `
		      Username: ` + master.Username + `
		      Password: ` + master.Password + `
		      Csrf: ` + strconv.FormatBool(master.Csrf)
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetTravisCi(ciReference *ci.Ci) string {
	str := ""

	if nil != ciReference.Travis {
		str = `
		travis:
		  enabled: ` + strconv.FormatBool(ciReference.Travis.Enabled) +
			GetTravisCiMasters(ciReference.Travis)

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
		    - name: ` + master.Name +
				getCiPermissions(master.Permission) + `
		      Address: ` + master.Address + `
		      BaseUrl: ` + master.BaseUrl + `
		      githubToken: ` + master.GithubToken + `
		      numberOfRepositories: ` + strconv.FormatInt(int64(master.NumberOfRepositories), 10) +
				getFilteredRepositories(master.FilteredRepositories, "filteredRepositories", "- ")
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetWerckerCi(ciReference *ci.Ci) string {
	str := ""

	if nil != ciReference.Wercker {
		str = `
		wercker:
		  enabled: ` + strconv.FormatBool(ciReference.Wercker.Enabled) +
			GetWerckerCiMasters(ciReference.Wercker)

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
		    - name: ` + master.Name +
				getCiPermissions(master.Permission) + `
		      Address: ` + master.Address + `
		      user: ` + master.User + `
		     token: ` + master.Token
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetConcourseCi(ciReference *ci.Ci) string {
	str := ""

	if nil != ciReference.Concourse {
		str = `
		concourse:
		  enabled: ` + strconv.FormatBool(ciReference.Concourse.Enabled) +
			GetConcourseCiMasters(ciReference.Concourse)

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
		    - name: ` + master.Name +
				getCiPermissions(master.Permission) + `
		      url: ` + master.Url + `
		      username: ` + master.Username + `
		      password: ` + master.Password
		}
	} else {
		str += `
		  masters: []`
	}

	return str
}

func GetGcbCi(ciReference *ci.Ci) string {
	str := ""

	if nil != ciReference.Gcb {
		str = `
		gcb:
		  enabled: ` + strconv.FormatBool(ciReference.Gcb.Enabled) +
			GetGcbCiAccounts(ciReference.Gcb)

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

func GetCodebuildCi(ciReference *ci.Ci) string {
	str := ""

	if nil != ciReference.Codebuild {
		str = `
		codebuild:
		  enabled: ` + strconv.FormatBool(ciReference.Codebuild.Enabled) +
			GetCodebuildCiAccounts(ciReference.Codebuild)

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

func getCiPermissions(permissionsRef *permissions.Permissions) string {
	str := ""

	if nil != permissionsRef {
		str += `
		    permissions:
		      READ:`
		for _, permissionRead := range permissionsRef.READ {
			str += `
		        - ` + permissionRead
		}
		str += `
		      WRITE:`
		for _, permissionWrite := range permissionsRef.WRITE {
			str += `
		        - ` + permissionWrite
		}
	} else {
		str += `
		      permissions: {}`
	}

	return str
}

func getFilteredRepositories(stringArray []string, fieldName string, valueToAppend string) string {
	str := ""

	if nil != stringArray {
		str += `
		      ` + fieldName + `:`
		for _, stringValue := range stringArray {
			str += `
		        ` + valueToAppend + stringValue
		}
	} else {
		str += `
		      ` + fieldName + `: []`
	}

	return str
}
