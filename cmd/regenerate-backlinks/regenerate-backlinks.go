package main

import (
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

func main() {
	folder, err := sys.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	notes, err := core.GetNotesDetailed(folder, true, false)
	if err != nil {
		log.Fatal(err)
	}

	core.SaveBacklinksInFiles(notes)
}
