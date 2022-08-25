package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/outline"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
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

	targetNote, err := getTargetNote(path, notes)
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

func getTargetNote(path sys.Path, notes []*core.Note) (*core.Note, error) {
	_, id, _ := zk.ParseNoteFilename(sys.GetFilename(path))

	var targetNote *core.Note
	for _, note := range notes {
		if note.Id == id {
			targetNote = note
			break
		}
	}

	if targetNote == nil {
		return nil, errors.New("provided note path wasn't correctly parsed as a zk note")
	}

	return targetNote, nil
}
