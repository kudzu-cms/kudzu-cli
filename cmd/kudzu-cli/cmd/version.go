package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "prints kudzu-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "kudzu-cli %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
