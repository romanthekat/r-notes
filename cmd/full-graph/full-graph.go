package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/render"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
)

//TODO seems to be too big on 700+ notes
func main() {
	folder, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes, err := core.GetNotes(folder)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("preparing graph")
	g, graph, finishFunc := render.InitGraphviz()
	defer finishFunc()

	noteToNodeMap := make(map[string]*cgraph.Node)
	for _, note := range notes {
		noteToNodeMap[note.Id] = render.GetNode(graph, note.Name)
	}

	for _, note := range notes {
		for _, link := range note.Links {
			edge := render.GetEdge(graph, noteToNodeMap[note.Id], noteToNodeMap[link.Id], "link")
			edge.SetLabel("link")
		}
	}

	log.Println("rendering to file")
	render.SaveGraphToFile(g, graph, string(outputPath))
	log.Println("file saved to", outputPath)
}

func parseArguments() (sys.Path, sys.Path, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to rendered graph file")
	flag.Parse()

	if *notesPath == "" || *outputPath == "" {
		return "", "", fmt.Errorf("provide both 'notesPath' and 'outputPath'")
	}

	return sys.Path(*notesPath), sys.Path(*outputPath), nil
}
