package mst_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs/mst"
	"github.com/youngzhu/algs4-go/testutil"
)

var (
	tinyEWG, mediumEWG *mst.EdgeWeightedGraph
)

func init() {
	in := testutil.NewInReadWords("testdata/tinyEWG.txt")
	tinyEWG = mst.NewEdgeWeightedGraphIn(in)

	in = testutil.NewInReadWords("testdata/mediumEWG.txt.gz")
	mediumEWG = mst.NewEdgeWeightedGraphIn(in)
}

func ExampleEdgeWeightedGraph() {

	fmt.Print(tinyEWG)

	// Output:
	// vertices:8, edges:16
	// 0: 6-0 0.58000 0-2 0.26000 0-4 0.38000 0-7 0.16000
	// 1: 1-3 0.29000 1-2 0.36000 1-7 0.19000 1-5 0.32000
	// 2: 6-2 0.40000 2-7 0.34000 1-2 0.36000 0-2 0.26000 2-3 0.17000
	// 3: 3-6 0.52000 1-3 0.29000 2-3 0.17000
	// 4: 6-4 0.93000 0-4 0.38000 4-7 0.37000 4-5 0.35000
	// 5: 1-5 0.32000 5-7 0.28000 4-5 0.35000
	// 6: 6-4 0.93000 6-0 0.58000 3-6 0.52000 6-2 0.40000
	// 7: 2-7 0.34000 1-7 0.19000 0-7 0.16000 5-7 0.28000 4-7 0.37000
}

func ExampleLazyPrimMST_tinyEWG() {
	mst := mst.NewLazyPrimMST(*tinyEWG)

	for _, e := range mst.Edges() {
		fmt.Println(e)
	}

	fmt.Printf("%.5f\n", mst.Weight())

	// Output:
	// 0-7 0.16000
	// 1-7 0.19000
	// 0-2 0.26000
	// 2-3 0.17000
	// 5-7 0.28000
	// 4-5 0.35000
	// 6-2 0.40000
	// 1.81000
}

func ExampleLazyPrimMST_mediumEWG() {
	mst := mst.NewLazyPrimMST(*mediumEWG)

	for _, e := range mst.Edges() {
		fmt.Println(e)
	}

	fmt.Printf("%.5f\n", mst.Weight())
}

func ExamplePrimMST_tinyEWG() {
	mst := mst.NewPrimMST(*tinyEWG)

	for _, e := range mst.Edges() {
		fmt.Println(e)
	}

	fmt.Printf("%.5f\n", mst.Weight())

	// Output:
	// 1-7 0.19000
	// 0-2 0.26000
	// 2-3 0.17000
	// 4-5 0.35000
	// 5-7 0.28000
	// 6-2 0.40000
	// 0-7 0.16000
	// 1.81000
}

func ExampleKruskalMST() {
	mst := mst.NewKruskalMST(*tinyEWG)

	for _, e := range mst.Edges() {
		fmt.Println(e)
	}

	fmt.Printf("%.5f\n", mst.Weight())

	// Output:
	// 0-7 0.16000
	// 2-3 0.17000
	// 1-7 0.19000
	// 0-2 0.26000
	// 5-7 0.28000
	// 4-5 0.35000
	// 6-2 0.40000
	// 1.81000
}
