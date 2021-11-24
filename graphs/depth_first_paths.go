package graphs

import "github.com/youngzhu/algs4-go/fund"

// We remember the edge v-w that takes us to each vertex w for the first time 
// by setting edge[w] to v. In other words, v-w is the last edge on the known
// path from s to w. The result of the search is a tree rooted at the source;
// edgeTo[] is a parent-link representation of that tree.

type DepthFirstPaths struct {
	graph Graph
	source int // source vertex
	marked []bool // marked[v]: is there an s-v path?
	edgeTo []int // edgeTo[v]: last edge on s-v path
}

// Computes a path between the source vertex (s) and every other vertex in graph g
func NewDepthFirstPaths(g Graph, s int) DepthFirstPaths {
	g.validateVertex(s)

	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())

	path := DepthFirstPaths{g, s, marked, edgeTo}
	path.dfs(g, s)

	return path
}

// depth first search from v
func (p DepthFirstPaths) dfs(g Graph, v int) {
	p.marked[v] = true

	for _, w := range g.Adj(v) {
		if !p.marked[w] {
			p.edgeTo[w] = v
			p.dfs(g, w)
		}
	}
}

// Is there a path between the source vertex (s) and vertex (v)
func (p DepthFirstPaths) HasPathTo(v int) bool {
	p.graph.validateVertex(v)
	return p.marked[v]
}

// Returns a path between the source vertex (s) and vertex v
// or nil if no such path
func (p DepthFirstPaths) PathTo(v int) []int {
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