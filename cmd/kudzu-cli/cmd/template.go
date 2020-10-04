package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func getTemplate(name string) (string, error) {

	remoteTmpl := "https://raw.githubusercontent.com/kudzu-cms/kudzu-cli/" + Version + "/cmd/kudzu-cli/templates/" + name
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
