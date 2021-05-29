package main

import (
	"fmt"
	"strings"
)

// MarkdownSection is a section in a document
type MarkdownSection struct {
	Title string
	Content string
	Sections []MarkdownSection
}

// ParseTextToMarkdown parset text section to MarkdownSection object
func ParseTextToMarkdown(title string, contents string, baseLevel int) (MarkdownSection, error) {
	contentsAndHeadings := splitByHeadings(contents, baseLevel + 1)
	contentAboveHeadings := ""
	if len(contentsAndHeadings) > 0 {
		contentAboveHeadings = contentsAndHeadings[0]
	}

	sections := []MarkdownSection{}
	if len(contentsAndHeadings) > 1 {
		sectionsOfContent, err := contentsAndHeadingsToSections(contentsAndHeadings, baseLevel + 1)
		sections = sectionsOfContent
		if err != nil {
			return MarkdownSection{}, err
		}
	}

	return MarkdownSection{
		Title: title,
		Content: contentAboveHeadings,
		Sections: sections,
	}, nil
}

// MarkdownToText converts markdown to text
func MarkdownToText(section MarkdownSection, level int) string {
	out := ""
	if level > 0 && len(section.Title) > 0 {
		out += makeHeading(section.Title, level) + "\n\n"
	}

	if len(section.Content) > 0 {
		out += section.Content + "\n\n"
	}

	for _, subsection := range section.Sections {
		out += MarkdownToText(subsection, level + 1)
	}

	if level > 0 && len(section.Sections) == 0 {
		out += "\n\n&nbsp;\n\n"
	}

	return out
}

// FindTopFirstSectionWithTitle searches top-down by hierarchy for the top-hierarchy section of a title
func FindTopFirstSectionWithTitle(section MarkdownSection, searchHeading string) (bool, MarkdownSection) {
	if section.Title == searchHeading {
		return true, section
	}
	for _, subsection := range section.Sections {
		found, matchedSection := FindTopFirstSectionWithTitle(subsection, searchHeading)
		if found == true {
			return true, matchedSection
		}
	}
	return false, MarkdownSection{}
}

func splitByHeadings(contents string, level int) []string {
	lines := strings.Split(contents, "\n")
	out := []string{}
	textBuffer := ""
	for _, line := range lines {
		headingLevel := getLineHeadingLevel(line)
		if headingLevel == level {
			textBufferTrimmed := strings.Trim(textBuffer, "\n\r ")
			out = append(out, textBufferTrimmed)
			out = append(out, getLineHeadingText(line))
			textBuffer = ""
		} else {
			textBuffer += line + "\n"
		}
	}
	textBufferTrimmed := strings.Trim(textBuffer, "\n\r ")
	out = append(out, textBufferTrimmed)
	return out
}

func contentsAndHeadingsToSections(contentsAndHeadings []string, level int) ([]MarkdownSection, error) {
	sections := []MarkdownSection{}
	for i:= 1; i < len(contentsAndHeadings);i+=2 {
		heading := contentsAndHeadings[i]
		content := ""
		if (i+1) < len(contentsAndHeadings) {
			content = contentsAndHeadings[i+1]
		}
		section, err := ParseTextToMarkdown(heading, content, level)
		if err != nil {
			return []MarkdownSection{}, err
		}
		sections = append(sections, section)
	}

	return sections, nil
}

func getLineHeadingLevel(line string) int {
	trimmed := strings.TrimSpace(line)
	for i := 6;i >= 1;i-- {
		prefix := strings.Repeat("#", i)
		if strings.HasPrefix(trimmed, prefix) {
			return i
		}
	}
	return -1
}

func getLineHeadingText(line string) string {
	return strings.Trim(line, "# \n")
}

func makeHeading(title string, level int) string {
	return fmt.Sprintf("%s %s", strings.Repeat("#", level), title)
}
