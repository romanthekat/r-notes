package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
	"log"
	"path/filepath"
	"time"
)

const (
	notesDelimiter    = "    "
	markdownLineBreak = "  "
	tag               = "#index"
)

func main() {
	path, folder, err := sys.GetNoteFileArgument(sys.MdExtension)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("generating outline for path", path)

	notesPaths, err := sys.GetNotesPaths(folder, sys.MdExtension)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("found .md files:", len(notesPaths))
	log.Println("parsing links")

	notes := core.NewNotesByPaths(notesPaths)
	core.FillLinks(notes)

	targetNote, err := getTargetNote(path, notes)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("outline:\n")
	outline := getNotesOutline(targetNote, "", 3, nil)
	for _, line := range outline {
		fmt.Println(line)
	}

	indexTitle := fmt.Sprintf("index - %s", targetNote.Name)

	resultId, resultPath := getResultPath(path, indexTitle)
	fmt.Printf("writing to %s\n", resultPath)

	resultContent := []string{fmt.Sprintf("# %s %s", resultId, indexTitle), tag}
	resultContent = append(resultContent, outline...)

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

func getResultPath(path sys.Path, title string) (id string, resultPath sys.Path) {
	basePath := filepath.Dir(string(path))
	zkId := time.Now().Format("200601021504")
	return zkId, sys.Path(
		fmt.Sprintf("%s/%s %s.md", basePath, zkId, title))
}

func getNotesOutline(note *core.Note, padding string, levelsLeft int, result []string) []string {
	if levelsLeft == 0 {
		return result
	}

	if note == nil {
		return result
	}

	noteLink := core.GetNoteLink(note)
	result = append(result, fmt.Sprintf("%s- %s%s", padding, noteLink, markdownLineBreak))
	for _, child := range note.Links {
		result = getNotesOutline(child, padding+notesDelimiter, levelsLeft-1, result)
	}

	return result
}
