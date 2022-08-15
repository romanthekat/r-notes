package subgraph

import (
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/render"
	"log"
)

const MaxBacklinks = 2

func GetNotesForSubgraph(note *core.Note, levelsLeft int, ignoreTags []string) []*core.Note {
	var result []*core.Note

	notesMap := getNotesForSubgraphRecursive(note, levelsLeft, ignoreTags, make(map[*core.Note]any))
	for key := range notesMap {
		result = append(result, key)
	}

	return result
}

func getNotesForSubgraphRecursive(
	note *core.Note,
	levelsLeft int,
	ignoreTags []string,
	result map[*core.Note]any) map[*core.Note]any {
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

	for i, link := range note.Backlinks {
		if _, ok := result[link]; !ok {
			result[link] = struct{}{}
			addedNotes = append(addedNotes, link)

			if i > MaxBacklinks {
				break
			}
		}
	}

	for _, addedNote := range addedNotes {
		result = getNotesForSubgraphRecursive(addedNote, levelsLeft-1, ignoreTags, result)
	}

	return result
}

func GetNodes(notes []*core.Note, graph *cgraph.Graph) map[string]*cgraph.Node {
	nodesMap := make(map[string]*cgraph.Node)
	for _, note := range notes {
		nodesMap[note.Id] = render.GetNode(graph, note.Name, note.Tags)
	}
	return nodesMap
}

func RenderMainNodes(noteToNodeMap map[string]*cgraph.Node, note *core.Note) {
	node := noteToNodeMap[note.Id]

	node.SetColor(render.MainNodeColor)
	node.SetGroup(render.MainNodeGroup)
	node.SetStyle(cgraph.BoldNodeStyle)

	for _, link := range note.Links {
		if linkNode := noteToNodeMap[link.Id]; linkNode != nil {
			linkNode.SetColor(render.DirectLinksColor)
			linkNode.SetGroup(render.MainNodeGroup)
			linkNode.SetStyle(cgraph.BoldNodeStyle)
		}
	}
}
