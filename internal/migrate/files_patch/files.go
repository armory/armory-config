package files_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
)

func GetFiles(KustomizeData structs.Kustomize) string {

	return GetFilesData(KustomizeData)
}

func GetFilesData(KustomizeData structs.Kustomize) string {
	str := `
	#  profiles__rosco__packer__example-packer-config.json: |
	#    {
	#      "packerSetting": "someValue"
	#    }
	#  profiles__rosco__packer__my_custom_script.sh: |
	#    #!/bin/bash -e
	#    echo "hello world!"`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
