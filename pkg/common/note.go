package common

import (
	"fmt"
	"log"
	"sync"
)

type Note struct {
	Id   string
	Name string

	Content []string

	Links     []*Note
	Backlinks []*Note

	Path        Path
	loadContent *sync.Once
}

func (n *Note) String() string {
	return GetNoteLink(n)
}

func NewNote(id, name string, path Path, content []string) *Note {
	return &Note{Id: id, Name: name, Path: path, Content: content, loadContent: &sync.Once{}}
}

func NewNoteWithLinks(id, name string, path Path, content []string, links []*Note, backlinks []*Note) *Note {
	note := NewNote(id, name, path, content)

	note.Links = links
	note.Backlinks = backlinks

	return note
}

func (n *Note) HasId() bool {
	return n.Id != ""
}

func (n *Note) GetContent() []string {
	n.loadContent.Do(func() {
		if n.Content != nil {
			return
		}

		content, err := ReadFile(n.Path)
		if err != nil {
			panic(fmt.Sprintf("can't loadContent file %s content: %s", n, err))
		}

		n.Content = content
	})

	return n.Content
}

func NewNotesByPaths(paths []Path) []*Note {
	var notes []*Note

	for _, path := range paths {
		note := NewNoteByPath(path)
		if !note.HasId() {
			log.Printf("[ERROR] note with Path '%s' has no id - skipping it\n", note.Path)
			continue
		}

		notes = append(notes, note)
	}

	return notes
}

func NewNoteByPath(path Path) *Note {
	isZettel, id, name := ParseNoteFilename(GetFilename(path))
	note := NewNote(id, name, path, []string{})

	if isZettel && len(name) != 0 {
		return note
	}

	name, err := GetNoteNameByNoteContent(note.GetContent())
	if err != nil {
		panic(err)
	}

	note.Name = name
	return note
}
