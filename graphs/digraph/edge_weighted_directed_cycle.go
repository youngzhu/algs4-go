package digraph

// import "log"
import "github.com/youngzhu/algs4-go/fund"

// Does a given digraph have a directed cycle?
// Solves this problem using depth-first search

type EdgeWeightedDirectedCycle struct {
	marked []bool // marked[v]: has vertex v been marked?
	edgeTo []*DirectedEdge // edgeTo[v]: previous vertex on path to v
	onStack []bool // onStack[v]: is vertex on the stack?
	cycle *fund.Stack // directed cycle (or nil if no such cycle)
}

func NewEdgeWeightedDirectedCycle(g EdgeWeightedDigraph) *EdgeWeightedDirectedCycle {
	marked := make([]bool, g.V())
	edgeTo := make([]*DirectedEdge, g.V())
	onStack := make([]bool, g.V())
	dc := &EdgeWeightedDirectedCycle{marked, edgeTo, onStack, nil}

	for v := 0; v < g.V(); v++ {
		if !dc.marked[v] && dc.cycle == nil {
			dc.dfs(g, v)
		}
	}

	return dc
}

// run DFS and find a directed cycle 
// must use pointer (*), otherwise dc.cycle wouldn't change
func (dc *EdgeWeightedDirectedCycle) dfs(g EdgeWeightedDigraph, v int) {
	dc.onStack[v] = true
	dc.marked[v] = true

	for _, it := range g.Adj(v) {
		e := it.(*DirectedEdge)
		w := e.To()
		// short circuit if directed cycle found
		if dc.cycle != nil {
			return
		} 
		
		if !dc.marked[w] { // found new vertex, so recur
			dc.edgeTo[w] = e
			dc.dfs(g, w)
		} else if dc.onStack[w] { // trace back directed cycle
			cycle := fund.NewStack()

			x := e
			for x.From() != w {
				cycle.Push(x)
				x = dc.edgeTo[e.From()]
			}
			cycle.Push(x)
			// log.Println(w)
			dc.cycle = cycle
		}
	}

	dc.onStack[v] = false
}

// Does the digraph have a directed cycle?
func (dc *EdgeWeightedDirectedCycle) HasCycle() bool {
	return dc.cycle != nil 
}

func (dc *EdgeWeightedDirectedCycle) Cycle() fund.Iterator {
	if dc.HasCycle() {
		return dc.cycle.Iterator()
	}
	
	return nil
}