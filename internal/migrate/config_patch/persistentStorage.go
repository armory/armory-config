package config_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/persistentStorage"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetPersistentStorage(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage {
		str = GetPersistentStoreType(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage) +
			GetAzsStorage(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage) +
			GetGcsStorage(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage) +
			GetRedisStorage(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage) +
			GetS3Storage(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage) +
			GetOracleStorage(KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].PersistentStorage)
	}
	return str
}

func GetPersistentStoreType(storageReference *persistentStorage.PersistentStorage) string {
	str := `
	persistentStoreType: ` + storageReference.PersistentStoreType
	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetAzsStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.Azs &&
		"" != storageReference.Azs.StorageAccountName {
		str = `  azs:
			StorageAccountName: ` + storageReference.Azs.StorageAccountName + `
			StorageAccountKey: ` + storageReference.Azs.StorageAccountKey + `
			StorageContainerName: ` + storageReference.Azs.StorageContainerName
	} else {
		str = `  azs: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}

func GetGcsStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.Gcs &&
		"" != storageReference.Gcs.JsonPath {
		str = `  gcs:
			jsonPath: ` + storageReference.Gcs.JsonPath + `
			project: ` + storageReference.Gcs.Project + `
			bucket: ` + storageReference.Gcs.Bucket + `
			rootFolder: ` + storageReference.Gcs.RootFolder + `
			bucketLocation: ` + storageReference.Gcs.BucketLocation
	} else {
		str = `  gcs: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}

func GetRedisStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	/*if nil != storageReference.R &&
		"" != storageReference.Gcs.JsonPath {
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

func GetS3Storage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.S3 &&
		"" != storageReference.S3.Bucket {
		str = `  s3:
			bucket: ` + storageReference.S3.Bucket + `
			rootFolder: ` + storageReference.S3.RootFolder + `
			region: ` + storageReference.S3.Region + `
			pathStyleAccess: ` + strconv.FormatBool(storageReference.S3.PathStyleAccess) + `
			accessKeyId: ` + storageReference.S3.AccessKeyId + `
			secretAccessKey: ` + storageReference.S3.SecretAccessKey
	} else {
		str = `    s3: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}

func GetOracleStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.Oracle &&
		"" != storageReference.Oracle.BucketName {
		str = `  oracle:
			bucketName: ` + storageReference.Oracle.BucketName + `
			namespace: ` + storageReference.Oracle.Namespace + `
			compartmentId: ` + storageReference.Oracle.CompartmentId + `
			region: ` + storageReference.Oracle.Region + `
			userId: ` + storageReference.Oracle.UserId + `
			fingerprint: ` + storageReference.Oracle.Fingerprint + `
			sshPrivateKeyFilePath: ` + storageReference.Oracle.SshPrivateKeyFilePath + `
			privateKeyPassphrase: ` + storageReference.Oracle.PrivateKeyPassphrase + `
			tenancyId: ` + storageReference.Oracle.TenancyId
	} else {
		str = `  oracle: {}`
	}

	str = strings.Replace(str, "  ", "\n        ", -1)
	str = strings.Replace(str, "\t\t\t", "          ", -1)
	return str
}
