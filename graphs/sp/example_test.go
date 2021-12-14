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
