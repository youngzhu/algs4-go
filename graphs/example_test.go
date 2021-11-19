package graphs_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/util"
)

var tinyGraph *graphs.Graph

func dataInit() {
	in := util.NewInReadWords("testdata/tinyG.txt")
	tinyGraph = graphs.NewGraph(in)
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