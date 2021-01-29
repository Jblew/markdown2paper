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
  outMarkdown := MarkdownSection{Sections: []MarkdownSection{{}}}

	outlineMarkdown, err := loadMarkdownFromFile(params.OutlineFile)
	if err != nil {
		return err
	}

	outlineSectionMarkdown, err := extractMarkdownSection(outlineMarkdown, "Outline")
	if err != nil {
		return err
	}

	outlineTree := ParseTextToOutlineTree(outlineSectionMarkdown.Content, 0)
	paperContents := buildPaperContentsForOutline(outlineTree)

	outMarkdown.Sections[0].Sections = paperContents
	outMarkdown.Sections[0].Title = outlineMarkdown.Sections[0].Title
	textOut := MarkdownToText(outMarkdown, 0)
  return WriteTextToFile(params.OutFile, textOut)
}

func buildPaperContentsForOutline(outlineTree OutlineTree) []MarkdownSection {
	sections := []MarkdownSection{}
	for _, child := range outlineTree.Children {
		section := MarkdownSection{
			Title: child.Title,
			Sections: buildPaperContentsForOutline(child),
		}
		sections = append(sections, section)
	}
	return sections
}

func extractMarkdownSectionFromFile(path string, sectionTitle string) (MarkdownSection, error) {
	markdown, err := loadMarkdownFromFile(path)
	if err != nil {
		return MarkdownSection{}, err
	}

	return extractMarkdownSection(markdown, sectionTitle)
}


func extractMarkdownSection(markdown MarkdownSection, sectionTitle string) (MarkdownSection, error) {
	found, section := FindTopFirstSectionWithTitle(markdown, sectionTitle)
	if found == true {
		return section, nil
	}
	return MarkdownSection{}, fmt.Errorf("Section named \"%s\" not found", sectionTitle)
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
