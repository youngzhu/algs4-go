package mst

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/sorting/pq"
)

// Eager implementation. To improve the lazy implementation of Prim's algorithm,
// we might try to delete ineligible edges from the priority queue, so that the
// priority queue contains only the crossing edges. But we can eliminate even 
// more edges. The key is to note that our only interest is in the minimal edge
// from each non-tree vertex to a tree vertex. When we add a vertex v to the tree,
// the only possible change with respect to each non-tree vertex w is that adding
// v brings w closer than before to the tree. In short, we do not need to keep on
// the priority queue all of the edges from w to verteices tree—we just need to 
// keep track of the minimum-weight edge and check whether the addition of v to the
// tree necessitates that we update that minimum (because of an edge v-w that has
// lower weight), which we can do as we process each edge in s adjacency list. In
// other words, we maintain on the priority queue just on edge for each non-tree 
// vertex: the shortest edge that connects it to the tree.

type distance float64

func (d distance) CompareTo(x pq.Item) int {
	xx := x.(distance)
	if d < xx {
		return -1
	} else if d > xx {
		return 1
	} else {
		return 0
	}
}

type PrimMST struct {
	graph EdgeWeightedGraph
	edgeTo []*Edge // edgeTo[v]: shortest edge from tree vertex to non-tree vertex
	distTo []distance // distTo[v]: weight of shortest such edge
	marked []bool // marked[v]: true if v on tree, false otherwise
	ipq *pq.MinIndexPQ 
}

const positiveInfinity = 1000000.0

func NewPrimMST(g EdgeWeightedGraph) *PrimMST {
	n := g.V()
	edgeTo := make([]*Edge, n)
	distTo := make([]distance, n)
	marked := make([]bool, n)
	ipq := pq.NewMinIndexPQ(n)

	for v := 0; v < n; v++ {
		distTo[v] = positiveInfinity
	}

	mst := &PrimMST{g, edgeTo, distTo, marked, ipq}

	// run from each vertex to find minimum spanning forest
	for v := 0; v < n; v++ {
		if !mst.marked[v] {
			mst.prim(v)
		}
	}

	return mst
}

// run Prim's algorithm starting from vertex s
func (p *PrimMST) prim(s int) {
	p.distTo[s] = 0
	p.ipq.Insert(s, p.distTo[s])
	for ! p.ipq.IsEmpty() {
		v := p.ipq.Delete()
		p.scan(v)
	}
}

// scan vertex v
func (p *PrimMST) scan(v int) {
	p.marked[v] = true
	for _, edge := range p.graph.Adj(v) {
		e := edge.(*Edge)
		w := e.Other(v)
		if p.marked[w] { // v-w is obsolete edge
			continue
		}
		weight := distance(e.Weight())
		if weight < p.distTo[w] {
			p.distTo[w] = weight
			p.edgeTo[w] = e
			if p.ipq.Contains(w) {
				p.ipq.Decrease(w, p.distTo[w])
			} else {
				p.ipq.Insert(w, p.distTo[w])
			}
		}
	}
}

// Returns the edges in a MST
func (p *PrimMST) Edges() fund.Iterator {
	mst := fund.NewQueue()
	for v := 0;  v < len(p.edgeTo); v++ {
		e := p.edgeTo[v]
		if e != nil {
			mst.Enqueue(e)
		}
	}
	return mst.Iterator()
}

// Returns the sum of the edge weights in a MST
func (p *PrimMST) Weight() float64 {
	weight := 0.0
	for _, edge := range p.Edges() {
		e := edge.(*Edge)
		weight += e.Weight()
	}
	return weight
}