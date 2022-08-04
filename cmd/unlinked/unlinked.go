package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"log"
)

func main() {
	folder, err := core.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	notes := getNotes(folder)
	log.Println("notes without (back)links:")

	for _, note := range notes {
		if len(note.Links) == 0 && len(note.Backlinks) == 0 {
			fmt.Println(core.GetNoteLink(note))
		}
	}
}

func getNotes(folder core.Path) []*core.Note {
	paths, err := core.GetNotesPaths(folder, core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := core.NewNotesByPaths(paths)
	core.FillLinks(notes)

	return notes
}
