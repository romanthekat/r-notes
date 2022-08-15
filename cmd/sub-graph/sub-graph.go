package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/render"
	"github.com/romanthekat/r-notes/pkg/subgraph"
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
	log.Printf("mainNote path: %s\n", notePath)
	log.Printf("graph output path: %s\n", outputPath)
	log.Printf("graph depth: %d\n", graphDepth)
	log.Printf("tags to ignore: %s\n", ignoreTags)

	log.Println("obtaining notes")
	mainNote, _ := getNotes(notePath, folderPath)

	log.Println("preparing graph")
	g, graph, finishFunc := render.InitGraphviz()
	defer finishFunc()

	log.Println("getting notes for subgraph")
	notes := subgraph.GetNotesForSubgraph(mainNote, graphDepth, ignoreTags)
	log.Println("notes in graph:", len(notes))

	nodesMap := subgraph.GetNodes(notes, graph)
	subgraph.RenderMainNodes(nodesMap, mainNote)

	edgesMap := make(map[string]map[string]*cgraph.Edge)
	for _, note := range notes {
		edgesMap[note.Id] = make(map[string]*cgraph.Edge)

		for _, link := range note.Links {
			if linkNode := nodesMap[link.Id]; linkNode != nil {
				if edge, ok := edgesMap[link.Id][note.Id]; ok {
					edge.SetArrowHead(cgraph.NoneArrow)
					continue
				}

				edge := render.GetEdge(graph, nodesMap[note.Id], linkNode, "link"+note.Id+link.Id)
				if note.Id == mainNote.Id {
					edge.SetWeight(8.0)
				}
				edgesMap[note.Id][link.Id] = edge
			}
		}
	}

	log.Println("rendering to file")
	render.SaveGraphToFile(g, graph, string(outputPath))
	log.Println("graph saved to:", outputPath)
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
