package filetools

import (
	"io/ioutil"
	"regexp"
)

func FindLabels(path string) (*[]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile(`(?m:get_label\(['"]([^'"]*)['"]\))`)
	if err != nil {
		return nil, err
	}
	results := re.FindAllStringSubmatch(string(content), len(content))

	keysOnly := []string{}
	for _, res := range results {
		keysOnly = append(keysOnly, res[1])
	}
	return &keysOnly, nil
}
