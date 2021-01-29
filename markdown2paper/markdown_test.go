package main
import "testing"

func TestSplitByHeadings(t *testing.T) {
	got := splitByHeadings("# H1\n\nC1\n\n#H2\n\nC2", 1)
	if len(got) != 5 {
			t.Errorf("Should split into 5 sections, got %d, %+v", len(got), got)
			return
	}

	i := 0
	want := ""
	if got[i] != want {
		t.Errorf("got[%d] = %s, want \"%s\"", i, got[i], want)
	}

	i = 1
	want = "H1"
	if got[i] != want {
		t.Errorf("got[%d] = %s, want \"%s\"", i, got[i], want)
	}

	i = 2
	want = "C1"
	if got[i] != want {
		t.Errorf("got[%d] = %s, want \"%s\"", i, got[i], want)
	}

	i = 3
	want = "H2"
	if got[i] != want {
		t.Errorf("got[%d] = %s, want \"%s\"", i, got[i], want)
	}

	i = 4
	want = "C2"
	if got[i] != want {
		t.Errorf("got[%d] = %s, want \"%s\"", i, got[i], want)
	}
}


func TestGetLineHeadingLevel(t *testing.T) {
	got := getLineHeadingLevel("Text")
	if got != -1 {
		t.Errorf("getLineHeadingLevel(Text) = %d, want -1", got)
	}

	got = getLineHeadingLevel("")
	if got != -1 {
		t.Errorf("getLineHeadingLevel() = %d, want -1", got)
	}

	got = getLineHeadingLevel("# Text")
	if got != 1 {
		t.Errorf("getLineHeadingLevel(# Text) = %d, want 1", got)
	}

	got = getLineHeadingLevel("### Text")
	if got != 3 {
		t.Errorf("getLineHeadingLevel(### Text) = %d, want 3", got)
	}
}

func TestGetLineHeadingText(t *testing.T) {
	got := getLineHeadingText("Text")
	if got != "Text" {
		t.Errorf("getLineHeadingLevel(Text) = %s, want \"Text\"", got)
	}

	got = getLineHeadingText("## Text")
	if got != "Text" {
		t.Errorf("getLineHeadingLevel(## Text) = %s, want \"Text\"", got)
	}

	got = getLineHeadingText("# Text")
	if got != "Text" {
		t.Errorf("getLineHeadingLevel(# Text) = %s, want \"Text\"", got)
	}
}

func TestParseTextToMarkdown(t *testing.T) {
	got, err := ParseTextToMarkdown("doc.md", "Pre-title \n# Title\n\n Level 1 comment \n\n## Section A\n Section A text \n\n ## Section B\n\n\nSection B text\n", 0)
	if err != nil {
		t.Errorf("ParseTextToMarkdown errpr: %+v", err)
	}

	if got.Title != "doc.md" {
		t.Errorf("Invalid root title: %+v", got)
	}

	if got.Content != "Pre-title" {
		t.Errorf("Invalid root content: %+v", got)
	}

	if len(got.Sections) != 1 {
		t.Errorf("Invalid number of level 1 sections: %+v", got.Sections)
		return
	}

	topSection := got.Sections[0]

	if topSection.Title != "Title" {
		t.Errorf("Invalid level 1 section title: %+v", topSection.Title)
	}

	if topSection.Content != "Level 1 comment" {
		t.Errorf("Invalid level 1 section content: %+v", topSection.Content)
	}

	if len(topSection.Sections) != 2 {
		t.Errorf("Invalid number of level 2 sections: %+v", topSection.Sections)
		return
	}

	if topSection.Sections[0].Title != "Section A" {
		t.Errorf("Invalid level 2 section title: %+v", topSection.Sections[0].Title)
	}

	if topSection.Sections[0].Content != "Section A text" {
		t.Errorf("Invalid level 1 section content: %+v", topSection.Sections[0].Content)
	}

	if topSection.Sections[1].Title != "Section B" {
		t.Errorf("Invalid level 2 section title: %+v", topSection.Sections[1].Title)
	}

	if topSection.Sections[1].Content != "Section B text" {
		t.Errorf("Invalid level 1 section content: %+v", topSection.Sections[1].Content)
	}
}
