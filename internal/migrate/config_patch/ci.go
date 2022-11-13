package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/ci"
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
		jenkins:` +
			helpers.PrintFmtBool(`enabled: `, ciReference.Jenkins.Enabled, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, master.Name, 6, true) +
				getCiPermissions(master.Permissions) +
				helpers.PrintFmtStr(`address: `, master.Address, 7, true) +
				helpers.PrintFmtStr(`username: `, master.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, master.Password, 7, true) +
				helpers.PrintFmtBool(`csrf: `, master.Csrf, 7, true)
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
		travis:` +
			helpers.PrintFmtBool(`enabled: `, ciReference.Travis.Enabled, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, master.Name, 6, true) +
				getCiPermissions(master.Permission) +
				helpers.PrintFmtStr(`address: `, master.Address, 7, true) +
				helpers.PrintFmtStr(`baseUrl: `, master.BaseUrl, 7, true) +
				helpers.PrintFmtStr(`githubToken: `, master.GithubToken, 7, true) +
				helpers.PrintFmtInt(`numberOfRepositories: `, master.NumberOfRepositories, 7, true) +
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
		wercker:` +
			helpers.PrintFmtBool(`enabled: `, ciReference.Wercker.Enabled, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, master.Name, 6, true) +
				getCiPermissions(master.Permission) +
				helpers.PrintFmtStr(`address: `, master.Address, 7, true) +
				helpers.PrintFmtStr(`user: `, master.User, 7, true) +
				helpers.PrintFmtStr(`token: `, master.Token, 7, true)
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
		concourse:` +
			helpers.PrintFmtBool(`enabled: `, ciReference.Concourse.Enabled, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, master.Name, 6, true) +
				getCiPermissions(master.Permission) +
				helpers.PrintFmtStr(`url: `, master.Url, 7, true) +
				helpers.PrintFmtStr(`username: `, master.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, master.Password, 7, true)
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
		gcb:` +
			helpers.PrintFmtBool(`enabled: `, ciReference.Gcb.Enabled, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`permission: `, account.Permission, 7, true) +
				helpers.PrintFmtStr(`project: `, account.Project, 7, true) +
				helpers.PrintFmtStr(`subscriptionName: `, account.SubscriptionName, 7, true) +
				helpers.PrintFmtStr(`jsonKey: `, account.JsonKey, 7, true)
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
		codebuild:` +
			helpers.PrintFmtBool(`enabled: `, ciReference.Codebuild.Enabled, 5, true) +
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
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`permission: `, account.Permission, 7, true) +
				helpers.PrintFmtStr(`accountId: `, account.AccountId, 7, true) +
				helpers.PrintFmtStr(`assumeRole: `, account.AssumeRole, 7, true) +
				helpers.PrintFmtStr(`region: `, account.Region, 7, true)
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func getCiPermissions(permissionsRef *permissions.Permissions) string {
	str := ""

	if nil != permissionsRef.READ {
		str += `
		      permissions:
		        READ:`
		for _, permissionRead := range permissionsRef.READ {
			str += `
		          - ` + permissionRead
		}
	} else if nil != permissionsRef.WRITE {
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
