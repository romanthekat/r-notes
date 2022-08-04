package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"log"
)

func main() {
	folder, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes := getNotes(folder)

	notes = core.SortByRank(notes)
	//for _, note := range notes {
	//	fmt.Println(note.String())
	//}

	result := core.JoinContent(notes)
	core.WriteToFile(outputPath, result)

	log.Println("file saved to", outputPath)
}

func getNotes(folder core.Path) []*core.Note {
	paths, err := core.GetNotesPaths(folder, core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := core.NewNotesByPaths(paths)
	core.FillLinks(notes)
	core.FillTags(notes)

	return notes
}

func parseArguments() (core.Path, core.Path, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to result join file")
	flag.Parse()

	if *notesPath == "" || *outputPath == "" {
		return "", "", fmt.Errorf("provide both 'notesPath' and 'outputPath'")
	}

	return core.Path(*notesPath), core.Path(*outputPath), nil
}
