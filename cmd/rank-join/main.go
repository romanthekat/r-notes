package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

func main() {
	folder, outputPath, filterSubstring, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folder)
	if err != nil {
		log.Fatal(err)
	}

	if len(filterSubstring) > 0 {
		log.Println("filtering substring")
		notes = core.FilterNotesBySubstring(notes, filterSubstring)
	}

	notes = core.SortByRank(notes)

	result := core.JoinContent(notes)
	sys.WriteToFile(outputPath, result)

	log.Println("file saved to", outputPath)
}

func parseArguments() (sys.Path, sys.Path, string, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to result join file")
	filterSubstring := flag.String("filterSubstring", "", "a substring in id or name to filter by")
	flag.Parse()

	if *notesPath == "" || *outputPath == "" {
		return "", "", "", fmt.Errorf("provide '-notesPath' and '-outputPath', and optional '-filterSubstring'")
	}

	return sys.Path(*notesPath), sys.Path(*outputPath), *filterSubstring, nil
}
