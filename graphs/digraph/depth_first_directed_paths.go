package digraph

import "github.com/youngzhu/algs4-go/fund"

type DepthFirstDirectedPaths struct {
	graph Digraph
	source int // source vertex
	marked []bool // marked[v]: is there an s-v path?
	edgeTo []int // edgeTo[v]: last edge on s-v path
}

// Computes a path between the source vertex (s) and every other vertex in graph g
func NewDepthFirstDirectedPaths(g Digraph, s int) DepthFirstDirectedPaths {
	g.validateVertex(s)

	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())

	path := DepthFirstDirectedPaths{g, s, marked, edgeTo}
	path.dfs(g, s)

	return path
}

// depth first search from v
func (p DepthFirstDirectedPaths) dfs(g Digraph, v int) {
	p.marked[v] = true

	for _, it := range g.Adj(v) {
		w := it.(int)
		if !p.marked[w] {
			p.edgeTo[w] = v
			p.dfs(g, w)
		}
	}
}

// Is there a path between the source vertex (s) and vertex (v)
func (p DepthFirstDirectedPaths) HasPathTo(v int) bool {
	p.graph.validateVertex(v)
	return p.marked[v]
}

// Returns a path between the source vertex (s) and vertex v
// or nil if no such path
func (p DepthFirstDirectedPaths) PathTo(v int) []int {
	if !p.HasPathTo(v) {
		return nil
	}

	stack := fund.NewStack()

	for x := v; x != p.source; x = p.edgeTo[x] {
		stack.Push(fund.Item(x))
	}
	stack.Push(fund.Item(p.source))

	path := make([]int, stack.Size())
	i := 0
	for !stack.IsEmpty() {
		path[i] = stack.Pop().(int)
		i++
	}

	return path
}