package structs

import "github.com/austinthao5/golang_proto_test/config/deploymentConfigurations"

type Kustomize struct {
	Skip_validations           string
	Spin_flavor                string
	Output_dir                 string
	Halyard                    *deploymentConfigurations.HalFile
	ProfilesConfigFiles        map[string]string
	ServiceSettingsConfigFiles map[string]string
	CurrentDeploymentPos       int
	Header                     string
}

// Struct for the fields inside of deploymentConfigurations
type FieldType int64

const (
	F_array  FieldType = 0
	F_object FieldType = 1
	F_map    FieldType = 2
)

type Fields struct {
	Field_name   string
	Field_type   FieldType
	Field_parent string
}

func GetFieldTypeToString(value FieldType) string {
	str := ""

	if F_array == value {
		str = "[]"
	} else if F_object == value {
		str = "{}"
	} else if F_map == value {
		str = "{}"
	}

	return str
}
