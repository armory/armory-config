package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetArtifacts(KustomizeData structs.Kustomize) string {

	return GetBitbucketArtifacts(KustomizeData) +
		GetGcsArtifacts(KustomizeData) +
		GetOracleArtifacts(KustomizeData) +
		GetGithubArtifacts(KustomizeData) +
		GetGitlabArtifacts(KustomizeData) +
		GetGitrepoArtifacts(KustomizeData) +
		GetHttpArtifacts(KustomizeData) +
		GetHelmArtifacts(KustomizeData) +
		GetS3Artifacts(KustomizeData) +
		GetMavenArtifacts(KustomizeData)
}

func GetBitbucketArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	bitbucket:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Bitbucket.Enable +
	GetBitbucketAccounts(KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Bitbucket)
	*/

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

/*func GetBitbucketAccounts(bitbucket artifacts.Bitbucket) string {
	str := ""

	if nil != bitbucket.BitbucketAcc {
		str += `
	  accounts:
`
		for _, accounts := range bitbucket.BitbucketAcc {
			str += `
	    - name:` + accounts.Name + `
	    - Username:` + accounts.Username + `
	    - Password:` + accounts.Password + `
	    - Token:` + accounts.Token + `
	    - TokenFile:` + accounts.TokenFile + `
	    - UsernamePasswordFile:` + accounts.Password
		}
	} else {
		str += `	  accounts: []`
	}

	str += `
		accounts: []`
	return str
}*/

func GetGcsArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	gcs:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Gcs.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetOracleArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	oracle:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Oracle.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGithubArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	github:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Github.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGitlabArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	gitlab:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Gitlab.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGitrepoArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	gitrepo:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Gitrepo.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetHttpArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	http:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Http.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetHelmArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	helm:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Helm.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetS3Artifacts(KustomizeData structs.Kustomize) string {
	str := `
	s3:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.S3.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetMavenArtifacts(KustomizeData structs.Kustomize) string {
	str := `
	maven:
	  enabled: ` + "false" /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Artifacts.Maven.Enable */ + `
	  accounts: []`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
