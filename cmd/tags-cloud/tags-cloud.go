package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
	"sort"
)

func main() {
	folder, err := common.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	notes := getNotes(folder)
	common.FillTags(notes)

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
		fmt.Printf("%s\t%d\n", tag, tagsMap[tag])
	}
}

func getNotes(folder common.Path) []*common.Note {
	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)

	return notes
}
