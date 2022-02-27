package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
)

func main() {
	folder, err := common.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	notes := getNotes(folder)
	log.Println("notes without (back)links:")

	for _, note := range notes {
		if len(note.Links) == 0 && len(note.Backlinks) == 0 {
			fmt.Println(common.GetNoteLink(note))
		}
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
