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

func Rename(note *Note, newName string) error {
	oldPath := note.Path
	oldLink := GetNoteLink(note)

	newPath := getNewPath(oldPath, newName)

	err := updateNote(note, oldPath, newPath, newName)
	if err != nil {
		return err
	}

	newLink := GetNoteLink(note)
	updateBacklinks(note.Backlinks, oldLink, newLink)

	return nil
}

func getNewPath(oldPath sys.Path, newName string) sys.Path {
	return sys.Path(path.Join(filepath.Dir(string(oldPath)), newName))
}

func updateNote(note *Note, oldPath, newPath sys.Path, newName string) error {
	isZettel, id, name := zk.ParseNoteFilename(sys.GetFilename(newPath))
	if !isZettel {
		return fmt.Errorf("new name '%s' is not a correct zettel", newName)
	}

	note.Id = id
	note.Name = name
	note.Path = newPath

	for i, line := range note.GetContent() {
		if md.IsFirstLevelHeader(line) {
			note.Content[i] = "# " + note.Id + note.Name
		}
	}

	err := sys.RenameFile(oldPath, newPath)
	if err != nil {
		return err
	}

	sys.WriteToFile(newPath, note.GetContent())

	return nil
}

func updateBacklinks(backlinks []*Note, oldLink, newLink string) {
	for _, backlink := range backlinks {
		for i, line := range backlink.Content {
			if strings.Contains(line, oldLink) {
				backlink.Content[i] = strings.ReplaceAll(line, oldLink, newLink)
			}
		}

		sys.WriteToFile(backlink.Path, backlink.GetContent())
	}
}
