package sp

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/graphs/digraph"
)

// Computes shortest paths in an edge-weighted acyclic digraph
type AcyclicSP struct {
	graph digraph.EdgeWeightedDigraph
	source int
	distTo []graphs.Distance // distTo[v]: distance of shortest s->v path
	edgeTo []*digraph.DirectedEdge // edgeTo[v]: last edge on shortest s->v path
}

func NewAcyclicSP(g digraph.EdgeWeightedDigraph, s int) *AcyclicSP {
	n := g.V()
	distTo := make([]graphs.Distance, n)
	edgeTo := make([]*digraph.DirectedEdge, n)

	asp := &AcyclicSP{g, s, distTo, edgeTo}

	asp.validateVertex(s)

	for i := 0; i < n; i++ {
		distTo[i] = graphs.DistanceInfinity
	}
	distTo[s] = graphs.DistanceZero
	
	// visit vertices in topological order
	topo := digraph.NewTopologicalWeighted(g)
	if !topo.HasOrder() {
		panic("digraph is not acyclic")
	}
	for _, v := range topo.Order() {
		for _, edge := range g.Adj(v) {
			e := edge.(*digraph.DirectedEdge)
			asp.relax(e)
		}
	}

	return asp
}

// relax edge e
func (asp *AcyclicSP) relax(e *digraph.DirectedEdge) {
	v, w := e.From(), e.To()
	distance := graphs.Distance(e.Weight())
	if asp.distTo[w] > asp.distTo[v] + distance {
		asp.distTo[w] = asp.distTo[v] + distance
		asp.edgeTo[w] = e
	}
}

// Returns the length of a shortest path from the source vertex s to vertex v
func (asp AcyclicSP) DistTo(v int) float64 {
	asp.validateVertex(v)
	return float64(asp.distTo[v])
}

// HasPathTo return true if there is a path from the source vertex to vertx v
func (asp AcyclicSP) HasPathTo(v int) bool {
	asp.validateVertex(v)
	return asp.distTo[v] < graphs.DistanceInfinity
}

// Returns a shortest path from the source vertex to vertex v
func (asp *AcyclicSP) PathTo(v int) fund.Iterator {
	asp.validateVertex(v)
	stack := fund.NewStack()

	if asp.HasPathTo(v) {
		for e := asp.edgeTo[v]; e != nil; e = asp.edgeTo[e.From()] {
			stack.Push(e)
		}
	}

	return stack.Iterator()
}

func (t *AcyclicSP) validateVertex(v int) {
	if v < 0 || v >= len(t.distTo) {
		panic("invalidate vertex")
	}
}