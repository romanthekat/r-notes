package common

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

const BacklinksHeader = "## Backlinks"

func IsBacklinksHeader(line string) bool {
	return strings.TrimSpace(line) == BacklinksHeader
}

func GetNoteLink(note *Note) string {
	return fmt.Sprintf("%s [[%s]]", note.Name, note.Id)
}

// FillLinks TODO make links context aware - file line at least
func FillLinks(notes []*Note) []*Note {
	notesById := make(map[string]*Note)

	//lazy two-phase approach
	for _, note := range notes {
		notesById[note.Id] = note
	}

	for _, note := range notes {
		linksIds := getWikiLinks(note.GetContent())

		for _, linkId := range linksIds {
			linkedNote := notesById[linkId]
			if linkedNote == nil {
				log.Printf("[ERROR] note '%s' has broken link to id '%s'\n", note.Id, linkId)
				continue
			}

			note.Links = append(note.Links, linkedNote)
			linkedNote.Backlinks = append(linkedNote.Backlinks, note)
		}

		sort.Slice(note.Links, func(i, j int) bool {
			return note.Links[i].Id < note.Links[j].Id
		})

		sort.Slice(note.Backlinks, func(i, j int) bool {
			return note.Backlinks[i].Id < note.Backlinks[j].Id
		})
	}

	return notes
}

//getWikiLinks extracts [[LINK]] from provided Note content
//TODO guarantee order
func getWikiLinks(content []string) []string {
	set := make(map[string]struct{})
	re := regexp.MustCompile(`\[\[(.+?)]]`) //TODO compile once for app rather than once per Path

	for _, line := range content {
		if IsBacklinksHeader(line) {
			break
		}

		for _, match := range re.FindAllStringSubmatch(line, -1) {
			link := strings.TrimSpace(match[1])
			set[link] = struct{}{}
		}
	}

	var links []string
	for link := range set {
		links = append(links, link)
	}

	return links
}
