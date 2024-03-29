package core

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/md"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
	"path"
	"path/filepath"
	"strings"
)

func Rename(note *Note, newFilename string) error {
	oldPath := note.Path
	oldLink := GetNoteLink(note)

	newPath := getNewPath(oldPath, newFilename)

	err := updateNote(note, oldPath, newPath)
	if err != nil {
		return err
	}

	newLink := GetNoteLink(note)
	updateBacklinks(note.Backlinks, oldLink, newLink)

	return nil
}

func ChangeId(notes []*Note, note *Note, newId string) error {
	for _, note := range notes {
		if note.Id == newId {
			return fmt.Errorf("there is already note with provided id %s: %s", newId, GetNoteLink(note))
		}
	}

	if !zk.IsZkId(newId) {
		return fmt.Errorf("new id is not a correct zk id: %s", newId)
	}

	return Rename(note, getNewFilenameById(note, newId))
}

func getNewFilenameById(note *Note, newId string) string {
	return newId + " " + note.Name + filepath.Ext(string(note.Path))
}

func getNewPath(oldPath sys.Path, newFilename string) sys.Path {
	return sys.Path(path.Join(filepath.Dir(string(oldPath)), newFilename))
}

func updateNote(note *Note, oldPath, newPath sys.Path) error {
	err := updateNoteContent(note, newPath)
	if err != nil {
		return err
	}

	err = sys.RenameFile(oldPath, newPath)
	if err != nil {
		return err
	}

	sys.WriteToFile(newPath, note.GetContent())

	return nil
}

func updateNoteContent(note *Note, newPath sys.Path) error {
	isZettel, id, name := zk.ParseNoteFilename(sys.GetFilename(newPath))
	if !isZettel {
		return fmt.Errorf("new name '%s' is not a correct zettel", name)
	}

	note.Id = id
	note.Name = name
	note.Path = newPath

	syncNoteHeader(note)

	return nil
}

func syncNoteHeader(note *Note) string {
	for i, line := range note.GetContent() {
		if md.IsFirstLevelHeader(line) {
			note.Content[i] = "# " + note.Id + " " + note.Name
			return note.Content[i]
		}
	}

	return ""
}

func updateBacklinks(backlinks []*Note, oldLink, newLink string) {
	for _, backlink := range backlinks {
		content := updateLinks(backlink, oldLink, newLink)

		sys.WriteToFile(backlink.Path, content)
	}
}

func updateLinks(note *Note, oldLink string, newLink string) []string {
	for i, line := range note.Content {
		if strings.Contains(line, oldLink) {
			note.Content[i] = strings.ReplaceAll(line, oldLink, newLink)
		}
	}

	return note.GetContent()
}
