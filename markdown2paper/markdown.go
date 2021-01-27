package main

import "strings"

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
	sections, err := contentsAndHeadingsToSections(contentsAndHeadings, baseLevel + 1)
	if err != nil {
		return MarkdownSection{}, err
	}

	return MarkdownSection{
		Title: title,
		Content: contentAboveHeadings,
		Sections: sections,
	}, nil
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
	i := 1
	sections := []MarkdownSection{}
	for i < len(contentsAndHeadings) {
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
