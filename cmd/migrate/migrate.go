package migrate

import (
	"fmt"
	"log"
	"os"

	"github.com/austinthao5/golang_proto_test/internal/fileio"
	"github.com/austinthao5/golang_proto_test/internal/migrate"
	"github.com/austinthao5/golang_proto_test/internal/migrate/structs"
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

	// fmt.Println("Account: " + halyard.PrimaryAccount)
	fmt.Printf("\nHalyard: %#v \n\n", halyard)
	// fmt.Println("Aws.AccessKey: " + halyard.DeploymentConfigurations.providers.Aws.AccessKey)
	// fmt.Println("Aws.AccessKey: " + halyard.DeploymentConfiguration[0].Name)

	// Profiles stuff
	fmt.Println("Reading " + halconfig_dir + "/" + deployment_dir + "/profiles")

	profiles_config_files, err = fileio.GetFilesAndData(halconfig_dir+"/"+deployment_dir+"/profiles", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// for fileName, config_file := range profiles_config_files {
	// 	fmt.Println("fileName:" + fileName + "\nconfig_file:" + config_file)
	// }

	service_settings_config_files, err = fileio.GetFilesAndData(halconfig_dir+"/"+deployment_dir+"/service-settings", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// profile_settings :=

	// fmt.Println(halconfig)
	// fmt.Println(output)

	//Create Kustomize struct with current info
	KustomizeData := structs.Kustomize{
		Spin_flavor:                spin_flavor,
		Output_dir:                 output_dir,
		Halyard:                    halyard,
		ProfilesConfigFiles:        profiles_config_files,
		ServiceSettingsConfigFiles: service_settings_config_files,
	}
	// Create output files
	output_config(&KustomizeData)
}

func output_config(KustomizeData *structs.Kustomize) {

	err := fileio.EnsureDirectory(output_dir)
	if err != nil {
		log.Println(err)
	}

	err = migrate.CreateKustomization(KustomizeData)
	if err != nil {
		log.Fatal(err)
	}
}
