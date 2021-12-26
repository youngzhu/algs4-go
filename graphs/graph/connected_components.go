package graph

// Another client of DFS, find the connected components of a graph.
type ConnectedComponents struct {
	graph  Graph
	marked []bool // makred[v]: has vertex v been marked?
	id     []int  // id[v]: id of connected component containing v
	size   []int  // size[id]: number of vertices in given component
	count  int    // number of connected components
}

func NewConnectedComponents(g Graph) ConnectedComponents {
	marked := make([]bool, g.V())
	id := make([]int, g.V())
	size := make([]int, g.V())

	cc := ConnectedComponents{g, marked, id, size, 0}

	for v := 0; v < g.V(); v++ {
		if !cc.marked[v] {
			cc.dfs(v)
			cc.count++
		}
	}

	return cc
}

func (cc ConnectedComponents) dfs(v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	cc.size[cc.count]++
	for _, ww := range cc.graph.Adj(v) {
		w := ww.(int)
		if !cc.marked[w] {
			cc.dfs(w)
		}
	}
}

// Returns the component id of the connected component containing vertex v
func (cc ConnectedComponents) Id(v int) int {
	cc.graph.validateVertex(v)
	return cc.id[v]
}

// Returns the number of vertices in the connected compoent containing vertex v
func (cc ConnectedComponents) Size(v int) int {
	cc.graph.validateVertex(v)
	return cc.size[v]
}

// Returns the number of connected components in the graph
func (cc ConnectedComponents) Count() int {
	return cc.count
}

// Returns true if vertices v and w are in the same connected component
func (cc ConnectedComponents) Connected(v, w int) bool {
	cc.graph.validateVertex(v)
	cc.graph.validateVertex(w)
	return cc.id[v] == cc.id[w]
}
