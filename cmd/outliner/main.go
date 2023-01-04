package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/outline"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

const (
	tag = "#index"
)

func main() {
	path, folder, err := sys.GetNoteFileArgument(sys.MdExtension)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("generating outline for path", path)

	notes, err := core.GetNotesDetailed(folder, true, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("found notes files:", len(notes))

	targetNote, err := core.GetNoteById(path, notes)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("outline:\n")
	result := outline.GetNotesOutline(targetNote, "", 3, nil)
	for _, line := range result {
		fmt.Println(line)
	}

	indexTitle := fmt.Sprintf("index - %s", targetNote.Name)

	resultId, resultPath := outline.GetResultPath(path, indexTitle)
	fmt.Printf("writing to %s\n", resultPath)

	resultContent := []string{fmt.Sprintf("# %s %s", resultId, indexTitle), tag}
	resultContent = append(resultContent, result...)

	sys.WriteToFile(resultPath, resultContent)
}
