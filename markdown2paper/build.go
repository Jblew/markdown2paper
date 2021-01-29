package main

import (
	"fmt"
)

// BuildParams — parameters for building
type BuildParams struct {
	BibFile string
	OutlineFile string
	OutFile string
}

// Build actually builds the paper
func Build(params BuildParams) error {
  outMarkdown := MarkdownSection{}

	outlineMarkdown, err := extractMarkdownSectionFromFile(params.OutlineFile, "Outline")
	if err != nil {
		return err
	}

	outMarkdown.Sections = append(outMarkdown.Sections, MarkdownSection{Title:"A final paper", Sections: []MarkdownSection{outlineMarkdown}})

  return WriteTextToFile(params.OutFile, MarkdownToText(outMarkdown, 0))
}

func extractMarkdownSectionFromFile(path string, sectionTitle string) (MarkdownSection, error) {
	markdown, err := loadMarkdownFromFile(path)
	if err != nil {
		return MarkdownSection{}, err
	}

	found, section := FindTopFirstSectionWithTitle(markdown, sectionTitle)
	if found == true {
		return section, nil
	}
	return MarkdownSection{}, fmt.Errorf("Section named \"%s\" not found in %s", sectionTitle, path)
}


func loadMarkdownFromFile(path string) (MarkdownSection, error) {
	outlineContents, err := ReadFileToText(path)
  if err != nil {
    return MarkdownSection{}, err
  }
  section, err := ParseTextToMarkdown("", outlineContents, 0)
  if err != nil {
    return MarkdownSection{}, err
	}
	return section, nil
}
