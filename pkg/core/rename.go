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

    newPath := sys.Path(path.Join(filepath.Dir(string(oldPath)), newName))
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

    err := sys.RenameFile(oldPath, newPath)
    if err != nil {
        return err
    }

    newLink := GetNoteLink(note)
    for _, backlink := range note.Backlinks {
        for i, line := range backlink.Content {
            if strings.Contains(line, oldLink) {
                backlink.Content[i] = strings.ReplaceAll(line, oldLink, newLink)
            }
        }

        sys.WriteToFile(newPath, backlink.GetContent())
    }

    return nil
}
