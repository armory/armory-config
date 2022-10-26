package files_patch

import (
	"regexp"
	"strings"

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
	return str
}

// This function adds profile to the start of the string and changes / to _
func transformKey(str string) string {
	str = strings.Replace(str, "/", "_", -1)

	if '_' != str[0] {
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
