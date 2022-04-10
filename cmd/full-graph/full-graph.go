package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
	"path/filepath"
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
	g, graph, finishFunc := common.InitGraphviz()
	defer finishFunc()

	noteToNodeMap := make(map[string]*cgraph.Node)
	for _, note := range notes {
		noteToNodeMap[note.Id] = common.GetNode(graph, note.Name)
	}

	for _, note := range notes {
		for _, link := range note.Links {
			edge := common.GetEdge(graph, noteToNodeMap[note.Id], noteToNodeMap[link.Id], "link")
			edge.SetLabel("link")
		}
	}

	log.Println("rendering to file")
	common.SaveGraphToFile(g, graph, string(outputPath))
	log.Println("file saved to", outputPath)
}

func getNotes(folder common.Path) []*common.Note {
	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)
	common.FillTags(notes)

	return notes
}

func parseArguments() (common.Path, common.Path, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to rendered graph file")
	flag.Parse()

	if filepath.Ext(*notesPath) != common.MdExtension {
		return "", "", fmt.Errorf("specify %s path for generating graph", common.MdExtension)
	}

	if *notesPath == "" || *outputPath == "" {
		return "", "", fmt.Errorf("provide both 'notesPath' and 'outputPath'")
	}

	return common.Path(*notesPath), common.Path(*outputPath), nil
}
