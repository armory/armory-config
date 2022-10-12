// Package fileio supports reading and writing config files to the filesystem.
package fileio

import (
	"fmt"
	"io/fs"
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

// Check if the provided string is a directory
func IsDirectory(dir string) (bool, error) {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false, err
	}

	if fileInfo.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
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

func GetListOfFiles(dirToSearch string) ([]fs.FileInfo, error) {
	dir, err := os.Open(dirToSearch)
	if err != nil {
		return nil, err
	}

	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func FindFileInDir(dirToSearch string, fileToFind string) (string, error) {
	files, err := GetListOfFiles(dirToSearch)
	if err != nil {
		return "", err
	}

	fmt.Println("Reading " + dirToSearch)

	for _, file := range files {
		if file.Name() == fileToFind {
			fmt.Println("Found:" + fileToFind)
			fmt.Println("Full Path:" + dirToSearch + "/" + fileToFind)
			return dirToSearch + "/" + file.Name(), nil
		}
	}

	return "", fmt.Errorf("File:[%s] not found inside:[%s]", fileToFind, dirToSearch)
}
