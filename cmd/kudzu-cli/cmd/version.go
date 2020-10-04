package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Prints the version of kudzu-cli your project is using.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "kudzu-cli %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
