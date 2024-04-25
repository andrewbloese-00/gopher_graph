package main

import (
	"fmt"

	"github.com/andrewbloese-00/gopher_graph/internal/reader"
	"github.com/andrewbloese-00/gopher_graph/pkg/graph"
)

func main() {
	// router := routes.NewRouter()
	// port := 8080
	// addr := fmt.Sprintf(":%d", port)
	// fmt.Printf("Server listening on http://localhost%s\n", addr)
	// err := http.ListenAndServe(addr, router)
	// if err != nil {
	// 	panic(err)
	// }
	data := reader.ReadCannabisData("/Users/blaze/Development/gopher_graph/cannabis.csv")
	g, ok := graph.NewStrainGraph(), true
	fmt.Println("Rows: ", len(data))

	for i, strain := range data {
		g.AddNode(graph.NewNode(i, &strain))
		for j, o_strain := range data {
			if !ok {
				break
			}
			if i == j {
				continue
			}
			_, has_j := g.Nodes[j]
			if !has_j {
				g.AddNode(graph.NewNode(j, &o_strain))
			}
			commE, commF := strain.CommonEffectsWith(&o_strain), strain.CommonFlavorsWith(&o_strain)
			w1, w2 := len(commE), len(commF)
			w := float64(w1 + w2)
			if w <= 4 {
				continue
			}

			ok = g.AddEdgeDirected(g.Nodes[i], g.Nodes[j], float64(w))

		}
	}
	// vis := make([]*graph.Node, 0)
	// g.Print(g.Nodes[0], &vis, 0)
	// fmt.Println("Neighbors\n", g.Nodes[250].Neighbors)
	fmt.Printf("Neighbors of %v\n", g.Nodes[2000].Strain.Name)
	for _, neighbor := range g.Nodes[2000].Neighbors {
		fmt.Println("Name: ", neighbor.Strain.Name)
		fmt.Println("Effects Shared: ", g.Nodes[0].Strain.CommonEffectsWith(neighbor.Strain))
		fmt.Println("Flavors Shared: ", g.Nodes[0].Strain.CommonFlavorsWith(neighbor.Strain))

	}

}
