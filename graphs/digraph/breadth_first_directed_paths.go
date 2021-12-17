package digraph

import (
	"github.com/youngzhu/algs4-go/graphs"
	"github.com/youngzhu/algs4-go/fund"
)

type BreadthFirstDirectedPaths struct {
	graph Digraph
	source int // source vertex
	marked []bool // marked[v]: is there an s-v path?
	edgeTo []int // edgeTo[v]: previous edge on shortest s-v path
	distTo []int // distTo[v]: number of edges on shortest s-v path
}

// Computes the shortest path between the source vertex (s) 
// and every other vertex in graph g
func NewBreadthFirstDirectedPaths(g Digraph, s int) BreadthFirstDirectedPaths {
	g.validateVertex(s)

	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	distTo := make([]int, g.V())

	for v := 0; v < g.V(); v++ {
		distTo[v] = graphs.InfinityDistance
	}

	path := BreadthFirstDirectedPaths{g, s, marked, edgeTo, distTo}
	path.bfs()

	return path
}

// breadth first search from s
func (p BreadthFirstDirectedPaths) bfs() {
	queue := fund.NewQueue()

	p.marked[p.source] = true
	p.distTo[p.source] = 0
	queue.Enqueue(fund.Item(p.source))

	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)

		for _, it := range p.graph.Adj(v) {
			w := it.(int)
			if !p.marked[w] {
				p.marked[w] = true
				p.edgeTo[w] = v
				p.distTo[w] = p.distTo[v] + 1
				queue.Enqueue(fund.Item(w))
			}
		}
	}
}

// Is there a path between the source vertex (s) and vertex (v)
func (p BreadthFirstDirectedPaths) HasPathTo(v int) bool {
	p.graph.validateVertex(v)
	return p.marked[v]
}

// Returns the number of edges in a shortest path between the source vertex s
// and vertex v
func (p BreadthFirstDirectedPaths) DistTo(v int) int {
	p.graph.validateVertex(v)
	return p.distTo[v]
}

// Returns a path between the source vertex (s) and vertex v
// or nil if no such path
func (p BreadthFirstDirectedPaths) PathTo(v int) []int {
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