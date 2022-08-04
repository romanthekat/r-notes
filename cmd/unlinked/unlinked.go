package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

func main() {
	folder, err := sys.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	notes, err := core.GetNotesPartial(folder, true, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("notes without (back)links:")

	for _, note := range notes {
		if len(note.Links) == 0 && len(note.Backlinks) == 0 {
			fmt.Println(core.GetNoteLink(note))
		}
	}
}
