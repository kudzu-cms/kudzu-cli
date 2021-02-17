package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
		err := filepath.Walk(filepath.Join(".", "plugins"), func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {

				outPath := filepath.Join(".", ".plugins") + strings.TrimLeft(path, filepath.Join(".", "plugins"))
				outPath = strings.TrimRight(outPath, ".go") + ".so"

				buildDir := strings.TrimRight(outPath, info.Name()+".so")
				_, err := os.Stat(buildDir)
				if os.IsNotExist(err) {
					err = os.MkdirAll(buildDir, os.ModeDir|os.ModePerm)
					if err != nil {
						return err
					}
				}

				cmdArgs := []string{"build", "-buildmode=plugin"}
				// Add debugging build flags.
				// @see https://github.com/go-delve/delve/issues/865#issuecomment-480766102
				if debug {
					cmdArgs = append(cmdArgs, "-gcflags='all=-N -l'")
				}
				cmdArgs = append(cmdArgs, "-o", outPath, path)
				soBuildCmd := exec.Command("go", cmdArgs...)
				log.Println("Plugin: " + info.Name())
				log.Println("\tBuilding: " + soBuildCmd.String())
				err = soBuildCmd.Run()
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
			return nil
		})
		return err
	},
}

func init() {
	buildPluginsCmd.Flags().BoolVar(&debug, "debug", false, "build plugins with debugging flags")
	buildCmd.AddCommand(buildPluginsCmd)
	rootCmd.AddCommand(buildCmd)
}
