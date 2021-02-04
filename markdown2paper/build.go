package main

import (
	"fmt"
	"path/filepath"
)

// Build actually builds the paper
func Build(params Config) error {
	outMarkdown := MarkdownSection{Sections: []MarkdownSection{{}}}
	baseDir := filepath.Dir(params.OutlineFile)

	outlineMarkdown, err := loadMarkdownFromFile(params.OutlineFile)
	if err != nil {
		return err
	}

	outlineSectionMarkdown, err := extractMarkdownSection(outlineMarkdown, "Outline")
	if err != nil {
		return err
	}

	outlineTree := ParseTextToOutlineTree(outlineSectionMarkdown.Content, 0)
	paperContents, err := buildPaperContentsForOutline(outlineTree, baseDir)
	if err != nil {
		return err
	}

	paperContentsWithBibliography := ProcessPandocReferences(paperContents)

	outMarkdown.Sections[0].Sections = paperContentsWithBibliography
	outMarkdown.Sections[0].Title = outlineMarkdown.Sections[0].Title
	textOut := MarkdownToText(outMarkdown, 0)
  return WriteTextToFile(params.OutFile, textOut)
}

func buildPaperContentsForOutline(outlineTree OutlineTree, baseDir string) ([]MarkdownSection, error) {
	sections := []MarkdownSection{}
	for _, child := range outlineTree.Children {
		childContent, loadedSections, err := fetchOutlineItemContent(child, baseDir)
		if err != nil {
			return []MarkdownSection{}, err
		}
		childSections, err := buildPaperContentsForOutline(child, baseDir)
		if err != nil {
			return []MarkdownSection{}, err
		}
		section := MarkdownSection{
			Title: child.Title,
			Content: childContent,
			Sections: append(loadedSections, childSections...),
		}
		sections = append(sections, section)
	}
	return sections, nil
}

func fetchOutlineItemContent(item OutlineTree, baseDir string) (string, []MarkdownSection, error) {
	if len(item.LinkPath) == 0 {
		return "", []MarkdownSection{}, nil
	}
	path, err := filepath.Abs(filepath.Join(baseDir, item.LinkPath))
	if err != nil {
		return "", []MarkdownSection{}, err
	}
	rootSection, err := extractMarkdownSectionFromFile(path, "Final")
	if err != nil {
		return "", []MarkdownSection{}, err
	}
	return rootSection.Content ,rootSection.Sections, nil
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
