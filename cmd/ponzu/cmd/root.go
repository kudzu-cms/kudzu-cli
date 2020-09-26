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

	// for ponzu internal / core development
	gocmd string
	fork  string
	dev   bool

	year = fmt.Sprintf("%d", time.Now().Year())
)

// RootCmd is the main CLI command.
var RootCmd = &cobra.Command{
	Use: "ponzu",
	Long: `Ponzu is an open-source HTTP server framework and CMS, released under
the BSD-3-Clause license.
(c) 2016 - ` + year + ` Boss Sauce Creative, LLC`,
}

// Execute adds all child commands.
func Execute() {

	pflags := RootCmd.PersistentFlags()
	pflags.StringVar(&gocmd, "gocmd", "go", "custom go command if using beta or new release of Go")

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
