package cmd

import (
	"io/ioutil"
	"net/http"
)

func getTemplate(name string) (string, error) {

	remoteTmpl := "https://raw.githubusercontent.com/kudzu-cms/kudzu-cli/" + Version + "/cmd/kudzu-cli/templates/gen-content.tmpl"
	resp, err := http.Get(remoteTmpl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
