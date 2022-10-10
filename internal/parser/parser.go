// Package fileio supports reading and writing config files to the filesystem.
package parser

import (
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"

	// "github.com/austinthao5/api/client/config"
	"github.com/austinthao5/golang_proto_test/config/deploymentconfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/validate"
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
// func ParseHalConfig(halPath string) (*config.Hal, error) {
func ParseHalConfig(halPath string) (*providers.AppEngine, error) {
	data, err := ioutil.ReadFile(halPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q. Error: %v", halPath, err)
	}

	// Debug
	// fmt.Println("===RAW hal data START===\n" + string(data) + "\n===RAW hal data END===")

	// hal := &config.Hal{}
	hal := &providers.AppEngine{}
	if err := Unmarshal(data, hal); err != nil {
		return nil, fmt.Errorf("unable to unmarshal %q: %v", halPath, err)
	}

	if err := validate.HalConfig(hal); err != nil {
		return nil, fmt.Errorf("unable to validate %q: %v", halPath, err)
	}

	return hal, nil
}

// WriteConfigs generates the service configs from the supplied hal, and writes
// them to the directory dir.
// func WriteConfigs(hal *config.Hal, dir string) error {
/*func WriteConfigs(hal *providers.AppEngine, dir string) error {
	if err := ensureDirectory(dir); err != nil {
		return err
	}

	services := transform.HalToServiceConfigs(hal)
	configFiles, err := transform.GenerateConfigFiles(services)
	if err != nil {
		return err
	}

	for _, file := range configFiles.GetConfigFile() {
		if err := ioutil.WriteFile(filepath.Join(dir, file.GetName()), file.GetContents(), 0666); err != nil {
			return err
		}
	}

	return nil
}*/

func ensureDirectory(d string) error {
	stat, err := os.Stat(d)
	if os.IsNotExist(err) {
		return os.MkdirAll(d, 0755)
	} else if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", d)
	}
	return nil
}
