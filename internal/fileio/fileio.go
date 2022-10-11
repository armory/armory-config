// Package fileio supports reading and writing config files to the filesystem.
package fileio

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/austinthao5/golang_proto_test/config/deploymentconfigurations/providers"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q. Error: %v", filename, err)
	}
	return data, err
}

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, err
}

func WriteConfigsTmp(filename string, data string) error {
	file, err := CreateFile(filename)
	if err != nil {
		return err
	}

	buffer := []byte(data)

	_, err = file.Write(buffer)

	return err
}

// WriteConfigs generates the service configs from the supplied hal, and writes
// them to the directory dir.
// func WriteConfigs(hal *config.Hal, dir string) error {
func WriteConfigs(hal *providers.AppEngine, dir string) error {
	if err := EnsureDirectory(dir); err != nil {
		return err
	}

	// services := transform.HalToServiceConfigs(hal)
	// configFiles, err := transform.GenerateConfigFiles(services)
	// if err != nil {
	// 	return err
	// }

	// for _, file := range configFiles.GetConfigFile() {
	// 	if err := ioutil.WriteFile(filepath.Join(dir, file.GetName()), file.GetContents(), 0666); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func EnsureDirectory(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755) //os.ModePerm
	} else if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", dir)
	}
	return nil
}
