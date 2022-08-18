package digraph

import (
	"fmt"

	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/testutil"
)

// An edge-weighted digraph, implemented using adjacency lists.
type EdgeWeightedDigraph struct {
	vertices int         // number of vertices in this digraph
	edges    int         // number of edges in this digraph
	adj      []*fund.Bag // adj[v]: adjaceny list for vertex v
	indegree []int       // indegree[v]: indegree of vertex v
}

// New an empty edge-weighted graph with n vertices and 0 edges
func NewEdgeWeightedDigraphN(n int) *EdgeWeightedDigraph {
	if n < 0 {
		panic("number of vertices must be non-negative")
	}

	indegree := make([]int, n)
	adj := make([]*fund.Bag, n)
	for v := 0; v < n; v++ {
		adj[v] = fund.NewBag()
	}

	return &EdgeWeightedDigraph{vertices: n, adj: adj, indegree: indegree}
}

func NewEdgeWeightedDigraphIn(in *testutil.In) *EdgeWeightedDigraph {
	if in == nil {
		panic("argument is nil")
	}

	vertices := in.ReadInt()
	if vertices < 0 {
		panic("Number of vertices must be non-negative")
	}
	indegree := make([]int, vertices)
	adj := make([]*fund.Bag, vertices)
	for v := 0; v < vertices; v++ {
		adj[v] = fund.NewBag()
	}

	edges := in.ReadInt()
	if edges < 0 {
		panic("Number of edges must be non-negative")
	}

	g := &EdgeWeightedDigraph{vertices, 0, adj, indegree}

	var v, w int
	for i := 0; i < edges; i++ {
		v, w = in.ReadInt(), in.ReadInt()
		g.validateVertex(v)
		g.validateVertex(w)
		weight := in.ReadFloat()
		e := NewDirectedEdge(v, w, weight)
		g.AddEdge(e)
	}

	return g
}

func (g *EdgeWeightedDigraph) AddEdge(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	g.validateVertex(v)
	g.validateVertex(w)
	g.adj[v].Add(e)
	g.indegree[w]++
	g.edges++
}

// Returns the edges incident on vertex v
func (g *EdgeWeightedDigraph) Adj(v int) fund.Iterator {
	g.validateVertex(v)
	return g.adj[v].Iterator()
}

// Returns the number of directed edges incident from vertex v
// This is known as the outdegree of vertex v
func (g *EdgeWeightedDigraph) Outdegree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

// Returns the number of directed edges incident to vertex v
// This is known as the indegree of vertex v
func (g *EdgeWeightedDigraph) Indegree(v int) int {
	g.validateVertex(v)
	return g.indegree[v]
}

// Returns all edges in this edge-weighted digraph
func (g *EdgeWeightedDigraph) Edges() fund.Iterator {
	bag := fund.NewBag()

	for v := 0; v < g.V(); v++ {
		for _, edge := range g.Adj(v) {
			e := edge.(*DirectedEdge)
			bag.Add(e)
		}
	}

	return bag.Iterator()
}

// Returns the number of vertices in this edge-weighted digraph
func (g *EdgeWeightedDigraph) V() int {
	return g.vertices
}

// Returns the number of edges in this edge-weighted digraph
func (g *EdgeWeightedDigraph) E() int {
	return g.edges
}

// Returns a string representation of this edge-weighted digraph
func (g *EdgeWeightedDigraph) String() string {
	s := fmt.Sprintf("vertices:%d, edges:%d\n", g.V(), g.E())
	for v := 0; v < g.V(); v++ {
		s += fmt.Sprintf("%d:", v)
		for _, e := range g.Adj(v) {
			s += fmt.Sprintf(" %v", e)
		}
		s += fmt.Sprintln()
	}

	return s
}

func (g *EdgeWeightedDigraph) validateVertex(v int) {
	if v < 0 || v >= g.V() {
		panic("invalidate vertex")
	}
}
