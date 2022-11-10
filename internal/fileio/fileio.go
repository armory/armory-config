// Package fileio supports reading and writing config files to the filesystem.
package fileio

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
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

// This function returns a map that uses the full file path as a key and as value the files data
func GetFilesAndData(profilesDir string, currentSubDir string) (map[string]string, error) {
	profiles_config_files := make(map[string]string)

	profiles_files, err := GetListOfFiles(profilesDir)
	if err != nil {
		return nil, err
	}

	for _, profiles_file := range profiles_files {

		isDir, err := IsDirectory(profilesDir + "/" + profiles_file.Name())
		if isDir {
			config_maps, err := GetFilesAndData(profilesDir+"/"+profiles_file.Name(), currentSubDir+"/"+profiles_file.Name())
			if err != nil {
				return nil, err
			}

			for fileName, config_map := range config_maps {
				profiles_config_files[fileName] = config_map
				// fmt.Println("fileName:" + fileName + " config_file: " + config_file)
			}
			continue
		}

		local_file, err := ReadFile(profilesDir + "/" + profiles_file.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var buf = string(local_file)

		if "" != currentSubDir {
			profiles_config_files[currentSubDir+"/"+profiles_file.Name()] = buf
		} else {
			profiles_config_files[profiles_file.Name()] = buf
		}

		// fmt.Println("Found " + profiles_file.Name())
	}

	return profiles_config_files, nil
}
