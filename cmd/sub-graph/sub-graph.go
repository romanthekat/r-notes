package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/render"
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/zk"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	notePath, folderPath, graphDepth, ignoreTags, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("note path: %s\n", notePath)
	log.Printf("graph output path: %s\n", outputPath)
	log.Printf("graph depth: %d\n", graphDepth)
	log.Printf("tags to ignore: %s\n", ignoreTags)

	log.Println("obtaining notes")
	note, _ := getNotes(notePath, folderPath)

	log.Println("preparing graph")
	g, graph, finishFunc := render.InitGraphviz()
	defer finishFunc()

	noteToNodeMap := make(map[string]*cgraph.Node)

	log.Println("getting notes for subgraph")
	notes := getNotesForSubgraph(note, graphDepth, ignoreTags)
	log.Println("notes in graph:", len(notes))

	for _, note := range notes {
		noteToNodeMap[note.Id] = render.GetNode(graph, note.Name, note.Tags)
	}

	renderMainNodesGroup(noteToNodeMap, note)

	noteToNoteMap := make(map[string]map[string]*cgraph.Edge)

	for _, note := range notes {
		noteToNoteMap[note.Id] = make(map[string]*cgraph.Edge)

		for _, link := range note.Links {
			if linkNode := noteToNodeMap[link.Id]; linkNode != nil {
				if edge, ok := noteToNoteMap[link.Id][note.Id]; ok {
					edge.SetArrowHead(cgraph.NoneArrow)
					continue
				}

				edge := render.GetEdge(graph, noteToNodeMap[note.Id], linkNode, "link"+note.Id+link.Id)
				noteToNoteMap[note.Id][link.Id] = edge
			}
		}
	}

	log.Println("rendering to file")
	render.SaveGraphToFile(g, graph, string(outputPath))
	log.Println("graph saved to:", outputPath)
}

func renderMainNodesGroup(noteToNodeMap map[string]*cgraph.Node, note *core.Note) {
	node := noteToNodeMap[note.Id]

	render.MarkMainNode(node)
	node.SetGroup(render.MainNodeGroup)

	for _, link := range note.Links {
		if linkNode := noteToNodeMap[link.Id]; linkNode != nil {
			linkNode.SetGroup(render.MainNodeGroup)
			linkNode.SetColor(render.DirectLinksColor)
		}
	}
}

func getNotesForSubgraph(note *core.Note, levelsLeft int, ignoreTags []string) []*core.Note {
	var result []*core.Note

	notesMap := getNotesForSubgraphRecursive(note, levelsLeft, ignoreTags, make(map[*core.Note]struct{}))
	for key := range notesMap {
		result = append(result, key)
	}

	return result
}

func getNotesForSubgraphRecursive(note *core.Note, levelsLeft int, ignoreTags []string,
	result map[*core.Note]struct{}) map[*core.Note]struct{} {
	if levelsLeft <= 0 {
		return result
	}

	if len(note.Tags) > 0 {
		for _, tag := range ignoreTags {
			if _, ok := note.Tags[tag]; ok {
				log.Printf("[INFO] note with Path '%s' ignored because of tag '%s'\n", note.Path, tag)
				return result
			}
		}
	}

	result[note] = struct{}{}
	var addedNotes []*core.Note

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
		result = getNotesForSubgraphRecursive(addedNote, levelsLeft-1, ignoreTags, result)
	}

	return result
}

func getNotes(notePath, folderPath sys.Path) (*core.Note, []*core.Note) {
	isZettel, id, _ := zk.ParseNoteFilename(sys.GetFilename(notePath))
	if !isZettel {
		log.Fatal(fmt.Errorf("provided note filename is not a correct zk note"))
	}

	notes, err := core.GetNotes(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	var targetNote *core.Note
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

func parseArguments() (sys.Path, sys.Path, int, []string, sys.Path, error) {
	notePath := flag.String("notePath", "", "a path to note file")
	outputPath := flag.String("outputPath", "./", "a path to rendered graph file")
	graphDepth := flag.Int("depth", 2, "graph depth to render")
	ignoreTags := flag.String("ignoreTags", "", "comma seperated list of tags which is used to ignore/filter notes")
	flag.Parse()

	if filepath.Ext(*notePath) != sys.MdExtension {
		return "", "", -1, nil, "", fmt.Errorf("specify %s path for generating graph", sys.MdExtension)
	}

	if *notePath == "" || *outputPath == "" {
		return "", "", -1, nil, "", fmt.Errorf("provide both 'notePath' and 'outputPath'")
	}

	return sys.Path(*notePath), sys.Path(filepath.Dir(*notePath)),
		*graphDepth, strings.Split(*ignoreTags, ","), sys.Path(*outputPath),
		nil
}
