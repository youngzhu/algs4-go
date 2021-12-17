package digraph

type DirectedDFS struct {
	marked []bool // marked[v]: true if v is reachable from source (s)
	count int // number of vertices reachable from source (s)
}

// Computes the vertices in graph (g) that are reachable from
// the source vertex (s)
func NewDirectedDFS(g Digraph, s int) DirectedDFS {
	g.validateVertex(s)
	marked := make([]bool, g.V())

	search := DirectedDFS{marked, 0}
	search.dfs(g, s)

	return search
}

func NewDirectedDFSN(g Digraph, sources []int) DirectedDFS {
	marked := make([]bool, g.V())

	search := DirectedDFS{marked, 0}
	for _, s := range sources {
		if !search.marked[s] {
			search.dfs(g, s)
		}
	}

	return search
}

// depth first search from s
func (d DirectedDFS) dfs(g Digraph, s int) {
	d.count++
	d.marked[s] = true

	for _, it := range g.Adj(s) {
		w := it.(int)
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

// Is there a path between the source vertex (s) and vertex (v)
func (d DirectedDFS) Marked(v int) bool {
	return d.marked[v]
}

// Returns the number of vertices connected to the source vertex (s)
func (d DirectedDFS) Count() int {
	return d.count
}