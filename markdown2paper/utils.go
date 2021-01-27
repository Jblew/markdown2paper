package main

import "io/ioutil"

// ReadFileToText reads file to text
func ReadFileToText(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
