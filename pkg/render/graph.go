package render

import (
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
)

const MainNodeGroup = "main"
const MainNodeColor = "red"
const DirectLinksColor = "blue"
const IndexNodeColor = "green"

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

func MarkMainNode(node *cgraph.Node) {
	node.SetColor(MainNodeColor)
}

func GetNode(graph *cgraph.Graph, noteName string, tags map[string]any) *cgraph.Node {
	node, err := graph.CreateNode(noteName)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := tags["index"]; ok {
		node.SetColor(IndexNodeColor)
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
