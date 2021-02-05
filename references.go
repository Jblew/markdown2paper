package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ProcessPandocReferences replaces all occurances of pandoc citation keys [@key] with footnotes
func ProcessPandocReferences(sections []MarkdownSection, bibliography Bibliography) []MarkdownSection {
	citationKeys := getCitationKeysOfSections(sections)
	bibliographyContent := ""

	for i, key := range citationKeys {
		bibliographyContent += fmt.Sprintf("[^%d]: %s\n\n", (i+1), bibliography.FormatMarkdownByKey(key))
	}

	transformedSections := applyFootnotesToSections(sections, citationKeys)
	bibliographySection := MarkdownSection{
		Title: "Bibliography",
		Content: bibliographyContent,
	}
	return append(transformedSections, bibliographySection)
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

func applyFootnotesToSections(sections []MarkdownSection, keys []string) []MarkdownSection {
	return applyFootnotesToSection(MarkdownSection{ Sections: sections }, keys).Sections
}

func applyFootnotesToSection(section MarkdownSection, keys []string) MarkdownSection {
	childSectionsTransformed := []MarkdownSection{}
	for _, childSection := range section.Sections {
		childSectionsTransformed = append(childSectionsTransformed, applyFootnotesToSection(childSection, keys))
	}

	return MarkdownSection{
		Title: applyFootnotesToText(section.Title, keys),
		Content: applyFootnotesToText(section.Content, keys),
		Sections: childSectionsTransformed,
	}
}

func applyFootnotesToText(text string, keys []string) string {
	out := text
	for i, key := range keys {
		footnote := fmt.Sprintf("[^%d]", i+1)
		searchKey := fmt.Sprintf("[@%s]", key)
		out = strings.ReplaceAll(out, searchKey, footnote)
	}
	return out
}

var citationKeyRe = regexp.MustCompile(`(?m)\[@([^\]]+)\]`)
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
