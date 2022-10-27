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
		str = GetBitbucketArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetGcsArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetOracleArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetGithubArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetGitlabArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetGitrepoArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetHttpArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetHelmArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetS3Artifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts) +
			GetMavenArtifacts(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Artifacts)
	}
	return str
}

func GetBitbucketArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Bitbucket {
		str = `
		bitbucket:
		  enabled: ` + strconv.FormatBool(artifactsReference.Bitbucket.Enabled) +
			GetBitbucketAccounts(artifactsReference.Bitbucket)
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

func GetGcsArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Gcs {
		str = `
		gcs:
		  enabled: ` + strconv.FormatBool(artifactsReference.Gcs.Enabled) +
			GetGcsAccounts(artifactsReference.Gcs)
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

func GetOracleArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Oracle {
		str = `
		oracle:
		  enabled: ` + strconv.FormatBool(artifactsReference.Oracle.Enabled) +
			GetOracleAccounts(artifactsReference.Oracle)
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

func GetGithubArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Github {
		str = `
		github:
		  enabled: ` + strconv.FormatBool(artifactsReference.Github.Enabled) +
			GetGithubAccounts(artifactsReference.Github)
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

func GetGitlabArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Gitlab {
		str = `
		gitlab:
		  enabled: ` + strconv.FormatBool(artifactsReference.Gitlab.Enabled) +
			GetGitlabAccounts(artifactsReference.Gitlab)
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

func GetGitrepoArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Gitrepo {
		str = `
		gitrepo:
		  enabled: ` + strconv.FormatBool(artifactsReference.Gitrepo.Enabled) +
			GetGitrepoAccounts(artifactsReference.Gitrepo)
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

func GetHttpArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Http {
		str = `
		http:
		  enabled: ` + strconv.FormatBool(artifactsReference.Http.Enabled) +
			GetHttpAccounts(artifactsReference.Http)
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

func GetHelmArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Helm {
		str = `
		helm:
		  enabled: ` + strconv.FormatBool(artifactsReference.Helm.Enabled) +
			GetHelmAccounts(artifactsReference.Helm)
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

func GetS3Artifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.S3 {
		str = `
		s3:
		  enabled: ` + strconv.FormatBool(artifactsReference.S3.Enabled) +
			GetS3Accounts(artifactsReference.S3)
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

func GetMavenArtifacts(artifactsReference *artifacts.Artifacts) string {
	str := ""

	if nil != artifactsReference.Maven {
		str = `
		maven:
		  enabled: ` + strconv.FormatBool(artifactsReference.Maven.Enabled) +
			GetMavenAccounts(artifactsReference.Maven)
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
