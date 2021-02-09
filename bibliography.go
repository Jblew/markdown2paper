package main
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nickng/bibtex"
)

// Bibliography is the bibliography of an article
type Bibliography struct {
	bibTex *bibtex.BibTex
}

// LoadBibliographyFromFile loads bibliography from file
func LoadBibliographyFromFile(path string) (Bibliography, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Bibliography{}, err
	}

	bibTex, err := bibtex.Parse(bytes.NewReader(b))
	if err != nil {
		return Bibliography{}, err
	}

	return Bibliography{ bibTex }, nil
}


// FormatMarkdownByKey formats bibliography entry by key
func (b *Bibliography) FormatMarkdownByKey(key string) string {
	entry := b.findEntryByKey(key)
	if entry != nil {
		entryType := entry.Type
		author := entry.Fields["author"]
		year := entry.Fields["year"]
		title := entry.Fields["title"]
		journal := entry.Fields["journal"]
		doi := entry.Fields["doi"]
		return fmt.Sprintf("%s [@%s] %s, %s, **%s**, *%s*, https://dx.doi.org/%s", entryType, key, author, year, title, journal, doi)
	}
	log.Printf("Warning: citation for key '%s' not found in bibtex file", key)
	return fmt.Sprintf("[@%s] Citation not found in bibtex file", key)
}

func (b *Bibliography) findEntryByKey(key string) *bibtex.BibEntry {
	for _, entry := range b.bibTex.Entries {
		if entry.CiteName == key {
			return entry
		}
	}
	return nil
}
