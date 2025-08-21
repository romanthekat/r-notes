package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
)

func main() {
	folder, outputPath, filterSubstring, filterMainNote, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folder)
	if err != nil {
		log.Fatal(err)
	}

	var mainNote *core.Note
	if filterMainNote != "" {
		log.Println("filtering by relevancy to main note")
		isZettel, id, _ := zk.ParseNoteFilename(filterMainNote)
		if !isZettel {
			log.Fatal(fmt.Errorf("provided note name is not a correct zk note"))
		}

		for _, note := range notes {
			if note.Id == id {
				mainNote = note
				break
			}
		}
		if mainNote == nil {
			log.Fatal(fmt.Errorf("note not found for input %s", filterMainNote))
		}

		notes = core.GetRelevantNotes(mainNote, notes)
	}

	if len(filterSubstring) > 0 {
		log.Println("filtering substring")
		notes = core.FilterNotesBySubstring(notes, filterSubstring)
	}

	notes = core.SortByRank(notes)

	result := core.JoinContent(notes)

	if outputPath != "" {
		sys.WriteToFile(outputPath, result)
		log.Println("file saved to", outputPath)
	} else {
		for _, line := range result {
			fmt.Println(line)
		}
	}
}

func parseArguments() (sys.Path, sys.Path, string, string, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "", "a path to result join file")
	filterSubstring := flag.String("filterSubstring", "", "a substring in id or name to filter by")
	filterMainNote := flag.String("filterMainNote", "", "full note to filter by relevance to")
	flag.Parse()

	if *notesPath == "" {
		return "", "", "", "", fmt.Errorf("provide '-notesPath', " +
			"and optional: '-outputPath', '-filterSubstring', '-filterMainNote'")
	}

	return sys.Path(*notesPath), sys.Path(*outputPath), *filterSubstring, *filterMainNote, nil
}
