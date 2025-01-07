package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

func main() {
	notesPath, oldId, newId, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("notes path: %s\n", notesPath)
	log.Printf("oldId: %s\n", oldId)
	log.Printf("newId: %s\n", newId)

	log.Println("obtaining notes")
	notes, err := core.GetNotes(notesPath)
	if err != nil {
		log.Fatal(err)
	}

	note, err := core.GetNoteById(notes, oldId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("updating note file and linked notes")
	err = core.ChangeId(notes, note, newId)
	if err != nil {
		log.Fatal(err)
	}
}

func parseArguments() (sys.Path, string, string, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	oldId := flag.String("oldId", "", "old note id")
	newId := flag.String("newId", "", "new note id")
	flag.Parse()

	if *notesPath == "" || *oldId == "" || *newId == "" {
		return "", "", "", fmt.Errorf("provide '-notesPath', '-oldId', '-newId'")
	}

	return sys.Path(*notesPath),
		*oldId,
		*newId,
		nil
}
