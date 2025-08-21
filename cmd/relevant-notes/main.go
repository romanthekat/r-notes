package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
)

func main() {
	notePath, folderPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("mainNote path: %s\n", notePath)

	log.Println("obtaining notes")
	mainNote, notes := getNotes(notePath, folderPath)

	log.Println("getting relevant notes")
	relevantNotes := core.GetRelevantNotes(mainNote, notes)

	for _, note := range relevantNotes {
		fmt.Println(core.GetNoteLink(note))
	}
}

func getNotes(notePath, folderPath sys.Path) (*core.Note, []*core.Note) {
	isZettel, id, _ := zk.ParseNoteFilename(sys.GetFilename(notePath))
	if !isZettel {
		log.Fatal(fmt.Errorf("provided note filename is not a correct zk note"))
	}

	notes, err := core.GetNotes(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	var targetNote *core.Note
	for _, note := range notes {
		if note.Id == id {
			targetNote = note
			break
		}
	}

	if targetNote == nil {
		log.Fatal(fmt.Errorf("provided note wasn't found within derived notes folder"))
	}

	return targetNote, notes
}

func parseArguments() (sys.Path, sys.Path, error) {
	notePath := flag.String("notePath", "", "a path to note file")
	flag.Parse()

	if *notePath == "" {
		return "", "", fmt.Errorf("provide '-notePath'")
	}

	if filepath.Ext(*notePath) != sys.MdExtension {
		return "", "", fmt.Errorf("specify %s note path for generating sub-graph", sys.MdExtension)
	}

	return sys.Path(*notePath), sys.Path(filepath.Dir(*notePath)), nil
}
