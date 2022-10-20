package migrate

import "strings"

func GetFiles(KustomizeData Kustomize) string {

	return GetFilesData(KustomizeData)
}

func GetFilesData(KustomizeData Kustomize) string {
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
