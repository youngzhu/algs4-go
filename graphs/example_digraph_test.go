package graphs_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/util"
)

var (
	tinyDigraph *graphs.Digraph
)

func init() {
	in := util.NewInReadWords("testdata/tinyDG.txt")
	tinyDigraph = graphs.NewDigraph(in)
}

func ExampleDigraph() {

	fmt.Println(tinyDigraph)

	// Output:
	// 13 vertices, 22 edges
	// 0: 5 1 
	// 1: 
	// 2: 0 3 
	// 3: 5 2 
	// 4: 3 2 
	// 5: 4 
	// 6: 9 4 8 0 
	// 7: 6 9
	// 8: 6 
	// 9: 11 10 
	// 10: 12 
	// 11: 4 12 
	// 12: 9 
}
