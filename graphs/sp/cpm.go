package sp

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs/digraph"
	"github.com/youngzhu/algs4-go/util"
)

// Critical path method.
// We consider the parallel precedence-constrained job scheduling problem: Given
// a set of jobs of specified duration to be completed, with precedence constraints
// that specify that certain jobs have to be completed before certain other jobs
// are begun, how can we schedule the jobs on identical processors (as many as needed)
// such that they are all completed in the minimum amount of time while still 
// respecting the constrains?
// 
// This problem can be solved by formulating it as a longest paths problem in an
// edge-weighted DAG: Create an edge-weighted DAG with a source a, a sink t, and
// two vertices for each job (a start vertex and an end vertex). For each job, 
// add an edge from its start vertex to its end vertex with weight equal to its
// duration. For each precedence constrain v->w, add a zero-weight edge from the
// end vertex corresponding to v to the beginning vertex corresponding to w. Also
// add zero-weight edges from the source to each job's start vertex and from each
// job's end vertex to the sink.
// Now, schedule each job at the time given by the length of its longest path 
// from the source.

// Reads the precedence constraints from file
// and prints a feasible schedule.
func CPM(path string) {
	in := util.NewInReadWords(path)

	// number of jobs
	n := in.ReadInt()

	// source and sink
	source, sink := 2*n, 2*n + 1

	// build network
	g := digraph.NewEdgeWeightedDigraphN(2*n+2)
	for i := 0; i < n; i++ {
		duration := in.ReadFloat()
		g.AddEdge(digraph.NewDirectedEdge(source, i, 0))
		g.AddEdge(digraph.NewDirectedEdge(i+n, sink, 0))
		g.AddEdge(digraph.NewDirectedEdge(i, i+n, duration))

		// precedence constraints
		m := in.ReadInt()
		for j := 0; j < m; j++ {
			precedent := in.ReadInt()
			g.AddEdge(digraph.NewDirectedEdge(n+i, precedent, 0))
		}
	}

	// compute longest path
	lp := NewAcyclicLP(*g, source)

	// print results
	columns := []string{"job", "start", "finish"}
	fmt.Printf("%4s %8s %8s\n", columns[0], columns[1], columns[2])
	fmt.Println("--------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%4d %7.1f %7.1f\n", i, lp.DistTo(i), lp.DistTo(i+n))
	}

	fmt.Printf("Finish time: %7.1f\n", lp.DistTo(sink))
}
