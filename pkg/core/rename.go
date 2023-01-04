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

    err := updateNote(note, newPath, newName)
    if err != nil {
        return err
    }

	err = sys.RenameFile(oldPath, newPath)
	if err != nil {
		return err
	}

	newLink := GetNoteLink(note)
	updateBacklinks(note.Backlinks, newPath, oldLink, newLink)

	return nil
}

func getNewPath(oldPath sys.Path, newName string) sys.Path {
    return sys.Path(path.Join(filepath.Dir(string(oldPath)), newName))
}

func updateNote(note *Note, newPath sys.Path, newName string) error {
    isZettel, id, name := zk.ParseNoteFilename(string(newPath))
    if !isZettel {
        return fmt.Errorf("new name '%s' is not a correct zettel", newName)
    }

    note.Id = id
    note.Name = name

    for i, line := range note.GetContent() {
        if md.IsFirstLevelHeader(line) {
            note.Content[i] = strings.ReplaceAll(line, note.Name, newName)
        }
    }

    return nil
}

func updateBacklinks(backlinks []*Note, newPath sys.Path, oldLink, newLink string) {
	for _, backlink := range backlinks {
		for i, line := range backlink.Content {
			if strings.Contains(line, oldLink) {
				backlink.Content[i] = strings.ReplaceAll(line, oldLink, newLink)
			}
		}

		sys.WriteToFile(newPath, backlink.GetContent())
	}
}
