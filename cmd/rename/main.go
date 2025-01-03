package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"path/filepath"
)

func main() {
	notePath, folderPath, newFilename, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("note path: %s\n", notePath)
	log.Printf("new filename: %s\n", newFilename)

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	mainNote, err := core.GetNoteByPath(notes, notePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("updating note file and linked notes")
	err = core.Rename(mainNote, newFilename)
	if err != nil {
		log.Fatal(err)
	}
}

func parseArguments() (sys.Path, sys.Path, string, error) {
	notePath := flag.String("notePath", "", "a path to note file")
	newFilename := flag.String("newFilename", "", "new note filename")
	flag.Parse()

	if *notePath == "" || *newFilename == "" {
		return "", "", "", fmt.Errorf("provide both 'notePath' and 'newFilename'")
	}

	if filepath.Ext(*notePath) != sys.MdExtension {
		return "", "", "", fmt.Errorf("specify %s note path for renaming", sys.MdExtension)
	}

	return sys.Path(*notePath), sys.Path(filepath.Dir(*notePath)),
		*newFilename,
		nil
}
