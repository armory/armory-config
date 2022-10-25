package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetPersistentStorage(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage {
		str = GetPersistentStoreType(KustomizeData) +
			GetAzsStorage(KustomizeData) +
			GetGcsStorage(KustomizeData) +
			GetRedisStorage(KustomizeData) +
			GetS3Storage(KustomizeData) +
			GetOracleStorage(KustomizeData)
	}
	return str
}

func GetPersistentStoreType(KustomizeData structs.Kustomize) string {
	str := `
	persistentStoreType: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.PersistentStoreType
	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetAzsStorage(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Azs &&
		"" != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Azs.StorageAccountName {
		str = `  azs:
			StorageAccountName: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Azs.StorageAccountName + `
			StorageAccountKey: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Azs.StorageAccountKey + `
			StorageContainerName: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Azs.StorageContainerName
	} else {
		str = `  azs: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}

func GetGcsStorage(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs &&
		"" != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.JsonPath {
		str = `  gcs:
			jsonPath: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.JsonPath + `
			project: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.Project + `
			bucket: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.Bucket + `
			rootFolder: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.RootFolder + `
			bucketLocation: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.BucketLocation
	} else {
		str = `  gcs: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}

func GetRedisStorage(KustomizeData structs.Kustomize) string {
	str := ""

	/*if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.R &&
		"" != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs.JsonPath {
			str = `  redis: `
		} else {
			str = `  redis: {}
	`
		}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str*/

	str = `
	redis: {}`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetS3Storage(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3 &&
		"" != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.Bucket {
		str = `  s3:
			bucket: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.Bucket + `
			rootFolder: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.RootFolder + `
			region: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.Region + `
			pathStyleAccess: ` + strconv.FormatBool(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.PathStyleAccess) + `
			accessKeyId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.AccessKeyId + `
			secretAccessKey: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3.SecretAccessKey
	} else {
		str = `    s3: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}

func GetOracleStorage(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle &&
		"" != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.BucketName {
		str = `  oracle:
			bucketName: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.BucketName + `
			namespace: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.Namespace + `
			compartmentId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.CompartmentId + `
			region: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.Region + `
			userId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.UserId + `
			fingerprint: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.Fingerprint + `
			sshPrivateKeyFilePath: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.SshPrivateKeyFilePath + `
			privateKeyPassphrase: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.PrivateKeyPassphrase + `
			tenancyId: ` + KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle.TenancyId
	} else {
		str = `  oracle: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}
