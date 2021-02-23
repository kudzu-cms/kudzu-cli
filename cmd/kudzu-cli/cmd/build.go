package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	debug bool
)

var buildCmd = &cobra.Command{
	Use:   "build <component>",
	Args:  cobra.ExactArgs(1),
	Short: "Builds components and plugins",
}

var buildPluginsCmd = &cobra.Command{
	Use:   "plugins [flags]",
	Args:  cobra.NoArgs,
	Short: "Builds content plugins",
	RunE: func(cmd *cobra.Command, args []string) error {

		info, statErr := os.Stat("plugins")
		if os.IsNotExist(statErr) || !info.IsDir() {
			log.Println("No Plugins to build")
			return nil
		}

		info, statErr = os.Stat(".plugins")

		cmdArgs := []string{info.Name()}
		_, file, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(file)
		soBuildCmd := exec.Command(filepath.Join(basepath, "..", "..", "..", "scripts", "build-plugins.sh"), cmdArgs...)
		soBuildCmd.Dir = "./plugins"
		output, err := soBuildCmd.Output()
		if err != nil {
			fmt.Println(string(output))
			return err
		}
		log.Println(string(output))

		return err
	},
}

func init() {
	buildPluginsCmd.Flags().BoolVar(&debug, "debug", false, "build plugins with debugging flags")
	buildCmd.AddCommand(buildPluginsCmd)
	rootCmd.AddCommand(buildCmd)
}
