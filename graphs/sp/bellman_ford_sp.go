package sp

import (
	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/graphs/digraph"
	"github.com/youngzhu/algs4-go/fund"
)

// Queue-based Bellman-Ford algorithm.
// The only edges that could lead to a change in distTo[] are those leaving a 
// vertex whose distTo[] value changed in the previous pass. To keep track of 
// such vertices, we use a FIFO queue.
type BellmanFordSP struct {
	graph digraph.EdgeWeightedDigraph
	source int
	distTo []graphs.Distance // distTo[v]: distance of shortest s->v path
	edgeTo []*digraph.DirectedEdge // edgeTo[v]: last edge on shortest s->v path
	onQueue []bool // onQueue[v]: is v currently on the queue?
	queue *fund.Queue // queue of vertices to relax
	cost int // number of calls to relax()
	cycle fund.Iterator // negative cycle (or nil if no such cycle)
}

func NewBellmanFordSP(g digraph.EdgeWeightedDigraph, s int) *BellmanFordSP {
	n := g.V()
	distTo := make([]graphs.Distance, n)
	edgeTo := make([]*digraph.DirectedEdge, n)
	onQueue := make([]bool, n)

	for i := 0; i < n; i++ {
		distTo[i] = graphs.DistanceInfinity
	}
	distTo[s] = graphs.DistanceZero

	queue := fund.NewQueue()
	queue.Enqueue(s)
	onQueue[s] = true

	// Bellman-Ford algorithm
	bf := &BellmanFordSP{graph: g,
			source: s,
		distTo: distTo,
		edgeTo: edgeTo, 
		onQueue: onQueue,
		queue: queue,}

	for !bf.queue.IsEmpty() && !bf.HasNegativeCycle() {
		v := queue.Dequeue().(int)
		onQueue[v] = false
		bf.relax(v)
	}

	return bf
}

// relax vertex v and put other endpoints on queue if changed
func (bf *BellmanFordSP) relax(v int) {
	for _, edge := range bf.graph.Adj(v) {
		e := edge.(*digraph.DirectedEdge)
		w := e.To()
		distance := graphs.Distance(e.Weight())
		if bf.distTo[w] > bf.distTo[v] + distance {
			bf.distTo[w] = bf.distTo[v] + distance
			bf.edgeTo[w] = e
			if !bf.onQueue[w] {
				bf.queue.Enqueue(w)
				bf.onQueue[w] = true
			}
		}

		bf.cost++
		if bf.cost % bf.graph.V() == 0 {
			bf.findNegativeCycle()
			if bf.HasNegativeCycle() {
				return // found a negative cycle
			}
		}
	}
}

// Is there a negative cycle reachable from the source vertex
func (bf *BellmanFordSP) HasNegativeCycle() bool {
	return bf.cycle != nil
}

// Returns a negative cycle reachable from the source vertex
// or nil if there is no such cycle
func (bf *BellmanFordSP) NegativeCycle() fund.Iterator {
	return bf.cycle
}

// finding a cycle in predecessor graph
func (bf *BellmanFordSP) findNegativeCycle() {
	n := len(bf.edgeTo)

	spt := digraph.NewEdgeWeightedDigraphN(n)
	for v := 0; v < n; v++ {
		if bf.edgeTo[v] != nil {
			spt.AddEdge(bf.edgeTo[v])
		}
	}

	finder := digraph.NewEdgeWeightedDirectedCycle(*spt)
	bf.cycle = finder.Cycle()
}

// Returns the length of a shortest path from the source vertex s to v
func (bf BellmanFordSP) DistTo(v int) float64 {
	bf.validateVertex(v)
	if bf.HasNegativeCycle() {
		panic("Negative cost cycle exists")
	}
	return float64(bf.distTo[v])
}

// Is there a path from the source s to v
func (bf BellmanFordSP) HasPathTo(v int) bool {
	bf.validateVertex(v)
	return bf.distTo[v] > graphs.DistanceNegativeInfinity
}

// Returns a shortest path from the source vertex to vertex v
func (bf BellmanFordSP) PathTo(v int) fund.Iterator {
	bf.validateVertex(v)
	if bf.HasNegativeCycle() {
		panic("Negative cost cycle exists")
	}

	stack := fund.NewStack()

	if bf.HasPathTo(v) {
		for e := bf.edgeTo[v]; e != nil; e = bf.edgeTo[e.From()] {
			stack.Push(e)
		}
	}

	return stack.Iterator()
}

func (bf *BellmanFordSP) validateVertex(v int) {
	if v < 0 || v >= len(bf.distTo) {
		panic("invalidate vertex")
	}
}