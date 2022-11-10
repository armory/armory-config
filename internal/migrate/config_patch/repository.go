package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/repository"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetRepository(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Repository {
		str = GetArtifactoryRepository(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Repository.Artifactory) +
			GetNexusRepository(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Repository.Nexus)
	}
	return str
}

func GetArtifactoryRepository(artifactoryRef *repository.Artifactory) string {
	str := ""

	if nil != artifactoryRef {
		str = `
		artifactory:` +
			helpers.PrintFmtBool(`enabled: `, artifactoryRef.Enabled, 5, true) +
			GetArtifactorySearches(artifactoryRef.Searches)
		str = strings.Replace(str, "\t", "    ", -1)
	} else {
		str = `
		artifactory:
		  enabled: false
		  searches: []`
	}

	return str
}

func GetArtifactorySearches(artifactorySearchesRef []*repository.ArtifactorySearch) string {
	str := ""

	if nil != artifactorySearchesRef {
		str += `
		  searches:`
		for _, searches := range artifactorySearchesRef {
			str += helpers.PrintFmtStr(`- name: `, searches.Name, 6, true) +
				getArtifactoryPermissions(searches.Permissions) +
				helpers.PrintFmtStr(`baseUrl: `, searches.BaseUrl, 7, true) +
				helpers.PrintFmtStr(`repo: `, searches.Repo, 7, true) +
				helpers.PrintFmtStr(`groupId: `, searches.GroupId, 7, true) +
				helpers.PrintFmtStr(`repoType: `, searches.RepoType, 7, true) +
				helpers.PrintFmtStr(`username: `, searches.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, searches.Password, 7, true)
		}
	} else {
		str += `
		  searches: []`
	}

	return str
}

func GetNexusRepository(nexusRef *repository.Nexus) string {
	str := ""

	if nil != nexusRef {
		str = `
		nexus:` +
			helpers.PrintFmtBool(`enabled: `, nexusRef.Enabled, 5, true) +
			GetNexusSearches(nexusRef.Searches)
		str = strings.Replace(str, "\t", "    ", -1)
	} else {
		str = `
		nexus:
		  enabled: false
		  searches: []`
	}

	return str
}

func GetNexusSearches(nexusSearchesRef []*repository.NexusSearch) string {
	str := ""

	if nil != nexusSearchesRef {
		str += `
		  searches:`
		for _, searches := range nexusSearchesRef {
			str += helpers.PrintFmtStr(`- name: `, searches.Name, 6, true) +
				getArtifactoryPermissions(searches.Permissions) +
				helpers.PrintFmtStr(`baseUrl: `, searches.BaseUrl, 7, true) +
				helpers.PrintFmtStr(`repo: `, searches.Repo, 7, true) +
				helpers.PrintFmtStr(`nodeId: `, searches.NodeId, 7, true) +
				helpers.PrintFmtStr(`username: `, searches.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, searches.Password, 7, true)
		}
	} else {
		str += `
		  searches: []`
	}

	return str
}

func getArtifactoryPermissions(permissionsRef *permissions.Permissions) string {
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
