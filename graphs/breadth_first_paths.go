package graphs

import "github.com/youngzhu/algs4-go/fund"

// Breadth-first search.
// Depth-first search finds some path from a source vertex s to a target vertex v.
// We are often interested in finding the shortest such path (one with a minimal 
// number of edges). Breadth-first search is a classic method based on this goal.
// To find a shortest path from s to v, we start at s and check for v among all 
// the vertices that we can reach by following one edge, then we check for v among
// all the vertices that we can reach from s by following two edges, and so forth.

// To implement this strategy, we maintain a queue of all vertices that have been
// marked but whose adjacency lists have not been checked. We put the source vertex
// on the queue, then perform the following steps until the queue is empty:
// 1. Remove the next vertex v from the queue
// 2. Put onto the queue all unmarked vertices that are adjacent to v and mark them

const infinity = 10000

type BreadthFirstPaths struct {
	graph Graph
	source int // source vertex
	marked []bool // marked[v]: is there an s-v path?
	edgeTo []int // edgeTo[v]: previous edge on shortest s-v path
	distTo []int // distTo[v]: number of edges on shortest s-v path
}

// Computes the shortest path between the source vertex (s) 
// and every other vertex in graph g
func NewBreadthFirstPaths(g Graph, s int) BreadthFirstPaths {
	g.validateVertex(s)

	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	distTo := make([]int, g.V())

	for v := 0; v < g.V(); v++ {
		distTo[v] = infinity
	}

	path := BreadthFirstPaths{g, s, marked, edgeTo, distTo}
	path.bfs(g, s)

	return path
}

// breadth first search from s
func (p BreadthFirstPaths) bfs(g Graph, s int) {
	queue := fund.NewQueue()

	p.marked[s] = true
	p.distTo[s] = 0
	queue.Enqueue(fund.Item(s))

	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)

		for _, w := range g.Adj(v) {
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
func (p BreadthFirstPaths) HasPathTo(v int) bool {
	p.graph.validateVertex(v)
	return p.marked[v]
}

// Returns the number of edges in a shortest path between the source vertex s
// and vertex v
func (p BreadthFirstPaths) DistTo(v int) int {
	p.graph.validateVertex(v)
	return p.distTo[v]
}

// Returns a path between the source vertex (s) and vertex v
// or nil if no such path
func (p BreadthFirstPaths) PathTo(v int) []int {
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