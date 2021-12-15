package sp_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs/sp"
	"github.com/youngzhu/algs4-go/util"
)

var (
	tinyEWD *sp.EdgeWeightedDigraph
)

func init() {
	in := util.NewInReadWords("testdata/tinyEWD.txt")
	tinyEWD = sp.NewEdgeWeightedDigraphIn(in)
}

func ExampleEdgeWeightedDigraph() {

	fmt.Print(tinyEWD)

	// Output:
	// vertices:8, edges:15
	// 0: 0->2  0.26 0->4  0.38
	// 1: 1->3  0.29
	// 2: 2->7  0.34
	// 3: 3->6  0.52
	// 4: 4->7  0.37 4->5  0.35
	// 5: 5->1  0.32 5->7  0.28 5->4  0.35
	// 6: 6->4  0.93 6->0  0.58 6->2  0.40
	// 7: 7->3  0.39 7->5  0.28
}

func printDijkstraSP(dsp sp.DijkstraSP) {
	s := dsp.Source()
	for v := 0; v < tinyEWD.V(); v++ {
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

	printDijkstraSP(*dsp)
	
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
	in := util.NewInReadWords("testdata/nopath.txt")
	nopath := sp.NewEdgeWeightedDigraphIn(in)

	s := 7
	dsp := sp.NewDijkstraSP(*nopath, s)

	printDijkstraSP(*dsp)
	
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