package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Retrieves the given template from the remote repo for the running version of
// the CLI.
// @todo Consider setting up a local cache under ~/.kudzu/<version>/templates
func getTemplateFromRepo(name string) (string, error) {

	repoBase := "https://raw.githubusercontent.com/kudzu-cms/kudzu-cli"
	versionPath := Version + "/cmd/kudzu-cli/templates/" + name
	remoteTmpl := repoBase + "/" + versionPath
	resp, err := http.Get(remoteTmpl)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("Request for template " + name + " failed with: " + resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
