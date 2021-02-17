package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Args:    cobra.NoArgs,
	Aliases: []string{"v"},
	Short:   "Prints kudzu-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "kudzu-cli %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
