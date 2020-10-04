package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Prints the version of kudzu your project is using.",
	Long:    `Prints the version of kudzu your project is using.`,
	Example: `$ kudzu version
> kudzu v0.8.2
(or)
$ kudzu version
> kudzu v0.9.2`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "kudzu-cli %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
