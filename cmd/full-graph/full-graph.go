package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/core"
	"log"
)

//TODO seems to be too big on 700+ notes
func main() {
	folder, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes := getNotes(folder)

	log.Println("preparing graph")
	g, graph, finishFunc := core.InitGraphviz()
	defer finishFunc()

	noteToNodeMap := make(map[string]*cgraph.Node)
	for _, note := range notes {
		noteToNodeMap[note.Id] = core.GetNode(graph, note.Name)
	}

	for _, note := range notes {
		for _, link := range note.Links {
			edge := core.GetEdge(graph, noteToNodeMap[note.Id], noteToNodeMap[link.Id], "link")
			edge.SetLabel("link")
		}
	}

	log.Println("rendering to file")
	core.SaveGraphToFile(g, graph, string(outputPath))
	log.Println("file saved to", outputPath)
}

func getNotes(folder core.Path) []*core.Note {
	paths, err := core.GetNotesPaths(folder, core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := core.NewNotesByPaths(paths)
	core.FillLinks(notes)
	core.FillTags(notes)

	return notes
}

func parseArguments() (core.Path, core.Path, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to rendered graph file")
	flag.Parse()

	if *notesPath == "" || *outputPath == "" {
		return "", "", fmt.Errorf("provide both 'notesPath' and 'outputPath'")
	}

	return core.Path(*notesPath), core.Path(*outputPath), nil
}
