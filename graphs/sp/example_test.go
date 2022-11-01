package sp_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs/digraph"
	"github.com/youngzhu/algs4-go/graphs/sp"
	"github.com/youngzhu/algs4-go/testutil"
)

var (
	tinyEWD, tinyEWDAG *digraph.EdgeWeightedDigraph
)

func init() {
	in := testutil.NewInReadWords("testdata/tinyEWD.txt")
	tinyEWD = digraph.NewEdgeWeightedDigraphIn(in)

	in = testutil.NewInReadWords("testdata/tinyEWDAG.txt")
	tinyEWDAG = digraph.NewEdgeWeightedDigraphIn(in)
}

func printShortestPath(dsp sp.ShortestPaths, g digraph.EdgeWeightedDigraph, s int) {
	for v := 0; v < g.V(); v++ {
		if dsp.HasPathTo(v) {
			fmt.Printf("%d to %d (%.2f)", s, v, dsp.DistTo(v))
			for _, e := range dsp.PathTo(v) {
				fmt.Printf(" %v", e)
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d        no path\n", s, v)
		}
	}
}

func ExampleDijkstraSP() {
	s := 0
	dsp := sp.NewDijkstraSP(*tinyEWD, s)

	printShortestPath(*dsp, *tinyEWD, s)

	// Output:
	// 0 to 0 (0.00)
	// 0 to 1 (1.05) 0->4  0.38 4->5  0.35 5->1  0.32
	// 0 to 2 (0.26) 0->2  0.26
	// 0 to 3 (0.99) 0->2  0.26 2->7  0.34 7->3  0.39
	// 0 to 4 (0.38) 0->4  0.38
	// 0 to 5 (0.73) 0->4  0.38 4->5  0.35
	// 0 to 6 (1.51) 0->2  0.26 2->7  0.34 7->3  0.39 3->6  0.52
	// 0 to 7 (0.60) 0->2  0.26 2->7  0.34
}

func ExampleDijkstraSP_noPath() {
	in := testutil.NewInReadWords("testdata/nopath.txt")
	nopath := digraph.NewEdgeWeightedDigraphIn(in)

	s := 7
	dsp := sp.NewDijkstraSP(*nopath, s)

	printShortestPath(*dsp, *nopath, s)

	// Output:
	// 7 to 0        no path
	// 7 to 1 (0.60) 7->5  0.28 5->1  0.32
	// 7 to 2 (1.31) 7->3  0.39 3->6  0.52 6->2  0.40
	// 7 to 3 (0.39) 7->3  0.39
	// 7 to 4 (0.63) 7->5  0.28 5->4  0.35
	// 7 to 5 (0.28) 7->5  0.28
	// 7 to 6 (0.91) 7->3  0.39 3->6  0.52
	// 7 to 7 (0.00)
}

func ExampleAcyclicSP() {
	s := 5
	dsp := sp.NewAcyclicSP(*tinyEWDAG, s)

	printShortestPath(dsp, *tinyEWDAG, s)

	// Output:
	// 5 to 0 (0.73) 5->4  0.35 4->0  0.38
	// 5 to 1 (0.32) 5->1  0.32
	// 5 to 2 (0.62) 5->7  0.28 7->2  0.34
	// 5 to 3 (0.61) 5->1  0.32 1->3  0.29
	// 5 to 4 (0.35) 5->4  0.35
	// 5 to 5 (0.00)
	// 5 to 6 (1.13) 5->1  0.32 1->3  0.29 3->6  0.52
	// 5 to 7 (0.28) 5->7  0.28
}

func ExampleAcyclicLP() {
	s := 5
	alp := sp.NewAcyclicLP(*tinyEWDAG, s)

	printShortestPath(alp, *tinyEWDAG, s)

	// Output:
	// 5 to 0 (2.44) 5->1  0.32 1->3  0.29 3->6  0.52 6->4  0.93 4->0  0.38
	// 5 to 1 (0.32) 5->1  0.32
	// 5 to 2 (2.77) 5->1  0.32 1->3  0.29 3->6  0.52 6->4  0.93 4->7  0.37 7->2  0.34
	// 5 to 3 (0.61) 5->1  0.32 1->3  0.29
	// 5 to 4 (2.06) 5->1  0.32 1->3  0.29 3->6  0.52 6->4  0.93
	// 5 to 5 (0.00)
	// 5 to 6 (1.13) 5->1  0.32 1->3  0.29 3->6  0.52
	// 5 to 7 (2.43) 5->1  0.32 1->3  0.29 3->6  0.52 6->4  0.93 4->7  0.37
}

func ExampleCPM() {
	sp.CPM("testdata/jobsPC.txt")

	// Output:
	// job    start   finish
	// --------------------
	//    0     0.0    41.0
	//    1    41.0    92.0
	//    2   123.0   173.0
	//    3    91.0   127.0
	//    4    70.0   108.0
	//    5     0.0    45.0
	//    6    70.0    91.0
	//    7    41.0    73.0
	//    8    91.0   123.0
	//    9    41.0    70.0
	// Finish time:   173.0
}

func printBellmanFordSP(bf sp.BellmanFordSP, g digraph.EdgeWeightedDigraph, s int) {
	// print negative cycle
	if bf.HasNegativeCycle() {
		for _, e := range bf.NegativeCycle() {
			fmt.Println(e)
		}
		return
	}

	printShortestPath(bf, g, s)
}

func ExampleBellmanFordSP() {
	in := testutil.NewInReadWords("testdata/tinyEWDn.txt")
	tinyEWDn := digraph.NewEdgeWeightedDigraphIn(in)

	s := 0
	bf := sp.NewBellmanFordSP(*tinyEWDn, s)

	printBellmanFordSP(*bf, *tinyEWDn, s)

	// Output:
	// 0 to 0 (0.00)
	// 0 to 1 (0.93) 0->2  0.26 2->7  0.34 7->3  0.39 3->6  0.52 6->4 -1.25 4->5  0.35 5->1  0.32
	// 0 to 2 (0.26) 0->2  0.26
	// 0 to 3 (0.99) 0->2  0.26 2->7  0.34 7->3  0.39
	// 0 to 4 (0.26) 0->2  0.26 2->7  0.34 7->3  0.39 3->6  0.52 6->4 -1.25
	// 0 to 5 (0.61) 0->2  0.26 2->7  0.34 7->3  0.39 3->6  0.52 6->4 -1.25 4->5  0.35
	// 0 to 6 (1.51) 0->2  0.26 2->7  0.34 7->3  0.39 3->6  0.52
	// 0 to 7 (0.60) 0->2  0.26 2->7  0.34
}

func ExampleBellmanFordSP_negativeCycle() {
	in := testutil.NewInReadWords("testdata/tinyEWDnc.txt")
	tinyEWDnc := digraph.NewEdgeWeightedDigraphIn(in)

	s := 0
	bf := sp.NewBellmanFordSP(*tinyEWDnc, s)

	printBellmanFordSP(*bf, *tinyEWDnc, s)

	// Output:
	// 4->5  0.35
	// 5->4 -0.66
}

func ExampleArbitrage() {
	sp.Arbitrage("testdata/rates.txt")

	// Output:
	// 1000.00000 USD =  741.00000 EUR
	//  741.00000 EUR = 1012.20600 CAD
	// 1012.20600 CAD = 1007.14497 USD
}
