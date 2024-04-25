package tests

import (
	"fmt"

	"github.com/andrewbloese-00/gopher_graph/pkg/graph"
)

func Testgraph() {
	fmt.Println("\n\nTesting Graph (no strains)")
	//init graph
	g := graph.NewStrainGraph()

	//create nodes
	a, b, c, d, e, f :=
		graph.NewNode(1000),
		graph.NewNode(1001),
		graph.NewNode(1002),
		graph.NewNode(1003),
		graph.NewNode(1004),
		graph.NewNode(1005)

	//add nodes to graph
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddNode(d)
	g.AddNode(e)
	g.AddNode(f)

	//define some edges
	g.AddEdgeDirected(a, b, 1)
	g.AddEdgeDirected(a, d, 1)
	g.AddEdgeDirected(b, c, 1)
	g.AddEdgeDirected(d, e, 1)
	g.AddEdgeDirected(b, f, 1)
	g.AddEdgeDirected(f, a, 1)

	//print the graph starting at 'a', (dfs driven)
	visited := make([]*graph.Node, 0)
	g.Print(a, &visited, 0)

	fmt.Printf("Visited Array (Node Pointers):\n%v\n", visited)

}
