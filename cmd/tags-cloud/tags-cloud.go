package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"sort"
)

func main() {
	folder, err := sys.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	notes := getNotes(folder)
	core.FillTags(notes)

	tagsMap := make(map[string]int)
	for _, note := range notes {
		for tag := range note.Tags {
			if len(tag) > 0 {
				tagsMap[tag] += 1
			}
		}
	}

	tags := make([]string, 0, len(tagsMap))
	for tag := range tagsMap {
		tags = append(tags, tag)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tagsMap[tags[i]] > tagsMap[tags[j]]
	})

	for _, tag := range tags {
		fmt.Printf("%s x%d\n", tag, tagsMap[tag])
	}
}

func getNotes(folder sys.Path) []*core.Note {
	paths, err := sys.GetNotesPaths(folder, sys.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := core.NewNotesByPaths(paths)
	core.FillLinks(notes)

	return notes
}
