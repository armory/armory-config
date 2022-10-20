package config_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetPersistentStorage(KustomizeData structs.Kustomize) string {

	return GetPersistentStoreType(KustomizeData) +
		GetAzsStorage(KustomizeData) +
		GetGcsStorage(KustomizeData) +
		GetRedisStorage(KustomizeData) +
		GetS3Storage(KustomizeData) +
		GetOracleStorage(KustomizeData)
}

func GetPersistentStoreType(KustomizeData structs.Kustomize) string {
	str := `
	persistentStoreType: s3`
	str = strings.Replace(str, "\t", "        ", -1)
	return str
	// return `persistentStoreType: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].PersistentStorage.persistentStoreType
}

func GetAzsStorage(KustomizeData structs.Kustomize) string {
	str := ""

	/*if nil != KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].PersistentStorage.Azs {
			str = `    azs: `
		} else {
			str = `    azs: {}
	`
		}*/

	str = `
	azs: {}`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetGcsStorage(KustomizeData structs.Kustomize) string {
	str := ""

	/*if nil != KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].PersistentStorage.Gcs {
			str = `    gcs: `
		} else {
			str = `    gcs: {}
	`
		}*/

	str = `
	gcs:
	  rootFolder: front50
	  bucketLocation: ''`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetRedisStorage(KustomizeData structs.Kustomize) string {
	str := ""

	/*if nil != KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].PersistentStorage.Redis {
			str = `    redis: `
		} else {
			str = `    redis: {}
	`
		}*/

	str = `
	redis: {}`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetS3Storage(KustomizeData structs.Kustomize) string {
	str := ""

	/*if nil != KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].PersistentStorage.S3 {
			str = `    s3: `
		} else {
			str = `    s3: {}
	`
		}*/

	str = `
	s3:
	  bucket: spinnaker
	  rootFolder: spinnaker
	  region: us-west-2
	  pathStyleAccess: false
	  accessKeyId: XXXXX
	  secretAccessKey: XXXXX`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}

func GetOracleStorage(KustomizeData structs.Kustomize) string {
	str := ""

	/*if nil != KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].PersistentStorage.Oracle {
			str = `    oracle: `
		} else {
			str = `    oracle: {}
	`
		}*/

	str = `
	oracle: {}`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
