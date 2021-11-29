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


func ExampleDepthFirstDirectedPaths() {

	source := 3
	g := *tinyDigraph
	path := graphs.NewDepthFirstDirectedPaths(g, source)

	for v := 0; v < g.V(); v++ {
		if path.HasPathTo(v) {
			fmt.Printf("%d to %d: ", source, v)
			// fmt.Println(path.PathTo(v))
			for _, x := range path.PathTo(v) {
				if x == source {
					fmt.Print(x)
				} else {
					fmt.Printf("-%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d: not connected\n", source, v)
		}
	}

	// Output: 
	// 3 to 0: 3-5-4-2-0
	// 3 to 1: 3-5-4-2-0-1
	// 3 to 2: 3-5-4-2
	// 3 to 3: 3
	// 3 to 4: 3-5-4
	// 3 to 5: 3-5
	// 3 to 6: not connected
	// 3 to 7: not connected
	// 3 to 8: not connected
	// 3 to 9: not connected
	// 3 to 10: not connected
	// 3 to 11: not connected
	// 3 to 12: not connected
}

func ExampleBreadthFirstDirectedPaths() {
	if tinyCG == nil {
		dataInit()
	}

	g, source := *tinyDigraph, 3
	path := graphs.NewBreadthFirstDirectedPaths(g, source)

	for v := 0; v < g.V(); v++ {
		if path.HasPathTo(v) {
			fmt.Printf("%d to %d (%d): ", source, v, path.DistTo(v))
			// fmt.Println(path.PathTo(v))
			for _, x := range path.PathTo(v) {
				if x == source {
					fmt.Print(x)
				} else {
					fmt.Printf("->%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d (-): not connected\n", source, v)
		}
	}

	// Output: 
	// 3 to 0 (2): 3->2->0
	// 3 to 1 (3): 3->2->0->1
	// 3 to 2 (1): 3->2
	// 3 to 3 (0): 3
	// 3 to 4 (2): 3->5->4
	// 3 to 5 (1): 3->5
	// 3 to 6 (-): not connected
	// 3 to 7 (-): not connected
	// 3 to 8 (-): not connected
	// 3 to 9 (-): not connected
	// 3 to 10 (-): not connected
	// 3 to 11 (-): not connected
	// 3 to 12 (-): not connected	
}