package main

import (
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
	"strings"
)

func main() {
	log.Println("Obtaining notes")
	note, _ := getNotes()

	log.Println("Preparing graph")
	g, graph, finishFunc := common.InitGraphviz()
	defer finishFunc()

	log.Println("Creating map note to node")
	noteToNodeMap := make(map[string]*cgraph.Node)

	notes := getNotesForSubgraph(note, 2)
	log.Println("Notes in graph: ", len(notes))

	for _, note := range notes {
		noteToNodeMap[note.Id] = common.GetNode(graph, note.Name)
	}

	node := noteToNodeMap[note.Id]
	node.SetColor("red")

	log.Println("Creating links edges")
	for _, note := range notes {
		for _, link := range note.Links {
			linkNode := noteToNodeMap[link.Id]
			if linkNode != nil {
				common.GetEdge(graph, noteToNodeMap[note.Id], linkNode, "link")
			}
		}
	}

	log.Println("Rendering to file")
	graphPath := "/tmp/graph.png"
	common.SaveGraphToFile(g, graph, graphPath)
	log.Println("file saved to", graphPath)
}

func getNotesForSubgraph(note *common.Note, levelsLeft int) []*common.Note {
	var result []*common.Note

	notesMap := getNotesForSubgraphRecursive(note, levelsLeft, make(map[*common.Note]struct{}))
	for key := range notesMap {
		result = append(result, key)
	}

	return result
}

func getNotesForSubgraphRecursive(note *common.Note, levelsLeft int, result map[*common.Note]struct{}) map[*common.Note]struct{} {
	if levelsLeft <= 0 {
		return result
	}

	if strings.HasPrefix(note.Name, "index for '") {
		return result
	}

	result[note] = struct{}{}
	var addedNotes []*common.Note

	for _, link := range note.Links {
		if _, ok := result[link]; !ok {
			result[link] = struct{}{}
			addedNotes = append(addedNotes, link)
		}
	}

	for _, link := range note.Backlinks {
		if _, ok := result[link]; !ok {
			result[link] = struct{}{}
			addedNotes = append(addedNotes, link)
		}
	}

	for _, addedNote := range addedNotes {
		result = getNotesForSubgraphRecursive(addedNote, levelsLeft-1, result)
	}

	return result
}

func getNotes() (*common.Note, []*common.Note) {
	file, folder, err := common.GetNoteFileArgument(common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	isZettel, id, _ := common.ParseNoteFilename(common.GetFilename(file))
	if !isZettel {
		log.Fatal(fmt.Errorf("provided note filename is note a correct zk note"))
	}

	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)

	var targetNote *common.Note
	for _, note := range notes {
		if note.Id == id {
			targetNote = note
			break
		}
	}

	if targetNote == nil {
		log.Fatal(fmt.Errorf("provided note wasn't found within derived notes folder"))
	}

	return targetNote, notes
}
