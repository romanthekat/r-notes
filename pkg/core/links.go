package core

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/md"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
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
	return fmt.Sprintf("[[%s %s]]", note.Id, note.Name)
}

// FillLinks TODO make links context aware - file line at least
// TODO consider creating benchmarks
func FillLinks(notes []*Note) []*Note {
	notesById := make(map[string]*Note)

	//lazy two-phase approach
	for _, note := range notes {
		notesById[note.Id] = note
	}

	for _, note := range notes {
		linksIds := getWikilinks(note.GetContent())

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

// SaveBacklinksInFiles
// TODO can be done concurrently, if all necessary info is already in memory
func SaveBacklinksInFiles(notes []*Note) {
	for _, note := range notes {
		content, err := generateContentWithBacklinks(note)
		if err != nil {
			log.Printf("[ERROR] %s: skip file due to error: %s\n", GetNoteLink(note), err)
			continue
		}

		if !IsSameContent(content, note.GetContent()) {
			sys.WriteToFile(note.Path, content)
		}
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
	if len(note.Backlinks) == 0 {
		return nil
	}

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

		if found && md.IsMarkdownHeader(line) && line != "## ..." {
			return found, index, fmt.Errorf("there is a markdown header after backlinks header - structure is incorrect")
		}
	}

	return found, index, nil

}

// getWikilinks extracts [[LINK_ID]] from provided Note content
func getWikilinks(content []string) []string {
	set := make(map[string]struct{})
	re := regexp.MustCompile(`\[\[([A-Za-z\p{Cyrillic}\d]+?[A-Za-z\p{Cyrillic}\d:\-Σ/.,_ ]+?)]]`) //TODO compile once for app

	for _, line := range content {
		if IsBacklinksHeader(line) {
			break
		}

		for _, match := range re.FindAllStringSubmatch(line, -1) {
			link := strings.TrimSpace(match[1])

			isFullFormat, id := isLinkFullFormat(link)
			if isFullFormat {
				set[id] = struct{}{}
				continue
			}

			set[link] = struct{}{}
		}
	}

	var links []string
	for link := range set {
		links = append(links, link)
	}

	sort.Strings(links)

	return links
}

func isLinkFullFormat(link string) (bool, string) {
	if strings.Contains(link, " ") {
		parts := strings.Split(link, " ")
		if len(parts) > 0 && zk.IsZkId(parts[0]) {
			return true, parts[0]
		}
	}

	return false, ""
}
