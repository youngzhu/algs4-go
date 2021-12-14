package sp

import (
	"fmt"
)

type DirectedEdge struct {
	v, w int
	weight float64
}

func NewDirectedEdge(v, w int, weight float64) *DirectedEdge {
	if v < 0 || w < 0 {
		panic("vertex index must be a non-negative integer")
	}

	return &DirectedEdge{v, w, weight}
}

// Returns the tail vertex of the directed edge
func (e *DirectedEdge) From() int {
	return e.v
}

// Returns the head vertex of the directed edge
func (e *DirectedEdge) To() int {
	return e.w
}

func (e *DirectedEdge) Weight() float64 {
	return e.weight
}

func (e *DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %5.2f", e.v, e.w, e.weight)
}