package config_patch

import (
	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/persistentStorage"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
	return helpers.PrintFmtStr(`persistentStoreType: `, storageReference.PersistentStoreType, 4, true)
}

func GetAzsStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.Azs &&
		"" != storageReference.Azs.StorageAccountName {
		str = `
		azs:` +
			helpers.PrintFmtStr(`storageAccountName: `, storageReference.Azs.StorageAccountName, 5, true) +
			helpers.PrintFmtStr(`storageAccountKey: `, storageReference.Azs.StorageAccountKey, 5, true) +
			helpers.PrintFmtStr(`storageContainerName: `, storageReference.Azs.StorageContainerName, 5, true)
	} else {
		str = `
		azs: {}`
	}

	return str
}

func GetGcsStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.Gcs &&
		"" != storageReference.Gcs.Bucket {
		str = `
		gcs:` +
			helpers.PrintFmtStr(`jsonPath: `, storageReference.Gcs.JsonPath, 5, true) +
			helpers.PrintFmtStr(`project: `, storageReference.Gcs.Project, 5, true) +
			helpers.PrintFmtStr(`bucket: `, storageReference.Gcs.Bucket, 5, true) +
			helpers.PrintFmtStr(`rootFolder: `, storageReference.Gcs.RootFolder, 5, true) +
			helpers.PrintFmtStr(`bucketLocation: `, storageReference.Gcs.BucketLocation, 5, true)
	} else {
		str = `
		gcs: {}`
	}

	return str
}

func GetRedisStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	//Todo Missing proto
	/*if nil != storageReference.Redis &&
		"" != storageReference.Gcs.JsonPath {
		str = `
		redis: ` +
			helpers.PrintFmtStr(`AAAAa: `, storageReference.Redis.AAAAa, 5, true) +
	} else {
			str = `
		redis: {}`
	}*/

	str = `
		redis: {}`

	return str
}

func GetS3Storage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.S3 &&
		"" != storageReference.S3.Bucket {
		str = `
		s3:` +
			helpers.PrintFmtStr(`bucket: `, storageReference.S3.Bucket, 5, true) +
			helpers.PrintFmtStr(`rootFolder: `, storageReference.S3.RootFolder, 5, true) +
			helpers.PrintFmtStr(`region: `, storageReference.S3.Region, 5, true) +
			helpers.PrintFmtBool(`pathStyleAccess: `, storageReference.S3.PathStyleAccess, 5, true) +
			helpers.PrintFmtStr(`accessKeyId: `, storageReference.S3.AccessKeyId, 5, true) +
			helpers.PrintFmtStr(`secretAccessKey: `, storageReference.S3.SecretAccessKey, 5, true)
	} else {
		str = `
		s3: {}`
	}

	return str
}

func GetOracleStorage(storageReference *persistentStorage.PersistentStorage) string {
	str := ""

	if nil != storageReference.Oracle &&
		"" != storageReference.Oracle.BucketName {
		str = `
		oracle:` +
			helpers.PrintFmtStr(`bucketName: `, storageReference.Oracle.BucketName, 5, true) +
			helpers.PrintFmtStr(`namespace: `, storageReference.Oracle.Namespace, 5, true) +
			helpers.PrintFmtStr(`compartmentId: `, storageReference.Oracle.CompartmentId, 5, true) +
			helpers.PrintFmtStr(`region: `, storageReference.Oracle.Region, 5, true) +
			helpers.PrintFmtStr(`userId: `, storageReference.Oracle.UserId, 5, true) +
			helpers.PrintFmtStr(`fingerprint: `, storageReference.Oracle.Fingerprint, 5, true) +
			helpers.PrintFmtStr(`sshPrivateKeyFilePath: `, storageReference.Oracle.SshPrivateKeyFilePath, 5, true) +
			helpers.PrintFmtStr(`privateKeyPassphrase: `, storageReference.Oracle.PrivateKeyPassphrase, 5, true) +
			helpers.PrintFmtStr(`tenancyId: `, storageReference.Oracle.TenancyId, 5, true)
	} else {
		str = `
		oracle: {}`
	}

	return str
}
