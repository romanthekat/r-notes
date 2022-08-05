package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

func main() {
	folder, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folder)
	if err != nil {
		log.Fatal(err)
	}

	notes = core.SortByRank(notes)

	result := core.JoinContent(notes)
	sys.WriteToFile(outputPath, result)

	log.Println("file saved to", outputPath)
}

func parseArguments() (sys.Path, sys.Path, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to result join file")
	flag.Parse()

	if *notesPath == "" || *outputPath == "" {
		return "", "", fmt.Errorf("provide both 'notesPath' and 'outputPath'")
	}

	return sys.Path(*notesPath), sys.Path(*outputPath), nil
}
