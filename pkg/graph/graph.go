package graph

import (
	"fmt"

	"github.com/andrewbloese-00/gopher_graph/pkg/strains"
)

type Graph struct {
	Size  int
	Nodes map[int]*Node
}

type Node struct {
	Id        int
	Neighbors []*Node
	Weights   map[int]float64 //weights to neighbors
	Strain    *strains.StrainEntry
}

func NewStrainGraph() *Graph {
	return &Graph{Size: 0, Nodes: make(map[int]*Node)}
}

func NewNode(id int, strain *strains.StrainEntry) *Node {
	return &Node{Id: id, Neighbors: make([]*Node, 0), Weights: make(map[int]float64), Strain: strain}
}

func (g *Graph) AddNode(n *Node) bool {
	_, hasNode := g.Nodes[n.Id]
	if hasNode {
		return false
	}

	g.Nodes[n.Id] = n
	return true

}

func (g *Graph) AddEdgeDirected(u *Node, v *Node, weight float64) bool {
	//ensure nodes are in the graph
	_, hasU := g.Nodes[u.Id]
	_, hasV := g.Nodes[v.Id]

	if !hasU || !hasV {
		return false
	}

	return u.AddEdgeDirected(v, weight)
}

func (g *Graph) Print(start *Node, visited *[]*Node, cost float64) {
	fmt.Printf("Visits: %d | Node: %d | Edge Cost : %.1f \n ", len(*visited), start.Id, cost)
	*visited = append(*visited, start)
	for _, neighbor := range start.Neighbors {
		if !HasNode(visited, neighbor) {
			g.Print(neighbor, visited, start.Weights[neighbor.Id])
		}
	}
}

// Add a directed edge from u to v on the graph
func (u *Node) AddEdgeDirected(v *Node, weight float64) bool {
	//check if edge already exists
	existsAtU := u.HasNeighbor(v)
	existsAtV := v.HasNeighbor(u)
	if existsAtU > -1 || existsAtV > -1 {
		fmt.Println("An edge already exists between 'u' and 'v'")
		return false
	}

	u.Neighbors = append(u.Neighbors, v)
	u.Weights[v.Id] = weight
	return true
}

// check if neighbor exists on a node and return its index if found (-1 otherwise)
func (n *Node) HasNeighbor(other *Node) int {
	for i, neighbor := range n.Neighbors {
		if neighbor == other {
			return i
		}
	}
	return -1
}

func HasNode(arr *[]*Node, node *Node) bool {
	for _, n := range *arr {
		if n == node {
			return true
		}
	}
	return false
}
