package main
import (
	"testing"
)

func TestParseTextToOutline(t *testing.T) {
	got := ParseTextToOutlineTree("1. Background\n   1. Subsection A\n\t1. Subsection B\n2. Methods\n   1. Subsection A\n      1. Subsubsection", 0)

	if len(got.Children) != 2 {
		t.Errorf("Invalid length of root children")
		return
	}

	if got.Children[0].Title != "Background" {
		t.Errorf("Invalid got.Children[0].Title")
	}

	if len(got.Children[0].Children) != 2 {
		t.Errorf("Invalid length of children of got.Children[0]")
		return
	}

	if got.Children[0].Children[0].Title != "Subsection A" {
		t.Errorf("Invalid got.Children[0].Children[0].Title")
	}

	if got.Children[0].Children[1].Title != "Subsection B" {
		t.Errorf("Invalid got.Children[0].Children[1].Title")
	}

	if got.Children[1].Title != "Methods" {
		t.Errorf("Invalid got.Children[1].Title")
	}

	if len(got.Children[1].Children) != 1 {
		t.Errorf("Invalid length of children of got.Children[1]")
		return
	}

	if got.Children[1].Children[0].Title != "Subsection A" {
		t.Errorf("Invalid got.Children[1].Children[0].Title")
	}

	if len(got.Children[1].Children[0].Children) != 1 {
		t.Errorf("Invalid length of children of got.Children[1].Children[0")
		return
	}

	if got.Children[1].Children[0].Children[0].Title != "Subsubsection" {
		t.Errorf("Invalid got.Children[1].Children[0].Children[0].Title")
	}
}

func TestTextToPunctors(t *testing.T) {
	got := textToPunctors("Text before \n \n\n 1. Background\n   1. Subsection A\n\t1. Subsection B\n2. Methods\n   1. Subsection A\n      1. Subsubsection")

	if len(got) != 6 {
		t.Errorf("Invalid number of punctors detected")
		return
	}

	if got[0].Level != 0 {
		t.Errorf("Invalid level of punctor 0")
	}

	if got[1].Level != 1 {
		t.Errorf("Invalid level of punctor 1")
	}

	if got[2].Level != 1 {
		t.Errorf("Invalid level of punctor 2")
		return
	}

	if got[3].Level != 0 {
		t.Errorf("Invalid level of punctor 3")
		return
	}

	if got[4].Level != 1 {
		t.Errorf("Invalid level of punctor 4")
		return
	}

	if got[5].Level != 2 {
		t.Errorf("Invalid level of punctor 5")
		return
	}
}
