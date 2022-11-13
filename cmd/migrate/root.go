package migrate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flag variables
var (
	halconfig_dir                 string
	output_dir                    string
	override_deployment_dir       string
	spin_flavor                   string
	halconfig                     string
	skip_validations              string
	writefiles                    bool
	profiles_config_files         map[string]string
	service_settings_config_files map[string]string
)

var rootCmd = &cobra.Command{
	Use:   "armory-config",
	Short: "Convert - a simple CLI that converts Halyard configuration to Operator Kustomize format",
	Long: `armory-config takes in the entire hal directory and organizes/restructures the files into
kustomize format. The output should be a fully buildable and deployable set of kustomize files`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Specify your Halyard Directory with the -h flag and the output directory with the -o flag\nFor example:\n$CMD -h ~/.hal -o test/operator/config")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
