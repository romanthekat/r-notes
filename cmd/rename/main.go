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
	notePath, folderPath, newName, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("note path: %s\n", notePath)
	log.Printf("new name: %s\n", newName)

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	mainNote, err := core.GetNoteById(notePath, notes)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("updating note file and linked notes")
	err = core.Rename(mainNote, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func parseArguments() (sys.Path, sys.Path, string, error) {
	notePath := flag.String("notePath", "", "a path to note file")
	newName := flag.String("newName", "", "new note (file)name, w/o extension")
	flag.Parse()

	if *notePath == "" || *newName == "" {
		return "", "", "", fmt.Errorf("provide both 'notePath' and 'newName'")
	}

	if filepath.Ext(*notePath) != sys.MdExtension {
		return "", "", "", fmt.Errorf("specify %s note path for renaming", sys.MdExtension)
	}

	return sys.Path(*notePath), sys.Path(filepath.Dir(*notePath)),
		*newName,
		nil
}
