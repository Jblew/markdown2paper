package main

import (
	"fmt"
	"regexp"
)

// ProcessPandocReferences replaces all occurances of pandoc citation keys [@key] with footnotes
func ProcessPandocReferences(sections []MarkdownSection) []MarkdownSection {
	citationKeys := getCitationKeysOfSections(sections)
	bibliographyContent := ""

	for i, key := range citationKeys {
		bibliographyContent += fmt.Sprintf("%d. %s\n\n", (i+1), key)
	}

	bibliography := MarkdownSection{
		Title: "Bibliography",
		Content: bibliographyContent,
	}
	return append(sections, bibliography)
}

func getCitationKeysOfSections(sections []MarkdownSection) []string {
	keys := []string{}
	for _, section := range sections {
		keysInTitle := getCitationKeysFromText(section.Title)
		keys = append(keys, keysInTitle...)

		keysInContent := getCitationKeysFromText(section.Content)
		keys = append(keys, keysInContent...)

		keysInChildren := getCitationKeysOfSections(section.Sections)
		keys = append(keys, keysInChildren...)
	}

	return StringSliceRemoveDuplicates(keys)
}

var citationKeyRe = regexp.MustCompile(`(?m)\[@([a-zA-Z0-9]+)\]`)
func getCitationKeysFromText(text string) []string {
	keys := []string{}
	for _, match := range citationKeyRe.FindAllStringSubmatch(text, -1) {
		if len(match) != 2 {
			continue
		}
		key := match[1]
		if StringSliceContains(keys, key) != true {
			keys = append(keys, key)
		}
	}
	return keys
}
