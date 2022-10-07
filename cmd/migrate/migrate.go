package migrate

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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

			migrator(halconfig_dir, output_dir, deployment_dir)

		}
	},
}

// Run this function when the CLI starts up. This initializes the flags for the Convert command.
func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVar(&halconfig_dir, "halconfig", "~/.hal", "Provide the Halconfig directory to convert")
	convertCmd.Flags().StringVar(&output_dir, "output", "./operatorConfig", "Select an output directory")
	convertCmd.Flags().StringVar(&deployment_dir, "deployment", "default", "Select the deployment name being used by Halyard. This is the subdirectory in the ./hal folder where the profiles and service-settings live")
}

// Read contents of halconfig and setup profiles patches for each service.
func migrator(halconfig_dir string, output_dir string, deployment_dir string) {

	// Halconfig stuff
	dir, err := os.Open(halconfig_dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reading " + halconfig_dir)

	for _, file := range files {
		// if file.Mode().IsRegular() {
		// 	fmt.Println(file.Name(), file.Size(), "bytes")
		// }

		if file.Name() == "config" {
			fmt.Println("Found Halonfig!")

			config, err := os.Open(halconfig_dir + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer config.Close()
			var buf bytes.Buffer
			io.Copy(&buf, config)
			halconfig = buf.String()
		}

	}

	// Profiles stuff
	fmt.Println("Reading " + halconfig_dir + "/" + deployment_dir + "/profiles")
	profile_folder, err := os.Open(halconfig_dir + "/" + deployment_dir + "/profiles")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer profile_folder.Close()

	profiles_files, err := profile_folder.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, profiles_files := range profiles_files {
		// fmt.Println(profiles_files.Name(), profiles_files.Size(), "bytes")

		local_file, err := os.Open(halconfig_dir + "/" + deployment_dir + "/profiles/" + profiles_files.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer local_file.Close()
		var buf bytes.Buffer
		io.Copy(&buf, local_file)
		config_files := make(map[string]string)
		config_files[profiles_files.Name()] = buf.String()
		fmt.Println("Found " + profiles_files.Name())

	}
	// profile_settings :=

	// fmt.Println(halconfig)
	// fmt.Println(output)

	// Create output files
	output_config(output_dir)

}

func output_config(output_dir string) {

	if _, err := os.Stat(output_dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(output_dir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	kustomization, err := os.Create(output_dir + "/kustomization.yml")
	if err != nil {
		log.Fatal(err)
	}

	str := `resources:
	- SpinnakerService.yml
	
	# Apply the patches top down order
	patchesStrategicMerge:`

	data := []byte(str)

	kustomization.Write(data)

	base_config, err := os.Create(output_dir + "/SpinnakerService.yml")
	if err != nil {
		log.Fatal(err)
	}

	str = `apiVersion: spinnaker.armory.io/v1alpha2
	kind: SpinnakerService
	metadata:
		name: spinnaker
	spec:
		spinnakerConfig:
		config:
			version: 0.0.0   # the version of Spinnaker to be deployed, see config-patch.yml
			persistentStorage:
			persistentStoreType: s3
			s3:
				bucket: mybucket
				rootFolder: front50
		profiles:
			clouddriver:
			artifacts:
				http:
				enabled: true
				accounts: []
			deck:
			settings-local.js: |
				window.spinnakerSettings.feature.kustomizeEnabled = true;
				window.spinnakerSettings.feature.artifactsRewrite = true;
			echo: {}    # is the contents of ~/.hal/default/profiles/echo.yml
			fiat: {}    # is the contents of ~/.hal/default/profiles/fiat.yml
			front50: {} # is the contents of ~/.hal/default/profiles/front50.yml
			gate: {}    # is the contents of ~/.hal/default/profiles/gate.yml
			igor: {}    # is the contents of ~/.hal/default/profiles/igor.yml
			kayenta: {} # is the contents of ~/.hal/default/profiles/kayenta.yml
			orca: {}    # is the contents of ~/.hal/default/profiles/orca.yml
			rosco: {}   # is the contents of ~/.hal/default/profiles/rosco.yml
	
		service-settings:
			clouddriver: {}
			deck: {}
			echo: {}
			fiat: {}
			front50: {}
			gate: {}
			igor: {}
			kayenta: {}
			orca: {}
			rosco: {}
		files: {}

		expose:
		type: service
		service:
			type: LoadBalancer
			overrides: {}
	`

	data = []byte(str)

	base_config.Write(data)
}