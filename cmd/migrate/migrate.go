package migrate

import (
	"fmt"
	"log"
	"os"

	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/migrate"
	"github.com/austinthao5/golang_proto_test/internal/parser"

	"github.com/spf13/cobra"
)

// Set the Convert Command usage. Expected usage is "armory-config convert --halconfig ~/.hal --output ./test_output"
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a halconfig to Operator config",
	Args:  cobra.ExactArgs(0),
	// Run this function when the command is called.
	Run: func(cmd *cobra.Command, args []string) {
		// res := stringer.Reverse(args[0])
		// fmt.Println(res)

		if len(halconfig_dir) > 0 && len(output_dir) > 0 {

			migrator(halconfig_dir, output_dir, deployment_dir, spin_flavor)

		}
	},
}

// Run this function when the CLI starts up. This initializes the flags for the Convert command.
func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVar(&halconfig_dir, "halconfig", "~/.hal", "Provide the Halconfig directory to convert")
	convertCmd.Flags().StringVar(&output_dir, "output", "./operatorConfig", "Select an output directory")
	convertCmd.Flags().StringVar(&deployment_dir, "deployment", "default", "Select the deployment name being used by Halyard. This is the subdirectory in the ./hal folder where the profiles and service-settings live")
	convertCmd.Flags().StringVar(&spin_flavor, "spin_flavor", "ARMORY", "Select Spinnaker Operator flavor to use (ARMORY, OSS)")
}

func getProfiles(profilesDir string, currentSubDir string) (map[string]string, error) {
	config_files := make(map[string]string)

	profiles_files, err := fileio.GetListOfFiles(profilesDir)
	if err != nil {
		return nil, err
	}

	for _, profiles_files := range profiles_files {

		isDir, err := fileio.IsDirectory(profilesDir + "/" + profiles_files.Name())
		if isDir {
			config_maps, err := getProfiles(profilesDir+"/"+profiles_files.Name(), currentSubDir+"/"+profiles_files.Name())
			if err != nil {
				return nil, err
			}

			for fileName, config_map := range config_maps {
				config_files[fileName] = config_map
				// fmt.Println("fileName:" + fileName + " config_file: " + config_file)
			}
			continue
		}

		local_file, err := fileio.ReadFile(profilesDir + "/" + profiles_files.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var buf = string(local_file)

		if "" != currentSubDir {
			config_files[currentSubDir+"/"+profiles_files.Name()] = buf
		} else {
			config_files[profiles_files.Name()] = buf
		}

		fmt.Println("Found " + profiles_files.Name())
	}

	return config_files, nil
}

// Read contents of halconfig and setup profiles patches for each service.
func migrator(halconfig_dir string, output_dir string, deployment_dir string, spin_flavor string) {

	// Halconfig stuff
	halconfig_file, err := fileio.FindFileInDir(halconfig_dir, "config")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	halyard, err := parser.ParseHalConfig(halconfig_file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Account: " + halyard.PrimaryAccount)
	fmt.Printf("\nHalyard: %#v \n\n", halyard)

	// Profiles stuff
	fmt.Println("Reading " + halconfig_dir + "/" + deployment_dir + "/profiles")

	config_files, err = getProfiles(halconfig_dir+"/"+deployment_dir+"/profiles", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// for fileName, config_file := range config_files {
	// 	fmt.Println("fileName:" + fileName + "\nconfig_file:" + config_file)
	// }

	// profile_settings :=

	// fmt.Println(halconfig)
	// fmt.Println(output)

	// Create output files
	output_config(output_dir, spin_flavor)

}

func output_config(output_dir string, spin_flavor string) {
	KustomizeData := migrate.Kustomize{Spin_flavor: spin_flavor, Output_dir: output_dir}

	err := fileio.EnsureDirectory(output_dir)
	if err != nil {
		log.Println(err)
	}

	err = KustomizeData.CreateKustomization()
	if err != nil {
		log.Fatal(err)
	}
}
