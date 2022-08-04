package core

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
	"log"
	"sync"
)

type Note struct {
	Id   string
	Name string

	Content []string

	Links     []*Note
	Backlinks []*Note

	Tags map[string]any

	Path        sys.Path
	loadContent *sync.Once
}

func (n *Note) String() string {
	return GetNoteLink(n)
}

func NewNote(id, name string, path sys.Path, content []string) *Note {
	return &Note{Id: id, Name: name, Path: path, Content: content, loadContent: &sync.Once{}}
}

func NewNoteWithLinks(id, name string, path sys.Path, content []string, links []*Note, backlinks []*Note) *Note {
	note := NewNote(id, name, path, content)

	note.Links = links
	note.Backlinks = backlinks

	return note
}

func NewNoteWithTags(id, name string, path sys.Path, content []string, tags map[string]any) *Note {
	note := NewNote(id, name, path, content)

	note.Tags = tags

	return note
}

func (n *Note) HasId() bool {
	return n.Id != ""
}

func (n *Note) GetContent() []string {
	n.loadContent.Do(func() {
		if len(n.Content) > 0 {
			return
		}

		content, err := sys.ReadFile(n.Path)
		if err != nil {
			panic(fmt.Sprintf("can't loadContent file %s content: %s", n, err))
		}

		n.Content = content
	})

	return n.Content
}

func GetNotes(folder sys.Path) ([]*Note, error) {
	return GetNotesDetailed(folder, true, true)
}

func GetNotesDetailed(folder sys.Path, fillLinks, fillTabs bool) ([]*Note, error) {
	paths, err := sys.GetNotesPaths(folder, sys.MdExtension)
	if err != nil {
		return nil, err
	}

	notes := NewNotesByPaths(paths)
	if fillLinks {
		notes = FillLinks(notes)
	}
	if fillTabs {
		notes = FillTags(notes)
	}

	return notes, nil
}

func NewNotesByPaths(paths []sys.Path) []*Note {
	var notes []*Note

	for _, path := range paths {
		note := NewNoteByPath(path)
		if !note.HasId() {
			log.Printf("[INFO] note with Path '%s' has no id - skipping\n", note.Path)
			continue
		}

		notes = append(notes, note)
	}

	return notes
}

func NewNoteByPath(path sys.Path) *Note {
	isZettel, id, name := zk.ParseNoteFilename(sys.GetFilename(path))
	note := NewNote(id, name, path, []string{})

	if isZettel && len(name) != 0 {
		return note
	}

	if !isZettel {
		return note
	}

	name, err := zk.GetNoteNameByNoteContent(note.GetContent())
	if err != nil {
		log.Fatalf("[ERROR] error during parsing note in path '%s': %s", path, err)
	}

	note.Name = name
	return note
}
