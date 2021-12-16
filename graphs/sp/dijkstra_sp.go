package sp

import (
	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/sorting/pq"
)

// Dijkstra's algorithm.
// Dijkstra's algorithm initalizing dist[s] to 0 and all other distTo[] entries
// to positive infinity. Then, it repeatedly relaxes and adds to the tree a 
// non-tree vertex with the lowest distTo[] value, counting until all vertices
// are on the tree or non-tree vertex has a finite distTo[] value.

// Assumes all weights are non-negative.

type DijkstraSP struct {
	graph EdgeWeightedDigraph
	source int
	distTo []graphs.Distance // distTo[v]: distance of shortest s->v path
	edgeTo []*DirectedEdge // edgeTo[v]: last edge on shortest s->v path
	ipq *pq.MinIndexPQ // priority queue of vertices
}

func NewDijkstraSP(g EdgeWeightedDigraph, s int) *DijkstraSP {
	for _, edge := range g.Edges() {
		e := edge.(*DirectedEdge)
		if e.Weight() < 0 {
			panic("negative weight")
		}
	}

	n := g.V()
	distTo := make([]graphs.Distance, n)
	edgeTo := make([]*DirectedEdge, n)

	g.validateVertex(s)

	for i := 0; i < n; i++ {
		distTo[i] = graphs.DistanceInfinity
	}
	distTo[s] = graphs.DistanceZero

	ipq := pq.NewMinIndexPQ(n)
	ipq.Insert(s, distTo[s])

	sp := &DijkstraSP{g, s, distTo, edgeTo, ipq}

	// relax vertices in order of distance from s
	for !sp.ipq.IsEmpty() {
		v := sp.ipq.Delete()
		for _, edge := range g.Adj(v) {
			e := edge.(*DirectedEdge)
			sp.relax(e)
		}
	}

	return sp
}

// relax edge e and update pq if changed
func (sp *DijkstraSP) relax(e *DirectedEdge) {
	v, w := e.From(), e.To()
	distance := graphs.Distance(e.Weight())
	if sp.distTo[w] > sp.distTo[v] + distance {
		sp.distTo[w] = sp.distTo[v] + distance
		sp.edgeTo[w] = e

		if sp.ipq.Contains(w) {
			sp.ipq.Decrease(w, sp.distTo[w])
		} else {
			sp.ipq.Insert(w, sp.distTo[w])
		}
	}
}

// Returns the length of a shortest path from the source vertex s to vertex v
func (sp DijkstraSP) DistTo(v int) float64 {
	sp.graph.validateVertex(v)
	return float64(sp.distTo[v])
}

// Rtrurn true if there is a path from the source vertex to vertx v
func (sp DijkstraSP) HasPathTo(v int) bool {
	sp.graph.validateVertex(v)
	return sp.distTo[v] < graphs.DistanceInfinity
}

// Returns a shortest path from the source vertex to vertex v
func (sp DijkstraSP) PathTo(v int) fund.Iterator {
	sp.graph.validateVertex(v)
	stack := fund.NewStack()

	if sp.HasPathTo(v) {
		for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
			stack.Push(e)
		}
	}

	return stack.Iterator()
}

func (sp *DijkstraSP) Source() int {
	return sp.source
}