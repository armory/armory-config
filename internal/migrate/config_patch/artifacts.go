package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/artifacts"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
		bitbucket:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Bitbucket.Enabled, 5, true) +
			GetBitbucketAccounts(artifactsReference.Bitbucket)
	}

	return str
}

func GetBitbucketAccounts(bitbucket *artifacts.Bitbucket) string {
	str := ""

	if nil != bitbucket.Accounts {
		str += `
		  accounts:`
		for _, account := range bitbucket.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`usernamePasswordFile: `, account.UsernamePasswordFile, 7, true)
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
		gcs:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Gcs.Enabled, 5, true) +
			GetGcsAccounts(artifactsReference.Gcs)
	}

	return str
}

func GetGcsAccounts(gcs *artifacts.Gcs) string {
	str := ""

	if nil != gcs.Accounts {
		str += `
		  accounts:`
		for _, account := range gcs.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`jsonPath: `, account.JsonPath, 7, true)
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
		oracle:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Oracle.Enabled, 5, true) +
			GetOracleAccounts(artifactsReference.Oracle)
	}

	return str
}

func GetOracleAccounts(oracle *artifacts.Oracle) string {
	str := ""

	if nil != oracle.Accounts {
		str += `
		  accounts:`
		for _, account := range oracle.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`namespace: `, account.Namespace, 7, true) +
				helpers.PrintFmtStr(`region: `, account.Region, 7, true) +
				helpers.PrintFmtStr(`userId: `, account.UserId, 7, true) +
				helpers.PrintFmtStr(`fingerprint: `, account.Fingerprint, 7, true) +
				helpers.PrintFmtStr(`sshPrivateKeyFilePath: `, account.SshPrivateKeyFilePath, 7, true) +
				helpers.PrintFmtStr(`privateKeyPassphrase: `, account.PrivateKeyPassphrase, 7, true) +
				helpers.PrintFmtStr(`tenancyId: `, account.TenancyId, 7, true)
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
		github:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Github.Enabled, 5, true) +
			GetGithubAccounts(artifactsReference.Github)
	}

	return str
}

func GetGithubAccounts(github *artifacts.Github) string {
	str := ""

	if nil != github.Accounts {
		str += `
		  accounts:`
		for _, account := range github.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`userId: `, account.UserId, 7, true) +
				helpers.PrintFmtStr(`token: `, account.Token, 7, true) +
				helpers.PrintFmtStr(`tokenFile: `, account.TokenFile, 7, true)
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
		gitlab:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Gitlab.Enabled, 5, true) +
			GetGitlabAccounts(artifactsReference.Gitlab)
	}

	return str
}

func GetGitlabAccounts(gitlab *artifacts.Gitlab) string {
	str := ""

	if nil != gitlab.Accounts {
		str += `
		  accounts:`
		for _, account := range gitlab.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`token: `, account.Token, 7, true) +
				helpers.PrintFmtStr(`tokenFile: `, account.TokenFile, 7, true)
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
		gitrepo:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Gitrepo.Enabled, 5, true) +
			GetGitrepoAccounts(artifactsReference.Gitrepo)
	}

	return str
}

func GetGitrepoAccounts(gitrepo *artifacts.Gitrepo) string {
	str := ""

	if nil != gitrepo.Accounts {
		str += `
		  accounts:`
		for _, account := range gitrepo.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`usernamePasswordFile: `, account.UsernamePasswordFile, 7, true) +
				helpers.PrintFmtStr(`token: `, account.Token, 7, true) +
				helpers.PrintFmtStr(`sshPrivateKeyFilePath: `, account.SshPrivateKeyFilePath, 7, true) +
				helpers.PrintFmtStr(`sshPrivateKeyPassphrase: `, account.SshPrivateKeyPassphrase, 7, true) +
				helpers.PrintFmtStr(`sshKnownHostsFilePath: `, account.SshKnownHostsFilePath, 7, true) +
				helpers.PrintFmtBool(`sshTrustUnknownHosts: `, account.SshTrustUnknownHosts, 7, true)
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
		http:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Http.Enabled, 5, true) +
			GetHttpAccounts(artifactsReference.Http)
	}

	return str
}

func GetHttpAccounts(http *artifacts.Http) string {
	str := ""

	if nil != http.Accounts {
		str += `
		  accounts:`
		for _, account := range http.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`usernamePasswordFile: `, account.UsernamePasswordFile, 7, true)
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
		helm:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Helm.Enabled, 5, true) +
			GetHelmAccounts(artifactsReference.Helm)
	}

	return str
}

func GetHelmAccounts(helm *artifacts.Helm) string {
	str := ""

	if nil != helm.Accounts {
		str += `
		  accounts:`
		for _, account := range helm.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`repository: `, account.Repository, 7, true) +
				helpers.PrintFmtStr(`username: `, account.Username, 7, true) +
				helpers.PrintFmtStr(`password: `, account.Password, 7, true) +
				helpers.PrintFmtStr(`usernamePasswordFile: `, account.UsernamePasswordFile, 7, true)
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
		s3:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.S3.Enabled, 5, true) +
			GetS3Accounts(artifactsReference.S3)
	}

	return str
}

func GetS3Accounts(s3 *artifacts.S3) string {
	str := ""

	if nil != s3.Accounts {
		str += `
		  accounts:`
		for _, account := range s3.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`apiEndpoint: `, account.ApiEndpoint, 7, true) +
				helpers.PrintFmtStr(`apiRegion: `, account.ApiRegion, 7, true) +
				helpers.PrintFmtStr(`region: `, account.Region, 7, true) +
				helpers.PrintFmtStr(`awsAccessKeyId: `, account.AwsAccessKeyId, 7, true) +
				helpers.PrintFmtStr(`awsSecretAccessKey: `, account.AwsSecretAccessKey, 7, true)
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
		maven:` +
			helpers.PrintFmtBool(`enabled: `, artifactsReference.Maven.Enabled, 5, true) +
			GetMavenAccounts(artifactsReference.Maven)
	}

	return str
}

func GetMavenAccounts(maven *artifacts.Maven) string {
	str := ""

	if nil != maven.Accounts {
		str += `
		  accounts:`
		for _, account := range maven.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`repositoryUrl: `, account.RepositoryUrl, 7, true)
		}
	} else {
		str += `
		  accounts: []`
	}

	return str
}
