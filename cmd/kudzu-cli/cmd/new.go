package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [flags] <project name>",
	Short: "creates a project directory of the name supplied as a parameter",
	Long: `Creates a project directory of the name supplied as a parameter
immediately following the 'new' option in the $GOPATH/src directory. Note:
'new' depends on the program 'git' and possibly a network connection. If
there is no local repository to clone from at the local machine's $GOPATH,
'new' will attempt to clone the 'github.com/kudzu-cms/kudzu' package from
over the network.`,
	Example: `$ kudzu new github.com/nilslice/proj
> New kudzu project created at $GOPATH/src/github.com/nilslice/proj`,
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := "kudzu-project"
		if len(args) > 0 {
			projectName = args[0]
		} else {
			msg := "Please provide a project name."
			msg += "\nThis will create a directory within your $GOPATH/src."
			return fmt.Errorf("%s", msg)
		}
		return newProjectInDir(projectName, args[1])
	},
}

func newProjectInDir(name string, path string) error {
	projPath := path + "/" + name
	return createProjectInDir(projPath)
}

func createProjectInDir(path string) error {

	// create the directory or overwrite it
	err := os.MkdirAll(path, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("kudzu-cli", "gen", "page", "title:\"string\"")
	cmd.Dir = path
	cmd.Run()

	cmd = exec.Command("go", "mod", "init", "kudzu-project")
	cmd.Dir = path
	cmd.Run()

	cmd = exec.Command("go", "mod", "tidy", "kudzu-project")
	cmd.Dir = path
	cmd.Run()

	fmt.Println("New kudzu project created at", path)
	return nil
}

func init() {
	newCmd.Flags().StringVar(&fork, "fork", "", "modify repo source for kudzu core development")
	newCmd.Flags().BoolVar(&dev, "dev", false, "modify environment for kudzu core development")

	rootCmd.AddCommand(newCmd)
}
