package files_patch

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
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
	str = strings.Replace(str, "/", "__", -1)

	if str[0] != '_' {
		str = "__" + str
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
	str = reg.ReplaceAllString(str, `$1        `)

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

	// Check for KubeconfigFile Strings and append them to the files-patch.yml file.
	str := ``
	account := KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Providers.Kubernetes.Accounts
	filename := ``
	s := ``

	var fileList []string
	for i := range account {
		// Don't look for a file if a secret is passed here. Instead, just ignore it.
		if !strings.Contains(account[i].KubeconfigFile, `encrypted:`) && !strings.Contains(account[i].KubeconfigFile, `encryptedFile:`) {

			fileAlreadyAdded := false
			test_file, err := fileio.ReadFile(account[i].KubeconfigFile)

			if err != nil {
				fmt.Println(err)
			} else {
				s = string(test_file)
				// Check if First Character of the string is a /. If so, we remove it.
				if account[i].KubeconfigFile[0:1] == `/` {
					filename = strings.Replace(account[i].KubeconfigFile, `/`, ``, 1)
				} else {
					filename = account[i].KubeconfigFile
				}
				// Replace / with __ so that the path gets mounted onto Halyard properly
				filename = strings.Replace(filename, `/`, `__`, -1)

				// Check for duplicate files. If duplicate files are referenced, don't add them again.
				for k := range fileList {
					if filename == fileList[k] {
						fileAlreadyAdded = true
					}

				}

				// If the file isn't a duplicate file, add it to the files-patch.yml file.
				if !fileAlreadyAdded {
					str += helpers.PrintFmtStr(filename, `: |`, 3, true) + ` ` + helpers.PrintFmtStr(``, filesFormatContent(s), 4, true)
					fileList = append(fileList, filename)
					// fmt.Println(filename)
				}
			}
		}
	}

	// Check for KubeconfigFile Strings and append them to the files-patch.yml file.
	pubsub := KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Pubsub.Google.Subscriptions
	for i := range pubsub {
		if !strings.Contains(pubsub[i].TemplatePath, `encrypted:`) && !strings.Contains(pubsub[i].TemplatePath, `encryptedFile:`) {

			fileAlreadyAdded := false
			test_file, err := fileio.ReadFile(pubsub[i].TemplatePath)

			if err != nil {
				fmt.Println(err)
			} else {
				s = string(test_file)
				// Check if First Character of the string is a /. If so, we remove it.
				if pubsub[i].TemplatePath[0:1] == `/` {
					filename = strings.Replace(pubsub[i].TemplatePath, `/`, ``, 1)
				} else {
					filename = pubsub[i].TemplatePath
				}
				filename = strings.Replace(filename, `/`, `__`, -1)
				for k := range fileList {
					if filename == fileList[k] {
						fileAlreadyAdded = true
					}

				}
				if !fileAlreadyAdded {
					str += helpers.PrintFmtStr(filename, `: |`, 3, true) + ` ` + helpers.PrintFmtStr(``, filesFormatContent(s), 4, true)
					fileList = append(fileList, filename)
					// fmt.Println(filename)
				}
			}
		}
	}

	credentialPath := KustomizeData.Halyard.DeploymentConfigurations[KustomizeData.CurrentDeploymentPos].Security.Authz.GroupMembership.Google.CredentialPath

	if credentialPath != `` && !strings.Contains(credentialPath, `encrypted:`) && !strings.Contains(credentialPath, `encryptedFile:`) {
		test_file, err := fileio.ReadFile(credentialPath)

		if err != nil {
			fmt.Println(err)
		} else {
			s = string(test_file)
			// Check if First Character of the string is a /. If so, we remove it.
			if credentialPath[0:1] == `/` {
				filename = strings.Replace(credentialPath, `/`, ``, 1)
			} else {
				filename = credentialPath
			}
			filename = strings.Replace(filename, `/`, `__`, -1)
			str += helpers.PrintFmtStr(filename, `: |`, 3, true) + ` ` + helpers.PrintFmtStr(``, filesFormatContent(s), 4, true)
		}

		str = filesFormatFix(str)
	}
	return str
}
