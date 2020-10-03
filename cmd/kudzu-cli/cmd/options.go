package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var kudzuRepo = []string{"github.com", "bobbygryzynger", "kudzu"}

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

func vendorCorePackages(path string) error {
	vendorPath := filepath.Join(path, "cmd", "kudzu", "vendor", "github.com", "bobbygryzynger", "kudzu")
	err := os.MkdirAll(vendorPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	dirs := []string{"content", "management", "system"}
	for _, dir := range dirs {
		err = os.Rename(filepath.Join(path, dir), filepath.Join(vendorPath, dir))
		if err != nil {
			return err
		}
	}

	// create a user content directory at project root
	contentPath := filepath.Join(path, "content")
	err = os.Mkdir(contentPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func copyFileNoRoot(src, dst string) error {
	noRoot := strings.Split(src, string(filepath.Separator))[1:]
	path := filepath.Join(noRoot...)
	dstFile, err := os.Create(filepath.Join(dst, path))
	defer dstFile.Close()
	if err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	defer srcFile.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func copyAll(src, dst string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		sep := string(filepath.Separator)

		// base == the kudzu project dir + string(filepath.Separator)
		parts := strings.Split(src, sep)
		base := strings.Join(parts[:len(parts)-1], sep)
		base += sep

		target := filepath.Join(dst, path[len(base):])

		// if its a directory, make dir in dst
		if info.IsDir() {
			err := os.MkdirAll(target, os.ModeDir|os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// if its a file, move file to dir of dst
			err = os.Rename(path, target)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
