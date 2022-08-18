package mst

import (
	"fmt"

	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/testutil"
)

type EdgeWeightedGraph struct {
	vertices int // number of vertices
	edges    int // number of edges
	adj      []*fund.Bag
}

func NewEdgeWeightedGraphIn(in *testutil.In) *EdgeWeightedGraph {
	if in == nil {
		panic("argument is nil")
	}

	vertices := in.ReadInt()
	if vertices < 0 {
		panic("Number of vertices must be non-negative")
	}
	adj := make([]*fund.Bag, vertices)
	for v := 0; v < vertices; v++ {
		adj[v] = fund.NewBag()
	}

	edges := in.ReadInt()
	if edges < 0 {
		panic("Number of edges must be non-negative")
	}

	// AddEdge() will update g.edges
	g := &EdgeWeightedGraph{vertices, 0, adj}

	var v, w int
	for i := 0; i < edges; i++ {
		v, w = in.ReadInt(), in.ReadInt()
		g.validateVertex(v)
		g.validateVertex(w)
		weight := in.ReadFloat()
		e := NewEdge(v, w, weight)
		g.AddEdge(e)
	}

	return g
}

// Returns the number of vertices in this edge-weighted graph
func (g *EdgeWeightedGraph) V() int {
	return g.vertices
}

// Returns the number of edges in this edge-weighted graph
func (g *EdgeWeightedGraph) E() int {
	return g.edges
}

// Add the undirected edge to this edge-weighted graph
func (g *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	g.validateVertex(v)
	g.validateVertex(w)
	g.adj[v].Add(e)
	g.adj[w].Add(e)
	g.edges++
}

// Returns the edges incident on vertex v
func (g *EdgeWeightedGraph) Adj(v int) fund.Iterator {
	g.validateVertex(v)
	return g.adj[v].Iterator()
}

// Returns the degree of vertex v
func (g *EdgeWeightedGraph) Degree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

// Returns all edges in this edge-weighted graph
func (g *EdgeWeightedGraph) Edges() fund.Iterator {
	bag := fund.NewBag()
	for v := 0; v < g.V(); v++ {
		selfLoops := 0
		for _, edge := range g.Adj(v) {
			e := edge.(*Edge)
			if e.Other(v) > v {
				bag.Add(e)
			} else if e.Other(v) == v {
				// add only one copy of each slef loop (self loops will be consecutive)
				if selfLoops%2 == 0 {
					bag.Add(e)
				}
				selfLoops++
			}
		}

	}
	return bag.Iterator()
}

// Returns a string representation of this edge-weighted graph
func (g *EdgeWeightedGraph) String() string {
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

func (g *EdgeWeightedGraph) validateVertex(v int) {
	if v < 0 || v >= g.V() {
		panic("invalidate vertex")
	}
}
