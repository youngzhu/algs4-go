package graphs_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/util"
	"github.com/youngzhu/algs4-go/fund"
)

var (
	tinyDigraph *graphs.Digraph
	tinyDAG graphs.Digraph
)

func init() {
	in := util.NewInReadWords("testdata/tinyDG.txt")
	tinyDigraph = graphs.NewDigraph(in)

	in = util.NewInReadWords("testdata/tinyDAG.txt")
	tinyDAG = *graphs.NewDigraph(in)
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

func findDirectedCycle(g graphs.Digraph) {
	dc := graphs.NewDirectedCycle(g)

	if dc.HasCycle() {
		fmt.Print("Directed cycle: ")
		for _, v := range dc.Cycle() {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	} else {
		fmt.Println("No directed cycle")
	}
}

func ExampleDirectedCycle() {
	
	findDirectedCycle(*tinyDigraph)

	// Output:
	// Directed cycle: 3 5 4 3 
}

// DAG: Directed Acyclic Graph
// a digraph with no directed cycles
func ExampleDirectedCycle_dag() {

	findDirectedCycle(tinyDAG)

	// Output:
	// No directed cycle
}

func ExampleDepthFirstOrder() {

	dfo := graphs.NewDepthFirstOrder(tinyDAG)

	fmt.Printf("%-4s %-4s %-4s\n", "v", "pre", "post")
	fmt.Println("----------------")

	for v := 0; v < tinyDAG.V(); v++ {
		fmt.Printf("%-4d %-4d %-4d\n", v, dfo.Pre(v), dfo.Post(v))
	}

	fmt.Print("Preorder: ")
	for _, v := range dfo.Preorder() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Print("Postorder: ")
	for _, v := range dfo.Postorder() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Print("Reverse Postorder: ")
	for _, v := range dfo.ReversePostorder() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// literally match. Fail probably because of CRLF
	// Output:
	// v    pre  post
	// ----------------
	// 0    0    8
	// 1    3    2
	// 2    9    10
	// 3    10   9
	// 4    2    0
	// 5    1    1
	// 6    4    7
	// 7    11   11
	// 8    12   12
	// 9    5    6
	// 10   8    5
	// 11   6    4
	// 12   7    3
	// Preorder: 0 5 4 1 6 9 11 12 10 2 3 7 8
	// Postorder: 4 5 1 12 11 10 9 6 0 3 2 7 8
	// Reverse Postorder: 8 7 2 3 0 6 9 10 11 12 1 5 4

}

func ExampleSymbolDigraph() {
	sg := graphs.NewSymbolDigraph("testdata/routes.txt", " ")

	symbolDigraph(sg, "JFK")
	symbolDigraph(sg, "ATL")
	symbolDigraph(sg, "LAX")

	// Output:
	// JFK
	//     ORD
	//     ATL
	//     MCO
	// ATL
	//     MCO
	//     HOU
	// LAX
}

func symbolDigraph(sg graphs.SymbolDigraph, input string) {
	fmt.Println(input)

	s := sg.Index(input)
	g := sg.Digraph()
	for _, v := range g.Adj(s) {
		fmt.Printf("    %v\n", sg.Name(v))
	}

}

func ExampleTopological() {

	sg := graphs.NewSymbolDigraph("testdata/jobs.txt", "/")
	topologial := graphs.NewTopological(sg.Digraph())

	for _, v := range topologial.Order() {
		fmt.Println(sg.Name(v))
	}

	// Output:
	// Calculus
	// Linear Algebra
	// Introduction to CS
	// Advanced Programming
	// Algorithms
	// Theoretical CS
	// Artificial Intelligence
	// Robotics
	// Machine Learning
	// Neural Networks
	// Databases
	// Scientific Computing
	// Computational Biology

}


func sccExample(g graphs.Digraph) {
	scc := graphs.NewKosarajuSharirSCC(g)

	// number of connected components
	n := scc.Count()
	fmt.Printf("components: %d\n", n)

	// compute list of vertices in each connected component
	components := make([]*fund.Queue, n)
	for i := 0; i < n; i++ {
		components[i] = fund.NewQueue()
	}
	for v := 0; v < g.V(); v++ {
		components[scc.Id(v)].Enqueue(fund.Item(v))
	}

	// print results
	for i := 0; i < n; i++ {
		fmt.Printf("id-%d:", i)
		for _, v := range components[i].Iterator() {
			fmt.Printf(" %v", v)
		}
		fmt.Println()
	}
}

func ExampleKosarajuSharirSCC_tinyDG() {

	g := *tinyDigraph
	sccExample(g)
	
	// Output: 
	// components: 5
	// id-0: 1
	// id-1: 0 2 3 4 5
	// id-2: 9 10 11 12
	// id-3: 6 8
	// id-4: 7

}

func ExampleKosarajuSharirSCC_mediumDG() {
	in := util.NewInReadWords("testdata/mediumDG.txt")
	g := graphs.NewDigraph(in)

	sccExample(*g)
	
	// Output: 
	// components: 10
	// id-0: 21
	// id-1: 2 5 6 8 9 11 12 13 15 16 18 19 22 23 25 26 28 29 30 31 32 33 34 35 37 38 39 40 42 43 44 46 47 48 49
	// id-2: 14
	// id-3: 3 4 17 20 24 27 36
	// id-4: 41
	// id-5: 7
	// id-6: 45
	// id-7: 1
	// id-8: 0
	// id-9: 10

}
