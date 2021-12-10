package graphs

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/sorting/pq"
)

// Lazy implementation. We use priority queue to hold the crossing edges and find
// one of minimal weight. Each time that we add an edge to the tree, we also add
// a vertex to the tree. To maintain the set of crossing edges, we need to add 
// to the priority queue all edges from that vertex to any non-tree vertex. But
// we must to do more: andy edge connecting the vertex just added to a tree vertex
// that is already on the priority queue now becomes ineligible (it is no longer
// a crossing edge because it connects two tree vertices). The lazy implementation
// leaves such edges on the priority queue, deferring the ineligibility test to
// when we remove them.

// It relies on MinPQ
type LazyPrimMST struct {
	graph EdgeWeightedGraph

	weight float64 // total weight of MST
	mst *fund.Queue // edges in the MST
	marked []bool // marked[v]: true if v on tree
	epq *pq.MinPQ // edges with one endpoint in tree
}

func NewLazyPrimMST(g EdgeWeightedGraph) *LazyPrimMST {
	mst := fund.NewQueue()
	epq := pq.NewMinPQ()
	marked := make([]bool, g.V())

	lp := &LazyPrimMST{graph: g, mst: mst, marked: marked, epq: epq}

	// run Prim from all vertices to get a minimum spanning forest
	for v := 0; v < g.V(); v++ {
		if !lp.marked[v] {
			lp.prim(v)
		}
	}

	return lp
}

// run Prim's algorithm
func (lp *LazyPrimMST) prim(s int) {
	lp.scan(s)

	// better to stop when mst has V-1 edges smallest edge on pq
	for !lp.epq.IsEmpty() {
		e := lp.epq.Delete().(*Edge)
		// two endpoints
		v := e.Either()
		w := e.Other(v)

		// lazy, both v and w already scanned
		if lp.marked[v] && lp.marked[w] {
			continue
		}

		lp.mst.Enqueue(e) // add to MST
		lp.weight += e.weight

		// v becomes part of tree
		if !lp.marked[v] {
			lp.scan(v)
		}
		// w becomes part of tree
		if !lp.marked[w] {
			lp.scan(w)
		}
	}
}

// add all edges e incident to v onto pq 
// if the other endpoint has not yet been scanned
func (lp *LazyPrimMST) scan(v int) {
	lp.marked[v] = true
	for _, edge := range lp.graph.Adj(v) {
		e := edge.(*Edge)
		if !lp.marked[e.Other(v)] {
			lp.epq.Insert(e)
		}
	}
}

// Returns the edges in a MST
func (lp *LazyPrimMST) Edges() fund.Iterator {
	return lp.mst.Iterator()
}

// Returns the sum of the edge weights in a MST
func (lp *LazyPrimMST) Weight() float64 {
	return lp.weight
}