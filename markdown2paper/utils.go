package main

import (
	"io/ioutil"
	"os"
)

// ReadFileToText reads file to text
func ReadFileToText(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

// WriteTextToFile writes text to a file
func WriteTextToFile(path string, contents string) error {
	return ioutil.WriteFile(path, []byte(contents), os.ModeAppend | os.FileMode(0777))
}
