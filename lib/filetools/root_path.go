package filetools

import (
	"errors"
	"os"
	"path/filepath"
)

func GetRootPath() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for dir != "/" {
		_, err := os.Stat(dir + "/.git")
		if err == nil {
			return dir, nil
		}
		dir, err = filepath.Abs(dir + "/..")
		if err != nil {
			return "", err
		}
	}
	return "", errors.New("not part of a git repo")
}
