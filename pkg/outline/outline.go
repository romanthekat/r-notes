package outline

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/md"
	"github.com/romanthekat/r-notes/pkg/sys"
	"path/filepath"
	"time"
)

const notesDelimiter = "    "

func GetResultPath(path sys.Path, title string) (id string, resultPath sys.Path) {
	basePath := filepath.Dir(string(path))
	zkId := time.Now().Format("200601021504")
	return zkId, sys.Path(
		fmt.Sprintf("%s/%s %s.md", basePath, zkId, title))
}

func GetNotesOutline(note *core.Note, padding string, levelsLeft int, result []string) []string {
	if levelsLeft == 0 {
		return result
	}

	if note == nil {
		return result
	}

	noteLink := core.GetNoteLink(note)
	result = append(result, fmt.Sprintf("%s- %s%s", padding, noteLink, md.MarkdownLineBreak))
	for _, child := range note.Links {
		result = GetNotesOutline(child, padding+notesDelimiter, levelsLeft-1, result)
	}

	return result
}
