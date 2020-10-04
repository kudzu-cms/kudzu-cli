package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Version is the kudzu-cli version.
const Version = "master"

var rootCmd = &cobra.Command{
	Use:  "kudzu-cli",
	Long: `kudzu-cli is the command-line interface for interacting with Kudzu projects`,
}

// Execute adds all child commands.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
