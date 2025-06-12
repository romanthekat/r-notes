package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"strings"
)

func main() {
	folder, level, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folder)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("notes of top-%d levels:", level)
	topNotes := core.GetTopNotes(notes, level)
	for _, note := range topNotes {
		prefix := strings.Repeat("  ", (note.Level-1)*2)
		fmt.Printf("%s%s\n", prefix, core.GetNoteLink(note))
	}
}

func parseArguments() (sys.Path, int, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	level := flag.Int("level", 1, "how many note levels to print")
	flag.Parse()

	if *notesPath == "" {
		return "", -1, fmt.Errorf("provide '-notesPath'")
	}

	return sys.Path(*notesPath), *level, nil
}
