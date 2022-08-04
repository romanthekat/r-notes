package render

import (
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
)

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
		g.Close()
	}
}

func GetNode(graph *cgraph.Graph, name string) *cgraph.Node {
	node, err := graph.CreateNode(name)
	if err != nil {
		log.Fatal(err)
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
