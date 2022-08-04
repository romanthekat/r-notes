package main

import (
	"github.com/romanthekat/r-notes/pkg/core"
	"log"
)

func main() {
	folder, err := core.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	paths, err := core.GetNotesPaths(folder, core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := core.NewNotesByPaths(paths)
	core.FillLinks(notes)

	core.SaveBacklinksInFiles(notes)
}
