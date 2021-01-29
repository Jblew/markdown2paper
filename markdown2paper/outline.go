package main

import (
	"regexp"
	"strings"
)

// OutlineTree is a tree of outline links
type OutlineTree struct {
	Title string
	Path string
	Children []OutlineTree
}

// ParseTextToOutlineTree parses text contents to outline tree
func ParseTextToOutlineTree(text string, level int) OutlineTree {
	punctors := textToPunctors(text)
	children := parsePunctorsToOutlineTree(punctors)
	return OutlineTree{
		Children: children,
	}
}

type punctor struct {
	Level int
	Text string
}

func parsePunctorsToOutlineTree(punctors []punctor) []OutlineTree {
	if len(punctors) == 0 {
		return []OutlineTree{}
	}

	children := []OutlineTree{}

	baseLevel := punctors[0].Level
	currentChildTitle := punctors[0].Text
	currentChildPunctors := []punctor{}
	for i, currPunctor := range punctors {
		if i == 0 {
			continue
		}
		if currPunctor.Level == baseLevel {
			children = append(children, OutlineTree{Title: currentChildTitle, Children: parsePunctorsToOutlineTree(currentChildPunctors)})
			currentChildTitle = currPunctor.Text
			currentChildPunctors = []punctor{}
		} else {
			currentChildPunctors = append(currentChildPunctors, currPunctor)
		}
	}
	children = append(children, OutlineTree{Title: currentChildTitle, Children: parsePunctorsToOutlineTree(currentChildPunctors)})

	return children
}

var punctorRe = regexp.MustCompile(`(?m)^([\t ]*)[\d+-.]+[\t ]?(.*)$`)
func textToPunctors(text string) []punctor {
	out := []punctor{}
	for _, matches := range punctorRe.FindAllStringSubmatch(text, -1) {
		if len(matches) == 3 {
			level := whiteSpacesToLevel(matches[1])
			out = append(out, punctor{Level: level, Text:matches[2]})
		}
	}
	return out
}

func whiteSpacesToLevel(whitespaceText string) int {
	hardTabsCount := strings.Count(whitespaceText, "\t")
	softTabsCount := strings.Count(whitespaceText, "   ")
	return hardTabsCount + softTabsCount
}
