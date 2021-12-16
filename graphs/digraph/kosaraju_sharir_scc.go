package digraph

// Strong connectivity is an equivalence relation on the set of vertices:
// Reflexive: Every vertex v is strongly connected to itself
// Symmetric: If v is strongly connected to w, then w is strongly connected to v
// Transtive: If v is strongly connected to w and w is strongly connected to x,
// then v is also strongly connected to x

// implements as follows:
// 1. Given a digraph G, use DepthFirstOrder to compute the reverse postorder of its reverse
// 2. Run standard DFS on G, but consider the unmarked vertices in the order 
// just computed instead of the standard numerical order
// 3. All vertices reached on a call to the recursive dfs() from New method

type KosarajuSharirSCC struct {
	digraph Digraph
	marked []bool // marked[v]: has vertex v been visited?
	id []int // id[v]: id of strong component containing v
	count int // number of strongly-connected components
}

func NewKosarajuSharirSCC(g Digraph) KosarajuSharirSCC {
	// compute reverse postorder of reverse graph
	dfo := NewDepthFirstOrder(g.Reverse())

	marked := make([]bool, g.V())
	id := make([]int, g.V())

	scc := &KosarajuSharirSCC{digraph: g, marked: marked, id: id}

	// run DFS on g, using reverse postorder to guide calulation
	for _, v := range dfo.ReversePostorder() {
		i := v.(int)
		if !scc.marked[i] {
			scc.dfs(i)
			scc.count++
		}
	}

	return *scc
}

func (scc *KosarajuSharirSCC) dfs(v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count
	for _, it := range scc.digraph.Adj(v) {
		w := it.(int)
		if !scc.marked[w] {
			scc.dfs(w)
		}
	}
}

// Returns the number of strong components
func (scc KosarajuSharirSCC) Count() int {
	return scc.count
}

// Are vertices v and w in the same strong component?
func (scc KosarajuSharirSCC) StronglyConnected(v, w int) bool {
	scc.digraph.validateVertex(v)
	scc.digraph.validateVertex(w)
	return scc.id[v] == scc.id[w]
}

// Returns the component id of the strong component containing vertex v
func (scc KosarajuSharirSCC) Id(v int) int {
	scc.digraph.validateVertex(v)
	return scc.id[v]
}