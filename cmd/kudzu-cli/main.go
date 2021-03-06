// Package main is located in the cmd/kudzu-cli directory and contains the code to build
// and operate the command line interface (CLI) to manage kudzu systems. Here,
// you will find the code that is used to create new kudzu projects, generate
// code for content types and other files.
package main

import "github.com/kudzu-cms/kudzu-cli/cmd/kudzu-cli/cmd"

func main() {
	cmd.Execute()
}
