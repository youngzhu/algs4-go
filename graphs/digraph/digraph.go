package digraph

import (
	"fmt"

	"github.com/youngzhu/algs4-go/graphs/graph"
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/util"
)

// A directed graph (or digraph) is a set of vertices and a collection of directed
// edges that each connects an ordered pair of vertices. We say that a directed
// edge points from the first vertex in the pair and points to the second vertex in
// the pair. We use the names 0 through V-1 for the vertices in a V-vertex graph.
type IDigraph interface {
	graph.IGraph
}

// Use the adjacency-lists representation, where maintain a vertex-indexed array 
// of lists of the vertices connected by an edge to each vertex.
type Digraph struct {
	v int // number of vertices
	e int // number of edges
	adj []*fund.Bag // adj[v]: adjacency list for vertex v
	indegree []int // indegree[v]: indegree of vertex v
}

// New an empty digraph with v vertices
func NewDigraphN(v int) *Digraph {
	if v < 0 {
		panic("number of verties in a Digraph must be non-negative")
	}

	adj := make([]*fund.Bag, v)
	for i:=0; i < v; i++ {
		adj[i] = fund.NewBag()
	}

	indegree := make([]int, v)

	return &Digraph{v, 0, adj, indegree}
}

// New a digraph from the specified input stream.
// The format is the number of vertices V
// followed by the number of edges E
// followed by E pairs of vertices, with each entry separated by whitespace
func NewDigraph(in *util.In) *Digraph {
	v := in.ReadInt()

	g := NewDigraphN(v)

	e := in.ReadInt()
	if e < 0 {
		panic("number of edges in a Graph must be non-negative")
	}
	for i:=0; i < e; i++ {
		v1 := in.ReadInt()
		v2 := in.ReadInt()
		g.AddEdge(v1, v2)
	}

	return g
}

// Returns the number of vertices in this graph
func (g *Digraph) V() int {
	return g.v
}

// Returns the number of edges in this graph
func (g *Digraph) E() int {
	return g.e
}

// Adds the directed edge v1->v2 to this graph
func (g *Digraph) AddEdge(v1, v2 int) {
	g.validateVertex(v1)
	g.validateVertex(v2)
	g.e++
	g.adj[v1].Add(v2)
	g.indegree[v2]++
}

func (g *Digraph) Outdegree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

func (g *Digraph) Indegree(v int) int {
	g.validateVertex(v)
	return g.indegree[v]
}

func (g *Digraph) Adj(v int) fund.Iterator {
	g.validateVertex(v)
	return g.adj[v].Iterator()
}

// Returns the reverse of the digraph
func (g *Digraph) Reverse() *Digraph {
	reverse := NewDigraphN(g.V())

	for v := 0; v < g.V(); v++ {
		for _, it := range g.Adj(v) {
			w := it.(int)
			reverse.AddEdge(w, v)
		}
	}

	return reverse
}

// Returns a string representation of this graph
func (g *Digraph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		adjs := ""
		for _, w := range g.adj[i].Iterator() {
			adjs += fmt.Sprintf(" %d", w)
		}
		s += fmt.Sprintf("%d:%s\n", i, adjs)
	}
	return s
}

func (g *Digraph) validateVertex(v int) {
	if v < 0 || v >= g.v {
		panic("invalidate vertex")
	}
}