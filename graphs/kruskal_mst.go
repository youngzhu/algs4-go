package graphs

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/fund/uf"
	"github.com/youngzhu/algs4-go/sorting/pq"
)

// Kruskal's algorithm processes the edges in order of their weight values
// (smallest to largest), taking for the MST (coloring black) each edge that
// does not form a cycle with edges previously added, stopping after adding
// V-1 edges. The black edges form a forest of trees that evolves gradually
// into a single tree, the MST.
// To implement Kruskal's algorithm, we use a priority queue to consider the
// edges in order by weight, a union-find data structure to indentify those
// that cause cycles, and a queue to collect the MST edges.
type KruskalMST struct {
	weight float64 // weight of MST
	mst *fund.Queue // edges in MST
}

func NewKruskalMST(g EdgeWeightedGraph) *KruskalMST {
	minPQ := pq.NewMinPQ()

	for _, e := range g.Edges() {
		minPQ.Insert(e.(*Edge))
	}

	mst := fund.NewQueue()
	weight := 0.0

	// run greedy algorithm
	unionFind := uf.NewUF(g.V())
	for !minPQ.IsEmpty() && mst.Size() < g.V()-1 {
		e := minPQ.Delete().(*Edge)
		v := e.Either()
		w := e.Other(v)

		// v-w does not create a cycle
		if unionFind.Find(v) != unionFind.Find(w) {
			unionFind.Union(v, w) // merge v and w components
			mst.Enqueue(e) // add edge e to mst
			weight += e.Weight()
		}
	}

	return &KruskalMST{weight, mst}
}

// Returns the edges in a MST
func (k *KruskalMST) Edges() fund.Iterator {
	return k.mst.Iterator()
}

// Returns the sum of edge weights in a MST
func (k *KruskalMST) Weight() float64 {
	return k.weight
}