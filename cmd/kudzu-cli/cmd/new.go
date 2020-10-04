package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new <path> <module name>",
	Short:   "creates a project directory of the name supplied as a parameter",
	Example: `$ kudzu new ~/Code/go/kudzu-project github.com/bobbygryzynger/kudzu-project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := ""
		modname := ""
		if len(args) == 2 {
			path = args[0]
			modname = args[1]
		} else {
			msg := "Please provide a path and module name"
			return fmt.Errorf("%s", msg)
		}
		return createProjectInDir(path, modname)
	},
}

func createProjectInDir(path string, modname string) error {

	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return fmt.Errorf("The path %s already exists", path)
	}

	// @todo if an error occurs during project creation, clean up the project
	// directory
	err = os.MkdirAll(filepath.Join(path, "content"), os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("kudzu-cli", "gen", "content", "page", "title:string")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(path, "main.go"))
	defer file.Close()
	if err != nil {
		return err
	}

	tmplStr, err := getTemplate("gen-new-project-main.tmpl")
	if err != nil {
		return err
	}

	_, err = file.WriteString(tmplStr)
	if err != nil {
		return err
	}

	cmd = exec.Command("go", "mod", "init", modname)
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("New kudzu project created at", path)
	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
}
