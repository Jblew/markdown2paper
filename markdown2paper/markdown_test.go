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
