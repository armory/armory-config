// Package parser marshals and unmarshals protocol buffer messages in YAML format.
package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations"
	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

var nonStrict = protojson.UnmarshalOptions{DiscardUnknown: true}
var strict = protojson.UnmarshalOptions{DiscardUnknown: false}

// Marshal writes the given proto.Message in YAML format.
func Marshal(m proto.Message) ([]byte, error) {
	json, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(json)
}

// Unmarshal reads the given []byte into the given proto.Message, discarding
// any unknown fields in the input.
func Unmarshal(b []byte, m proto.Message) error {
	json, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}
	return nonStrict.Unmarshal(json, m)
}

// UnmarshalStrict reads the given []byte into the given proto.Message. If there
// are any unknown fields in the input, an error is returned.
func UnmarshalStrict(b []byte, m proto.Message) error {
	json, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}
	return strict.Unmarshal(json, m)
}

// ParseHalConfig reads the YAML file at halPath parses it into a *config.Hal.
// The config.Hal is validated; if there are any errors, they will be returned
// in error and *config.Hal will be nil.
func ParseHalConfig(halPath string) (*deploymentConfigurations.HalFile, error) {
	data, err := fileio.ReadFile(halPath)

	if err != nil {
		return nil, err
	}

	//We transform the data to have the correct name or format
	data = []byte(converter(string(data)))

	// Debug
	// fmt.Println("===RAW hal data START===\n" + string(data) + "\n===RAW hal data END===")

	hal := &deploymentConfigurations.HalFile{}
	if err := Unmarshal(data, hal); err != nil {
		return nil, fmt.Errorf("unable to unmarshal %q: %v", halPath, err)
	}

	//Debug JSON
	// fmt.Println("%#V", hal)

	return hal, nil
}

// This function converts specific fields to the correct format for parsing
func converter(str string) string {
	//Convert spin- to spin_ because go doesn't allow variables to be have - and the proto files cannot parse the customSizing data
	reg := regexp.MustCompile(`(\n)      spin-(.*)`)
	str = reg.ReplaceAllString(str, `$1      spin_$2`)

	//Convert github-status to githubStatus, same as above
	reg = regexp.MustCompile(`(\n)    github-status:`)
	str = reg.ReplaceAllString(str, `$1    githubStatus:`)

	//Add any missing fields
	str = updateMissingConfig(str)

	return str
}

// This function checks if there is any missing field and adds it to avoid Unmarshal error
func updateMissingConfig(str string) string {

	fields := []structs.Fields{{
		Field_name:   "huaweicloud",
		Field_type:   structs.F_object,
		Field_parent: "providers",
	}, {
		Field_name:   "dcos",
		Field_type:   structs.F_object,
		Field_parent: "providers",
	}}

	for _, field := range fields {
		if !strings.Contains(str, field.Field_name) {
			reg := regexp.MustCompile("(\n.*)(" + field.Field_parent + ":)\n")
			str = reg.ReplaceAllString(str, `$1$2$1  `+field.Field_name+": "+structs.GetFieldTypeToString(field.Field_type)+"\n")
		}
	}

	return str
}
