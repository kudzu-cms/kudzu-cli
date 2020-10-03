package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Prints the version of kudzu your project is using.",
	Long: `Prints the version of kudzu your project is using. Must be called from
within a kudzu project directory.`,
	Example: `$ kudzu version
> kudzu v0.8.2
(or)
$ kudzu version --cli
> kudzu v0.9.2`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := version(cli)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "kudzu-cli %s\n", p["version"])
	},
}

func version(isCLI bool) (map[string]interface{}, error) {
	kv := make(map[string]interface{})

	info := filepath.Join("cmd", "kudzu-cli", "kudzu.json")
	if isCLI {
		info = filepath.Join("cmd", "kudzu", "kudzu.json")
	}

	b, err := ioutil.ReadFile(info)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &kv)
	if err != nil {
		return nil, err
	}

	return kv, nil
}

func init() {
	versionCmd.Flags().BoolVar(&cli, "cli", false, "specify that information should be returned about the CLI, not project")
	rootCmd.AddCommand(versionCmd)
}
