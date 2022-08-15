package render

import (
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
	"strings"
)

const MainNodeGroup = "main"
const MainNodeColor = "red"
const DirectLinksColor = "blue"
const IndexNodeColor = "green"
const MaxLineLengthChars = 32

func InitGraphviz() (g *graphviz.Graphviz, graph *cgraph.Graph, finishFunc func()) {
	g = graphviz.New()

	var err error
	graph, err = g.Graph()
	if err != nil {
		log.Fatal(err)
	}

	return g, graph, func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}

		err = g.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetNode(graph *cgraph.Graph, noteName string, tags map[string]any) *cgraph.Node {
	if len(noteName) > MaxLineLengthChars {
		words := strings.Split(noteName, " ")
		header := strings.Builder{}

		for i, word := range words {
			header.WriteString(word)
			header.WriteRune(' ')

			if len(words)/2 == i {
				header.WriteString("\n")
			}
		}

		noteName = header.String()
	}

	node, err := graph.CreateNode(noteName)
	if err != nil {
		log.Fatal(err)
	}

	//node.SetShape(cgraph.OctagonShape)

	if _, ok := tags["index"]; ok {
		node.SetColor(IndexNodeColor)
		node.SetStyle(cgraph.BoldNodeStyle)
	}

	return node
}

func GetEdge(graph *cgraph.Graph, start, end *cgraph.Node, name string) *cgraph.Edge {
	edge, err := graph.CreateEdge(name, start, end)
	if err != nil {
		log.Fatal(err)
	}

	return edge
}

func SaveGraphToFile(g *graphviz.Graphviz, graph *cgraph.Graph, graphPath string) {
	if err := g.RenderFilename(graph, graphviz.PNG, graphPath); err != nil {
		log.Fatal(err)
	}
}
