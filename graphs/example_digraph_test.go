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

func ExampleDepthFirstSearch_singleSource() {
	dfs := graphs.NewDirectedDFS(*tinyDigraph, 2)

	// print out vertices reachable from soure
	for v := 0; v < tinyDigraph.V(); v++ {
		if dfs.Marked(v) {
			fmt.Printf("%v ", v)
		}
	}

	// Output:
	// 0 1 2 3 4 5
}

func ExampleDepthFirstSearch_multipleSource() {
	sources := []int{1, 2, 6}
	dfs := graphs.NewDirectedDFSN(*tinyDigraph, sources)

	// print out vertices reachable from soure
	for v := 0; v < tinyDigraph.V(); v++ {
		if dfs.Marked(v) {
			fmt.Printf("%v ", v)
		}
	}

	// Output:
	// 0 1 2 3 4 5 6 8 9 10 11 12
}