package graph_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/graphs/graph"
	"github.com/youngzhu/algs4-go/testutil"
)

var (
	tinyGraph *graph.Graph
	tinyCG    *graph.Graph
)

func dataInit() {
	in := testutil.NewInReadWords("testdata/tinyG.txt")
	tinyGraph = graph.NewGraph(in)

	in = testutil.NewInReadWords("testdata/tinyCG.txt")
	tinyCG = graph.NewGraph(in)
}

func ExampleGraph() {
	if tinyGraph == nil {
		dataInit()
	}

	fmt.Println(tinyGraph)

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

func ExampleGraph_copy() {
	if tinyGraph == nil {
		dataInit()
	}

	copy := graph.NewGraphCopy(*tinyGraph)
	fmt.Println(copy)

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

func ExampleGraph_MaxDegree() {
	if tinyGraph == nil {
		dataInit()
	}

	fmt.Printf("vertex of maximum degree: %d", tinyGraph.MaxDegree())

	// Output:
	// vertex of maximum degree: 4
}

func ExampleGraph_AvgDegree() {
	if tinyGraph == nil {
		dataInit()
	}

	fmt.Printf("average degree: %d", tinyGraph.AvgDegree())

	// Output:
	// average degree: 2
}

func ExampleGraph_NumberOfSelfLoops() {
	if tinyGraph == nil {
		dataInit()
	}

	fmt.Printf("number of self loops: %d", tinyGraph.NumberOfSelfLoops())

	// Output:
	// number of self loops: 0
}

func ExampleGraph_HasEdge() {
	if tinyGraph == nil {
		dataInit()
	}

	v, w := 0, 6
	fmt.Printf("has edge %d-%d: %v\n", v, w, tinyGraph.HasEdge(v, w))

	v, w = 6, 0
	fmt.Printf("has edge %d-%d: %v\n", v, w, tinyGraph.HasEdge(v, w))

	v, w = 0, 7
	fmt.Printf("has edge %d-%d: %v\n", v, w, tinyGraph.HasEdge(v, w))

	// Output:
	// has edge 0-6: true
	// has edge 6-0: true
	// has edge 0-7: false
}

func ExampleDepthFirstSearch() {
	if tinyGraph == nil {
		dataInit()
	}

	dfs(*tinyGraph, 0)

	fmt.Println()

	dfs(*tinyGraph, 9)

	// literally match. Fail probably because of CRLF
	// Output:
	// 0 1 2 3 4 5 6
	// NOT connected
	// 9 10 11 12
	// NOT connected

}

func dfs(g graph.Graph, s int) {
	search := graph.NewDepthFirstSearch(g, s)

	for v := 0; v < g.V(); v++ {
		if search.Marked(v) {
			fmt.Printf("%v ", v)
		}
	}

	fmt.Println()

	if search.Count() != g.V() {
		fmt.Print("NOT connected")
	} else {
		fmt.Print("connected")
	}
}

func ExampleDepthFirstPaths() {

	if tinyCG == nil {
		dataInit()
	}

	source := 0
	path := graph.NewDepthFirstPaths(*tinyCG, source)

	for v := 0; v < tinyCG.V(); v++ {
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
	// 0 to 0: 0
	// 0 to 1: 0-2-1
	// 0 to 2: 0-2
	// 0 to 3: 0-2-3
	// 0 to 4: 0-2-3-4
	// 0 to 5: 0-2-3-5

}

func ExampleBreadthFirstPaths() {
	if tinyCG == nil {
		dataInit()
	}

	source := 0
	path := graph.NewBreadthFirstPaths(*tinyCG, source)

	for v := 0; v < tinyCG.V(); v++ {
		if path.HasPathTo(v) {
			fmt.Printf("%d to %d (%d): ", source, v, path.DistTo(v))
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
			fmt.Printf("%d to %d (-): not connected\n", source, v)
		}
	}

	// Output:
	// 0 to 0 (0): 0
	// 0 to 1 (1): 0-1
	// 0 to 2 (1): 0-2
	// 0 to 3 (2): 0-2-3
	// 0 to 4 (2): 0-2-4
	// 0 to 5 (1): 0-5

}

func ExampleConnectedComponents() {
	if tinyGraph == nil {
		dataInit()
	}

	cc := graph.NewConnectedComponents(*tinyGraph)

	// number of connected components
	n := cc.Count()
	fmt.Printf("%d components\n", n)

	// compute list of vertices in each connected component
	components := make([]*fund.Queue, n)
	for i := 0; i < n; i++ {
		components[i] = fund.NewQueue()
	}
	for v := 0; v < tinyGraph.V(); v++ {
		components[cc.Id(v)].Enqueue(fund.Item(v))
	}

	// print results
	for i := 0; i < n; i++ {
		// fmt.Printf("id-%d:", i)
		for _, v := range components[i].Iterator() {
			fmt.Printf(" %v", v)
		}
		fmt.Printf("\n")
	}

	// Output:
	// 3 components
	//  0 1 2 3 4 5 6
	//  7 8
	//  9 10 11 12

}

func ExampleSymbolGraph_routes() {
	sg := graph.NewSymbolGraph("testdata/routes.txt", " ")

	symbolGraph(sg, "JFK")
	symbolGraph(sg, "KFC")

	// Output:
	// JFK
	//     ORD
	//     ATL
	//     MCO
	// KFC
	//     input not contain 'KFC'
}

func ExampleSymbolGraph_movies() {
	sg := graph.NewSymbolGraph("testdata/movies.txt.gz", "/")

	symbolGraph(sg, "Tin Men (1987)")
	symbolGraph(sg, "让子弹飞")

}

func symbolGraph(sg graph.SymbolGraph, input string) {
	fmt.Println(input)

	if sg.Contains(input) {
		s := sg.Index(input)
		g := sg.Graph()
		for _, v := range g.Adj(s) {
			fmt.Printf("    %v\n", sg.Name(v.(int)))
		}
	} else {
		fmt.Printf("    input not contain '%v'\n", input)
	}
}
