package graphs

import "fmt"

type Edge struct {
	v, w int
	weight float64
}

func NewEdge(v, w int, weight float64) Edge {
	if v < 0 || w < 0 {
		panic("vertex index must be a non-negative integer")
	}

	return Edge{v, w, weight}
}

/* The Either() and Other methods are usefule for accessing the edge's vertices */

// Return either endpoint of this edge
func (e Edge) Either() int {
	return e.v
}

// Returns the endpoint of this edge that is different from the given vertex
func (e Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	} else if vertex == e.w {
		return e.v
	} else {
		panic("Illegal endpoint")
	}
}

func (e Edge) Weight() float64 {
	return e.weight
}

// Returns a string representation of this edge
func (e Edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.v, e.w, e.weight)
}