package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
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
	path, folder, err := core.GetNoteFileArgument(core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("generating outline for path", path)
	_, id, _ := core.ParseNoteFilename(core.GetFilename(path))

	notesPaths, err := core.GetNotesPaths(folder, core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("found .md files:", len(notesPaths))
	log.Println("parsing links")

	notes := core.NewNotesByPaths(notesPaths)
	core.FillLinks(notes)

	var targetNote *core.Note
	for _, note := range notes {
		if note.Id == id {
			targetNote = note
			break
		}
	}
	if targetNote == nil {
		log.Fatal("provided note path wasn't correctly parsed as a zk note")
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

	core.WriteToFile(resultPath, resultContent)
}

func getResultPath(path core.Path, title string) (id string, resultPath core.Path) {
	basePath := filepath.Dir(string(path))
	zkId := time.Now().Format("200601021504")
	return zkId, core.Path(
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
