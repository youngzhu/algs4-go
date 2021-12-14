package graph

// Depth-first search is a classic recursive method for systematically examining
// each of the vertices and edges in a graph.
// To visit a vertex:
// - Mark it as having been visited
// - Visit (recursively) all the vertices that are adjacent to it and that have
// 	 not yet been marked

type DepthFirstSearch struct {
	marked []bool // marked[v]: is there an s-v path?
	count int // number of vertices connected to s
}

// Computes the vertices in graph (g) that are connected to
// the source vertex (s)
func NewDepthFirstSearch(g Graph, s int) DepthFirstSearch {
	g.validateVertex(s)
	marked := make([]bool, g.V())

	search := DepthFirstSearch{marked, 0}
	search.dfs(g, s)

	return search
}

// depth first search from s
func (d DepthFirstSearch) dfs(g Graph, s int) {
	d.count++
	d.marked[s] = true

	for _, w := range g.Adj(s) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

// Is there a path between the source vertex (s) and vertex (v)
func (d DepthFirstSearch) Marked(v int) bool {
	return d.marked[v]
}

// Returns the number of vertices connected to the source vertex (s)
func (d DepthFirstSearch) Count() int {
	return d.count
}