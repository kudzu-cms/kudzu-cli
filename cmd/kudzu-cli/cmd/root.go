package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	bind      string
	httpsport int
	port      int
	docsport  int
	https     bool
	devhttps  bool
	docs      bool
	cli       bool

	fork string
	dev  bool

	year = fmt.Sprintf("%d", time.Now().Year())
)

// Version is the kudzu-cli version.
const Version = "master"

var rootCmd = &cobra.Command{
	Use: "kudzu-cli",
	Long: `kudzu-cli is the command-line interface for interacting with Kudzu projects.
(c) 2016 - ` + year + ` Boss Sauce Creative, LLC`,
}

// Execute adds all child commands.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
