package files_patch

import (
	"log"
	"regexp"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetFiles(KustomizeData structs.Kustomize) string {
	str := ""

	if nil != KustomizeData.ProfilesConfigFiles {
		str = GetFilesData(KustomizeData)
	}
	return str
}

func GetFilesData(KustomizeData structs.Kustomize) string {
	str := ""

	for key, value := range KustomizeData.ProfilesConfigFiles {

		if !isFileBlacklisted(key) {
			key = transformKey(key)

			str += `
			` + key + `: |
			` + filesFormatContent(value)
		}
	}

	str = filesFormatFix(str)
	str = strings.TrimSuffix(str, `          `)
	return str
}

// This function adds profile to the start of the string and changes / to _
func transformKey(str string) string {
	str = strings.Replace(str, "/", "_", -1)

	if str[0] != '_' {
		str = "_" + str
	}

	str = "profile" + str
	return str
}

var blackListedFiles = [12]string{
	"clouddriver-local.yml",
	"settings-local.js",
	"echo-local.yml",
	"fiat-local.yml",
	"front50-local.yml",
	"gate-local.yml",
	"igor-local.yml",
	"kayenta-local.yml",
	"orca-local.yml",
	"rosco-local.yml",
	"dinghy-local.yml",
	"spinnaker-local.yml",
}

func isFileBlacklisted(fileName string) bool {
	isBlacklisted := false

	for _, blackListedFile := range blackListedFiles {
		if strings.Contains(fileName, blackListedFile) {
			isBlacklisted = true
		}
	}

	return isBlacklisted
}

// Function that edit the read file to have the correct spacing
func filesFormatContent(str string) string {
	reg := regexp.MustCompile(`(\n)`)
	str = reg.ReplaceAllString(str, `$1          `)

	return str
}

// Function that edit the final file to have the correct spacing
func filesFormatFix(str string) string {
	str = strings.Replace(str, "\t", "  ", -1)

	reg := regexp.MustCompile(`(: \|\n)`)
	str = reg.ReplaceAllString(str, `$1  `)

	return str
}

func WriteConfigFiles(KustomizeData structs.Kustomize) string {

	str := ``
	account := KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Kubernetes.Accounts
	filename := ``
	s := ``

	var fileList []string
	for i := range account {
		if !strings.Contains(account[i].KubeconfigFile, `encrypted:`) && !strings.Contains(account[i].KubeconfigFile, `encryptedFile:`) {

			fileAlreadyAdded := false
			test_file, err := fileio.ReadFile(account[i].KubeconfigFile)

			if err != nil {
				log.Fatal(err)
			}
			s = string(test_file)
			filename = strings.Replace(account[i].KubeconfigFile, `/`, `__`, -1)
			for k := range fileList {
				if filename == fileList[k] {
					fileAlreadyAdded = true
				}

			}
			if !fileAlreadyAdded {
				str += `
			` + filename + `: |
			` + filesFormatContent(s)
				fileList = append(fileList, filename)
				// fmt.Println(filename)
			}
		}
	}

	pubsub := KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google.Subscriptions
	for i := range pubsub {
		if !strings.Contains(pubsub[i].TemplatePath, `encrypted:`) && !strings.Contains(pubsub[i].TemplatePath, `encryptedFile:`) {

			fileAlreadyAdded := false
			test_file, err := fileio.ReadFile(pubsub[i].TemplatePath)

			if err != nil {
				log.Fatal(err)
			}
			s = string(test_file)
			filename = strings.Replace(pubsub[i].TemplatePath, `/`, `__`, -1)
			for k := range fileList {
				if filename == fileList[k] {
					fileAlreadyAdded = true
				}

			}
			if !fileAlreadyAdded {
				str += `
			` + filename + `: |
			` + filesFormatContent(s)
				fileList = append(fileList, filename)
				// fmt.Println(filename)
			}
		}
	}

	credentialPath := KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Google.CredentialPath

	if credentialPath != `` && !strings.Contains(credentialPath, `encrypted:`) && !strings.Contains(credentialPath, `encryptedFile:`) {
		test_file, err := fileio.ReadFile(credentialPath)

		if err != nil {
			log.Fatal(err)
		}
		s = string(test_file)

		filename = strings.Replace(credentialPath, `/`, `__`, -1)
		str += `
			` + filename + `: |
			` + filesFormatContent(s)
	}

	str = filesFormatFix(str)

	return str
}
