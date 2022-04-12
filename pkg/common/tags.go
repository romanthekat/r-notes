package common

import (
	"regexp"
	"sort"
	"strings"
	"sync"
)

// FillTags adds tags information to notes
func FillTags(notes []*Note) []*Note {
	var wg sync.WaitGroup
	wg.Add(len(notes))
	for _, note := range notes {
		go fillTags(note, &wg)
	}
	wg.Wait()

	return notes
}

func fillTags(note *Note, wg *sync.WaitGroup) {
	defer wg.Done()

	tags := make(map[string]any)
	for _, tag := range getTags(note.GetContent()) {
		tags[tag] = struct{}{}
	}

	note.Tags = tags
}

//getTags extracts #TAG from provided Note content
//TODO consider unification w/ links generation
//TODO consider generating tags and links w/ one content scanning
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
