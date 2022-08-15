package core

import (
	"regexp"
	"sort"
	"strings"
)

// FillTags adds tags information to notes
// TODO consider creating benchmarks
func FillTags(notes []*Note) []*Note {
	for _, note := range notes {
		tags := make(map[string]any)
		for _, tag := range getTags(note.GetContent()) {
			tags[tag] = struct{}{}
		}

		note.Tags = tags
	}

	return notes
}

// getTags extracts #TAG from provided Note content
// TODO consider unification w/ links generation
// TODO consider generating tags and links w/ one content scanning
func getTags(content []string) []string {
	set := make(map[string]struct{})

	re := regexp.MustCompile(`\B#([A-Za-z0-9\-_]+)`) //TODO once per app

	for _, line := range content {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			link := strings.TrimSpace(match[1])
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
