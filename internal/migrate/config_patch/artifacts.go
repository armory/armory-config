package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/artifacts"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts {
		str = GetBitbucketArtifacts(KustomizeData) +
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
	return str
}

func GetBitbucketArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Bitbucket {
		str = `
		bitbucket:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Bitbucket.Enabled) +
			GetBitbucketAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Bitbucket)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetBitbucketAccounts(bitbucket *artifacts.Bitbucket) string {
	str := ""

	if nil != bitbucket.Accounts {
		str += `
		  accounts:`
		for _, account := range bitbucket.Accounts {
			str += `
		    - name: ` + account.Name + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      usernamePasswordFile: ` + account.UsernamePasswordFile
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetGcsArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gcs {
		str = `
		gcs:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gcs.Enabled) +
			GetGcsAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gcs)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetGcsAccounts(gcs *artifacts.Gcs) string {
	str := ""

	if nil != gcs.Accounts {
		str += `
		  accounts:`
		for _, account := range gcs.Accounts {
			str += `
		    - name: ` + account.Name + `
		      jsonPath: ` + account.JsonPath
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetOracleArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Oracle {
		str = `
		oracle:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Oracle.Enabled) +
			GetOracleAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Oracle)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetOracleAccounts(oracle *artifacts.Oracle) string {
	str := ""

	if nil != oracle.Accounts {
		str += `
		  accounts:`
		for _, account := range oracle.Accounts {
			str += `
		    - name: ` + account.Name + `
		      namespace: ` + account.Namespace + `
		      region: ` + account.Region + `
		      userId: ` + account.UserId + `
		      fingerprint: ` + account.Fingerprint + `
		      sshPrivateKeyFilePath: ` + account.SshPrivateKeyFilePath + `
		      privateKeyPassphrase: ` + account.PrivateKeyPassphrase + `
		      tenancyId: ` + account.TenancyId
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetGithubArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Github {
		str = `
		github:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Github.Enabled) +
			GetGithubAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Github)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetGithubAccounts(github *artifacts.Github) string {
	str := ""

	if nil != github.Accounts {
		str += `
		  accounts:`
		for _, account := range github.Accounts {
			str += `
		    - name: ` + account.Name + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      userId: ` + account.UserId + `
		      token: ` + account.Token + `
		      tokenFile: ` + account.TokenFile
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetGitlabArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gitlab {
		str = `
		gitlab:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gitlab.Enabled) +
			GetGitlabAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gitlab)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetGitlabAccounts(gitlab *artifacts.Gitlab) string {
	str := ""

	if nil != gitlab.Accounts {
		str += `
		  accounts:`
		for _, account := range gitlab.Accounts {
			str += `
		    - name: ` + account.Name + `
		      token: ` + account.Token + `
		      tokenFile: ` + account.TokenFile
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetGitrepoArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gitrepo {
		str = `
		gitrepo:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gitrepo.Enabled) +
			GetGitrepoAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Gitrepo)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetGitrepoAccounts(gitrepo *artifacts.Gitrepo) string {
	str := ""

	if nil != gitrepo.Accounts {
		str += `
		  accounts:`
		for _, account := range gitrepo.Accounts {
			str += `
		    - name: ` + account.Name + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      usernamePasswordFile: ` + account.UsernamePasswordFile + `
		      token: ` + account.Token + `
		      tokenFile: ` + account.TokenFile + `
		      sshPrivateKeyFilePath: ` + account.SshPrivateKeyFilePath + `
		      sshPrivateKeyPassphrase: ` + account.SshPrivateKeyPassphrase + `
		      sshKnownHostsFilePath: ` + account.SshKnownHostsFilePath + `
		      sshTrustUnknownHosts: ` + strconv.FormatBool(account.SshTrustUnknownHosts)
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetHttpArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Http {
		str = `
		http:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Http.Enabled) +
			GetHttpAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Http)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetHttpAccounts(http *artifacts.Http) string {
	str := ""

	if nil != http.Accounts {
		str += `
		  accounts:`
		for _, account := range http.Accounts {
			str += `
		    - name: ` + account.Name + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      usernamePasswordFile: ` + account.UsernamePasswordFile
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetHelmArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Helm {
		str = `
		helm:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Helm.Enabled) +
			GetHelmAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Helm)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetHelmAccounts(helm *artifacts.Helm) string {
	str := ""

	if nil != helm.Accounts {
		str += `
		  accounts:`
		for _, account := range helm.Accounts {
			str += `
		    - name: ` + account.Name + `
		      repository: ` + account.Repository + `
		      username: ` + account.Username + `
		      password: ` + account.Password + `
		      usernamePasswordFile: ` + account.UsernamePasswordFile
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetS3Artifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.S3 {
		str = `
		s3:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.S3.Enabled) +
			GetS3Accounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.S3)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetS3Accounts(s3 *artifacts.S3) string {
	str := ""

	if nil != s3.Accounts {
		str += `
		  accounts:`
		for _, account := range s3.Accounts {
			str += `
		    - name: ` + account.Name + `
		      apiEndpoint: ` + account.ApiEndpoint + `
		      apiRegion: ` + account.ApiRegion + `
		      region: ` + account.Region + `
		      awsAccessKeyId: ` + account.AwsAccessKeyId + `
		      awsSecretAccessKey: ` + account.AwsSecretAccessKey
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}

func GetMavenArtifacts(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Maven {
		str = `
		maven:
		  enabled: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Maven.Enabled) +
			GetMavenAccounts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts.Maven)
		str = strings.Replace(str, "\t", "    ", -1)
	}

	return str
}

func GetMavenAccounts(maven *artifacts.Maven) string {
	str := ""

	if nil != maven.Accounts {
		str += `
		  accounts:`
		for _, account := range maven.Accounts {
			str += `
		    - name: ` + account.Name + `
		      repositoryUrl: ` + account.RepositoryUrl
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}
