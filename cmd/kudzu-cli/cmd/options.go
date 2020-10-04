package cmd

import (
	"fmt"
	"strings"
)

func getAnswer() (string, error) {
	var answer string
	_, err := fmt.Scanf("%s\n", &answer)
	if err != nil {
		if err.Error() == "unexpected newline" {
			answer = ""
		} else {
			return "", err
		}
	}

	answer = strings.ToLower(answer)

	return answer, nil
}
