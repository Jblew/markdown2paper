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

// StringSliceContains checks if slice contains
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

// StringSliceRemoveDuplicates removes duplicates
func StringSliceRemoveDuplicates(s []string) []string {
	out := []string {}
	for _, item := range s {
		if StringSliceContains(out, item) != true {
			out = append(out, item)
		}
	}
	return out
}

