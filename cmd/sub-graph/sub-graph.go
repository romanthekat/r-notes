package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	notePath, folderPath, graphDepth, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("note path: %s\n", notePath)
	log.Printf("graph output path: %s\n", outputPath)
	log.Printf("graph depth: %d\n", graphDepth)

	log.Println("obtaining notes")
	note, _ := getNotes(notePath, folderPath)

	log.Println("preparing graph")
	g, graph, finishFunc := common.InitGraphviz()
	defer finishFunc()

	noteToNodeMap := make(map[string]*cgraph.Node)

	log.Println("getting notes for subgraph")
	notes := getNotesForSubgraph(note, graphDepth)
	log.Println("notes in graph:", len(notes))

	for _, note := range notes {
		noteToNodeMap[note.Id] = common.GetNode(graph, note.Name)
	}

	node := noteToNodeMap[note.Id]
	node.SetColor("red")

	for _, note := range notes {
		for _, link := range note.Links {
			linkNode := noteToNodeMap[link.Id]
			if linkNode != nil {
				common.GetEdge(graph, noteToNodeMap[note.Id], linkNode, "link")
			}
		}
	}

	log.Println("rendering to file")
	common.SaveGraphToFile(g, graph, string(outputPath))
	log.Println("graph saved to:", outputPath)
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

func getNotes(notePath, folderPath common.Path) (*common.Note, []*common.Note) {
	isZettel, id, _ := common.ParseNoteFilename(common.GetFilename(notePath))
	if !isZettel {
		log.Fatal(fmt.Errorf("provided note filename is not a correct zk note"))
	}

	paths, err := common.GetNotesPaths(folderPath, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)
	common.FillTags(notes)

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

func parseArguments() (common.Path, common.Path, int, common.Path, error) {
	notePath := flag.String("notePath", "", "a path to note file")
	outputPath := flag.String("outputPath", "./", "a path to rendered graph file")
	graphDepth := flag.Int("depth", 2, "graph depth to render")
	flag.Parse()

	if filepath.Ext(*notePath) != common.MdExtension {
		return "", "", -1, "", fmt.Errorf("specify %s path for generating graph", common.MdExtension)
	}

	if *notePath == "" || *outputPath == "" {
		return "", "", -1, "", fmt.Errorf("provide both 'notePath' and 'outputPath'")
	}

	return common.Path(*notePath), common.Path(filepath.Dir(*notePath)), *graphDepth, common.Path(*outputPath), nil
}
