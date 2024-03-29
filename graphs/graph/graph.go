package graph

import (
	"fmt"

	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/testutil"
)

// A graph is a set of vertices and a collection of edges that each connect a
// pair of vertices.

// IGraph Undirected graph
type IGraph interface {
	V() int
	E() int
	AddEdge(v, w int)
	Adj(v int) fund.Iterator
}

// Graph
// implemented using an array of set.
// Parallel edges and self-loops allowed
type Graph struct {
	v   int         // number of vertices
	e   int         // number of edges
	adj []*fund.Bag //
}

// NewGraph
// New a graph from the specified input stream.
// The format is the number of vertices V
// followed by the number of edges E
// followed by E pairs of vertices, with each entry separated by whitespace
func NewGraph(in *testutil.In) *Graph {
	v := in.ReadInt()
	if v < 0 {
		panic("number of verties in a Graph must be non-negative")
	}

	adj := make([]*fund.Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = fund.NewBag()
	}

	g := &Graph{v, 0, adj}

	e := in.ReadInt()
	if e < 0 {
		panic("number of edges in a Graph must be non-negative")
	}
	for i := 0; i < e; i++ {
		v1 := in.ReadInt()
		v2 := in.ReadInt()
		g.AddEdge(v1, v2)
	}

	return g
}

// NewGraphN
// New an empty graph with v vertices and 0 edges
func NewGraphN(v int) *Graph {
	if v < 0 {
		panic("number of vertices in a Graph must be non-negative")
	}

	adj := make([]*fund.Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = fund.NewBag()
	}

	return &Graph{v, 0, adj}
}

// exercise 4.1.3

// NewGraphCopy
// new a graph that is a deep copy of g
func NewGraphCopy(g Graph) *Graph {
	v, e := g.V(), g.E()
	if v < 0 {
		panic("number of vertices in a Graph must be non-negative")
	}

	adj := make([]*fund.Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = fund.NewBag()
	}

	for i := 0; i < v; i++ {
		// reverse so that adjacency list is in same order
		// as original
		reverse := fund.NewStack()
		for _, w := range g.Adj(i) {
			reverse.Push(w)
		}
		for !reverse.IsEmpty() {
			adj[i].Add(reverse.Pop())
		}
	}

	return &Graph{v, e, adj}
}

// V
// Returns the number of vertices in this graph
func (g *Graph) V() int {
	return g.v
}

// E
// Returns the number of edges in this graph
func (g *Graph) E() int {
	return g.e
}

// Adds the undirected edge v1-v2 to this graph
func (g *Graph) AddEdge(v1, v2 int) {
	g.validateVertex(v1)
	g.validateVertex(v2)
	g.e++
	g.adj[v1].Add(v2)
	g.adj[v2].Add(v1)
}

// Returns the degree of the vertex v
func (g *Graph) Degree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

func (g *Graph) Adj(v int) fund.Iterator {
	g.validateVertex(v)

	return g.adj[v].Iterator()
}

// Returns a string representation of this graph
func (g *Graph) String() string {
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

func (g *Graph) validateVertex(v int) {
	if v < 0 || v >= g.v {
		panic("invalidate vertex")
	}
}

// Maximum degree
func (g *Graph) MaxDegree() int {
	max := 0

	for v := 0; v < g.v; v++ {
		if g.Degree(v) > max {
			max = g.Degree(v)
		}
	}

	return max
}

// Average degree
func (g *Graph) AvgDegree() int {
	// each edge incident on two vertices
	return 2 * g.e / g.v
}

// Number of self-loops
func (g *Graph) NumberOfSelfLoops() int {
	count := 0

	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v].Iterator() {
			if v == w {
				count++
			}
		}
	}

	return count / 2 // self loop appears in adjacency list twice
}

// exercise 4.1.4

// HasEdge ...
func (g *Graph) HasEdge(v, w int) bool {
	for _, ww := range g.Adj(v) {
		if ww == w {
			return true
		}
	}
	return false
}
