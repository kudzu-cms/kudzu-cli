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

var rootCmd = &cobra.Command{
	Use: "ponzu",
	Long: `Ponzu is an open-source HTTP server framework and CMS, released under
the BSD-3-Clause license.
(c) 2016 - ` + year + ` Boss Sauce Creative, LLC`,
}

// Execute adds all child commands.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
