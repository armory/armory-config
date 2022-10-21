package migrate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	halconfig_dir         string
	output_dir            string
	deployment_dir        string
	spin_flavor           string
	halconfig             string
	profiles_config_files map[string]string
)

var rootCmd = &cobra.Command{
	Use:   "armory-config",
	Short: "migrate - a simple CLI to transform and inspect strings",
	Long: `Halyard to Operator migrator.
migrate some stuff brah`,
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
