package sp

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/graphs/digraph"
)

// Single-source longest paths problem in edge-weighted DAGs.
// We can solve the problem by initializing the distTo[] values to negative
// infinity and switching the sense of the inequality in relax().
type AcyclicLP struct {
	graph  digraph.EdgeWeightedDigraph
	source int
	distTo []graphs.Distance       // distTo[v]: distance of longest s->v path
	edgeTo []*digraph.DirectedEdge // edgeTo[v]: last edge on longest s->v path
}

func NewAcyclicLP(g digraph.EdgeWeightedDigraph, s int) *AcyclicLP {
	n := g.V()
	distTo := make([]graphs.Distance, n)
	edgeTo := make([]*digraph.DirectedEdge, n)

	alp := &AcyclicLP{g, s, distTo, edgeTo}

	alp.validateVertex(s)

	for i := 0; i < n; i++ {
		distTo[i] = graphs.DistanceNegativeInfinity
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
			alp.relax(e)
		}
	}

	return alp
}

// relax edge e, update if find a longer path
func (alp *AcyclicLP) relax(e *digraph.DirectedEdge) {
	v, w := e.From(), e.To()
	distance := graphs.Distance(e.Weight())
	if alp.distTo[w] < alp.distTo[v]+distance {
		alp.distTo[w] = alp.distTo[v] + distance
		alp.edgeTo[w] = e
	}
}

// Returns the length of a longest path from the source vertex s to vertex v
func (alp AcyclicLP) DistTo(v int) float64 {
	alp.validateVertex(v)
	return float64(alp.distTo[v])
}

// HasPathTo return true if there is a path from the source vertex to vertx v
func (alp AcyclicLP) HasPathTo(v int) bool {
	alp.validateVertex(v)
	return alp.distTo[v] > graphs.DistanceNegativeInfinity
}

// Returns a longest path from the source vertex to vertex v
func (alp *AcyclicLP) PathTo(v int) fund.Iterator {
	alp.validateVertex(v)
	stack := fund.NewStack()

	if alp.HasPathTo(v) {
		for e := alp.edgeTo[v]; e != nil; e = alp.edgeTo[e.From()] {
			stack.Push(e)
		}
	}

	return stack.Iterator()
}

func (t *AcyclicLP) validateVertex(v int) {
	if v < 0 || v >= len(t.distTo) {
		panic("invalidate vertex")
	}
}
