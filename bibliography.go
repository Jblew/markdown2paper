package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Jblew/bibtex"
)

// Bibliography is the bibliography of an article
type Bibliography struct {
	bibTex *bibtex.BibTex
}

// FormatMarkdownByKey formats bibliography entry by key
func (b *Bibliography) FormatMarkdownByKey(key string) string {
	entry := b.findEntryByKey(key)
	if entry == nil {
		log.Printf("Warning: citation for key '%s' not found in bibtex file", key)
		return fmt.Sprintf("[@%s] Citation not found in bibtex file", key)
	}
	return formatBibtexEntry(entry)
}

func formatBibtexEntry(entry *bibtex.BibEntry) string {
	author := "anonymus";
	if authorField, ok := entry.Fields["author"]; ok {
		author = strings.ReplaceAll(strings.ReplaceAll(authorField.String(), "{", ""), "}", "")
	}
	year := ""
	if yearField, ok := entry.Fields["year"]; ok {
		year = yearField.String()
	}
	title := fmt.Sprintf("*%s*", entry.Fields["title"])
	titleNoBrackets := strings.ReplaceAll(strings.ReplaceAll(title, "{", ""), "}", "")
	journal := "";
	if journalField, ok := entry.Fields["journal"]; ok {
		journal = journalField.String()
	}
	doiFormatted := ""
	if doi, ok := entry.Fields["doi"]; ok {
		doiFormatted = fmt.Sprintf("https://dx.doi.org/%s", doi)
	}
	elems := []string{author, titleNoBrackets, journal, year, doiFormatted}
	return strings.Join(StringSliceRemoveEmpty(elems), ", ")
}

func (b *Bibliography) findEntryByKey(key string) *bibtex.BibEntry {
	keyLowerCase := strings.ToLower(key)
	for _, entry := range b.bibTex.Entries {
		if strings.ToLower(entry.CiteName) == keyLowerCase {
			return entry
		}
	}
	return nil
}

// LoadBibliographyFromPath loads bibliography from file
func LoadBibliographyFromPath(path string) (Bibliography, error) {
	bibliographyBytes, err := readBibtexFromPath(path)
	if err != nil {
		return Bibliography{}, err
	}

	if verbose {
		log.Printf("Loaded bibtex contents: %v", string(bibliographyBytes))
	}

	bibTex, err := bibtex.Parse(bytes.NewReader(bibliographyBytes))
	if err != nil {
		return Bibliography{}, err
	}
	if verbose {
		log.Printf("Loaded bibtex: %+v", bibTex.Entries)
	}

	return Bibliography{ bibTex }, nil
}

func readBibtexFromPath(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http") {
		return fetchBibtexFromHttp(path)
	}
	return ioutil.ReadFile(path)
}

func fetchBibtexFromHttp(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		if verbose {
			log.Printf("Failed to fetch bibtex from %s", path)
		}
		log.Fatalln(err)
	}
	if verbose {
		log.Printf("Fetched bibtex from %s", path)
	}
	return ioutil.ReadAll(resp.Body)
}
