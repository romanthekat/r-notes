package common

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

const BacklinksHeader = "## Backlinks"

//TODO more reliable parsing would be beneficial
func IsMarkdownHeader(line string) bool {
	return strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### ")
}

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
				log.Printf("[ERROR] note '%s' has broken link to id '%s'\n", GetNoteLink(note), linkId)
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

func SaveBacklinksInFiles(notes []*Note) {
	for _, note := range notes {
		content, err := generateContentWithBacklinks(note)
		if err != nil {
			log.Printf("[ERROR] %s: skip file due to error: %s\n", GetNoteLink(note), err)
			continue
		}

		WriteToFile(note.Path, content)
	}
}

func generateContentWithBacklinks(note *Note) ([]string, error) {
	backlinksContent := generateBacklinksContent(note)

	backlinksHeaderExists, backlinksHeaderIdx, err := findBacklinkHeader(note)
	if err != nil {
		return nil, err
	}

	content := append([]string{}, note.GetContent()...)

	copyUntilIdx := len(content)
	if backlinksHeaderExists {
		copyUntilIdx = backlinksHeaderIdx
	}

	content = append(content[:copyUntilIdx], backlinksContent...)
	return content, nil
}

func generateBacklinksContent(note *Note) []string {
	backlinksContent := []string{BacklinksHeader}

	for _, backlink := range note.Backlinks {
		backlinksContent = append(backlinksContent, "- "+GetNoteLink(backlink))
	}

	return backlinksContent
}

func findBacklinkHeader(note *Note) (found bool, index int, err error) {
	for i, line := range note.GetContent() {
		if IsBacklinksHeader(line) {
			found = true
			index = i
			continue
		}

		if found && IsMarkdownHeader(line) && line != "## ..." {
			return found, index, fmt.Errorf("there is a markdown header after backlinks header - structure is incorrect")
		}
	}

	return found, index, nil

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
