package graphs_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/util"
)
func ExampleGraph() {
	in := util.NewInReadWords("testdata/tinyG.txt")
	graph := graphs.NewGraph(in)

	fmt.Println(graph.String())

	// Output:
	// 13 vertices, 13 edges
	// 0: 6 2 1 5
	// 1: 0
	// 2: 0
	// 3: 5 4
	// 4: 5 6 3
	// 5: 3 4 0
	// 6: 0 4
	// 7: 8
	// 8: 7
	// 9: 11 10 12
	// 10: 9
	// 11: 9 12
	// 12: 11 9
}