package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

		info, statErr := os.Stat("plugins")
		if os.IsNotExist(statErr) || !info.IsDir() {
			log.Println("No Plugins to build")
			return nil
		}

		// Clean previously compiled plugins.
		info, statErr = os.Stat(".plugins")
		if !os.IsNotExist(statErr) && info.IsDir() {
			os.RemoveAll(".plugins")
		}

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

				outpathAbs, _ := filepath.Abs(outPath)
				sourcePathAbs, _ := filepath.Abs(path)
				cmdArgs := []string{filepath.Dir(sourcePathAbs), info.Name(), outpathAbs}
				if debug {
					cmdArgs = append(cmdArgs, "true")
				}
				_, callerFile, _, _ := runtime.Caller(0)
				callerBase := filepath.Dir(callerFile)
				rootPath := filepath.Join(callerBase, "..", "..", "..")
				soBuildCmd := exec.Command(filepath.Join(rootPath, "scripts", "build-plugins.sh"), cmdArgs...)
				soBuildCmd.Dir = "./plugins"
				output, err := soBuildCmd.CombinedOutput()
				if err != nil {
					fmt.Println(string(output))
					return err
				}
				log.Println(string(output))
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
