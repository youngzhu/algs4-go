package digraph

// Topological sort.
// Given a digraph, put the vertices in order such that all its directed edges
// point from a vertex earlier in the order to a vertex later in the order (or
// report that doing so is not possible).
// Remarkably, a reverse postorder in a DAG provides a topological order.

// A digraph has a topological order if and only if it is a DAG.
// Reverse postorder in a DAG is a topological sort.
// With depth-first search, we can topologically sort a DAG in time proportional to V+E.

type Topological struct {
	order []int // topological order
	rank []int // rank[v]: rank of vertex v in order
}

func NewTopological(g IDigraph) Topological {
	t := Topological{}

	finder := NewDirectedCycle(g)
	if !finder.HasCycle() {
		dfo := NewDepthFirstOrder(g)
		order := dfo.ReversePostorder()
		rank := make([]int, g.V())

		for i, v := range order {
			rank[v.(int)] = i
		}

		orderInt := make([]int, len(order))
		for i, v := range order {
			orderInt[i] = v.(int)
		}

		t.order = orderInt
		t.rank = rank
	}

	return t
}

// Returns a topological order if the digraph has a topologial order
func (t Topological) Order() []int {
	return t.order
}

// Does the digraph have a topological order
func (t Topological) HasOrder() bool {
	return t.order != nil
}

// Does the digraph is a DAG
// equivalently, does have a topological order
func (t Topological) IsDAG() bool {
	return t.HasOrder()
}

// The rank of vertex v in the topological order
// return -1, if the digraph is not a DAG
func (t Topological) Rank(v int) int {
	t.validateVertex(v)
	if t.HasOrder() {
		return t.rank[v]
	}
	return -1
}

func (t Topological) validateVertex(v int) {
	if v < 0 || v >= len(t.rank) {
		panic("invalidate vertex")
	}
}