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
		path := ""
		if len(args) == 2 {
			projectName = args[0]
			path = args[1]
		} else {
			msg := "Please provide a project name and path."
			msg += "\nThis will create a directory within your $GOPATH/src."
			return fmt.Errorf("%s", msg)
		}
		return newProjectInDir(projectName, path)
	},
}

func newProjectInDir(name string, path string) error {
	projPath := path + "/" + name
	return createProjectInDir(projPath)
}

func createProjectInDir(path string) error {

	// create the directory or overwrite it
	err := os.MkdirAll(path+"/content", os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	cmd := exec.Command("kudzu-cli", "gen", "content", "page", "title:string")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Create(path + "/main.go")
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

	cmd = exec.Command("go", "mod", "init", "kudzu-project")
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
	newCmd.Flags().StringVar(&fork, "fork", "", "modify repo source for kudzu core development")
	newCmd.Flags().BoolVar(&dev, "dev", false, "modify environment for kudzu core development")

	rootCmd.AddCommand(newCmd)
}
