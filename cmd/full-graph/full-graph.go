package main

import (
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
)

//TODO seems to be too big on 700+ notes
func main() {
	log.Println("Obtaining notes")
	notes := getNotes()

	log.Println("Preparing graph")
	g, graph, finishFunc := common.InitGraphviz()
	defer finishFunc()

	log.Println("Creating map note to node")
	noteToNodeMap := make(map[string]*cgraph.Node)
	for _, note := range notes {
		noteToNodeMap[note.Id] = common.GetNode(graph, note.Name)
	}

	log.Println("Creating links edges")
	for _, note := range notes {
		for _, link := range note.Links {
			edge := common.GetEdge(graph, noteToNodeMap[note.Id], noteToNodeMap[link.Id], "link")
			edge.SetLabel("link")
		}
	}

	log.Println("Rendering to file")
	graphPath := "/tmp/graph.png"
	common.SaveGraphToFile(g, graph, graphPath)
	log.Println("file saved to", graphPath)
}

func getNotes() []*common.Note {
	folder, err := common.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)

	return notes
}
