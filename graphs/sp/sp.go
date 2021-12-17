package sp

import "github.com/youngzhu/algs4-go/graphs"

// Shortest paths.
// An edge-weighted digraph is a digraph where we associate weights or costs 
// with each edge. A shortest path from vertex s to vertex t is a directed
// path from s to t with the property that no other such path has a lower weight.

type ShortestPaths interface {
	graphs.Paths
	
	DistTo(v int) float64
}