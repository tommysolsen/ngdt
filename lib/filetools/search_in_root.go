package filetools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetConfigJSON() (*ApiJSON, error) {
	file, err := FindInProject("api.json", ".", "config")
	if err != nil {
		return nil, err
	}
	// Read file if found
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var apiData ApiJSON
	err = json.Unmarshal(data, &apiData)
	if err != nil {
		return nil, err
	}
	return &apiData, err
}

// FindInProject searches a series of files in the project. And returns the first match it finds for the file
func FindInProject(filename string, paths ...string) (string, error) {
	root, err := GetRootPath()
	if err != nil {
		return "", err
	}

	resolvedPaths := make([]string, len(paths))
	for i, path := range paths {

		resolvedPaths[i] = filepath.Clean(strings.Join([]string{root, path, filename}, "/"))
	}
	for _, path := range resolvedPaths {
		fInfo, _ := os.Stat(path)
		if fInfo != nil {
			if fInfo.IsDir() == false {
				return path, nil
			}
		}
	}

	return "", fmt.Errorf("no error found for filename %s, looked in [%s]", filename, strings.Join(resolvedPaths, ", "))
}

func FindFilesOfType(extension string, baseDir string) []string {
	files := make([]string, 0)
	var walkerFunction = func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, "/vendor/") {
			return nil
		}
		if strings.Contains(path, "/storage/") {
			return nil
		}
		if strings.Contains(path, "/files/") {
			return nil
		}
		if !strings.Contains(path, "."+extension) {
			return nil
		}
		files = append(files, path)
		return nil
	}
	filepath.Walk(baseDir, walkerFunction)
	return files
}
