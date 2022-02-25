package common

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type Note struct {
	Id   string
	Name string
	Path Path

	//Backlinks *[]Note //TODO implement
	Links []*Note

	load    *sync.Once
	Content []string
}

func (n Note) String() string {
	return n.Name
}

func NewNote(id, name string, path Path, Links []*Note) *Note {
	return &Note{Id: id, Name: name, Path: path, Links: Links}
}

func (n *Note) GetContent() []string {
	n.load.Do(func() {
		content, err := ReadFile(n.Path)
		if err != nil {
			panic(fmt.Sprintf("can't load file %s content: %s", n, err))
		}

		n.Content = content
	})

	return n.Content
}

//GetFilesByWikiLinks parses wikilinks of a note and returns paths to correspondent files
//TODO Trie would be much better
func GetFilesByWikiLinks(currentPath Path, paths []Path, wikiLinks []string) []Path {
	var linkedFiles []Path

	for _, path := range paths {
		for _, link := range wikiLinks {
			if path != currentPath && strings.Contains(string(path), link) {
				linkedFiles = append(linkedFiles, path)
			}
		}
	}

	return linkedFiles
}

//GetWikiLinks extracts [[LINK]] from provided path content
//TODO make sure to guarantee order
func GetWikiLinks(content []string) []string {
	set := make(map[string]struct{})        //lack of golang sets ;(
	re := regexp.MustCompile(`\[\[(.+?)]]`) //TODO compile once for app rather than once per path

	for _, line := range content {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			link := strings.TrimSpace(match[1])
			set[link] = struct{}{}
		}
	}

	var links []string
	for link := range set {
		links = append(links, link)
	}

	return links
}

func GetNoteLink(note *Note) string {
	return fmt.Sprintf("%s [[%s]]", note.Name, note.Id)
}
