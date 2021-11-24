package main
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Jblew/bibtex"
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
	if entry == nil {
		log.Printf("Warning: citation for key '%s' not found in bibtex file", key)
		return fmt.Sprintf("[@%s] Citation not found in bibtex file", key)
	}
	return formatBibtexEntry(entry)
}

func formatBibtexEntry(entry *bibtex.BibEntry) string {
	entryType := entry.Type
	author := entry.Fields["author"]
	year := entry.Fields["year"]
	title := entry.Fields["title"]
	journal := entry.Fields["journal"]
	doiFormatted := "<no doi>"
	if doi, ok := entry.Fields["doi"]; ok {
		doiFormatted = fmt.Sprintf("https://dx.doi.org/%s", doi)
	}
	key := entry.CiteName
	return fmt.Sprintf("%s [@%s] %s, %s, **%s**, *%s*, %s", entryType, key, author, year, title, journal, doiFormatted)
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
