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

	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)
	for _, note := range notes {
		fmt.Printf("%s [[%s]]\n", note.Name, note.Id)

		if len(note.Links) > 0 {
			fmt.Printf("\tlinks: %s\n", note.Links)
		}
		if len(note.Backlinks) > 0 {
			fmt.Printf("\tbacklinks: %s\n", note.Backlinks)
		}
		fmt.Println("")
	}

}
