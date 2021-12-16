package digraph

import "github.com/youngzhu/algs4-go/fund"

// Depth-first orders.
// DFS search visits each vertex exactly once. Three vertex orderings are of
// interest in typical applications:
// 1. Preorder: Put the vertex on a queue before the recursive calls
// 2. Postorder: Put the vertex on a queue after the recursive calls
// 3. Reverse postorder: Put the vertex on a stack after the recursive calls

type DepthFirstOrder struct {
	marked []bool // marked[v]: has v marked
	pre []int // pre[v]: preorder number of v
	post []int // post[v]: postorder number of v
	preorder *fund.Queue // vertices in preorder
	postorder *fund.Queue // vertices in postorder
	preCounter int // counter for preorder numbering
	postCounter int // counter for postorder numbering
}

func NewDepthFirstOrder(g IDigraph) DepthFirstOrder {
	n := g.V()
	marked := make([]bool, n)
	pre := make([]int, n)
	post := make([]int, n)
	preorder := fund.NewQueue()
	postorder := fund.NewQueue()

	dfo := &DepthFirstOrder{
		marked: marked, 
		pre: pre, 
		post: post, 
		preorder: preorder, 
		postorder: postorder}

	for v := 0; v < n; v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}

	return *dfo
}

func NewDepthFirstOrderWeighted(g EdgeWeightedDigraph) DepthFirstOrder {
	n := g.V()
	marked := make([]bool, n)
	pre := make([]int, n)
	post := make([]int, n)
	preorder := fund.NewQueue()
	postorder := fund.NewQueue()

	dfo := &DepthFirstOrder{
		marked: marked, 
		pre: pre, 
		post: post, 
		preorder: preorder, 
		postorder: postorder}

	for v := 0; v < n; v++ {
		if !dfo.marked[v] {
			dfo.dfsWeighted(g, v)
		}
	}

	return *dfo
}

func (dfo *DepthFirstOrder) dfs(g IDigraph, v int) {
	dfo.marked[v] = true
	dfo.pre[v] = dfo.preCounter
	dfo.preCounter++
	dfo.preorder.Enqueue(v)

	for _, it := range g.Adj(v) {
		w := it.(int)
		if !dfo.marked[w] {
			dfo.dfs(g, w)
		}
	}

	dfo.postorder.Enqueue(v)
	dfo.post[v] = dfo.postCounter
	dfo.postCounter++
}

func (dfo *DepthFirstOrder) dfsWeighted(g EdgeWeightedDigraph, v int) {
	dfo.marked[v] = true
	dfo.pre[v] = dfo.preCounter
	dfo.preCounter++
	dfo.preorder.Enqueue(v)

	for _, it := range g.Adj(v) {
		w := it.(int)
		if !dfo.marked[w] {
			dfo.dfsWeighted(g, w)
		}
	}

	dfo.postorder.Enqueue(v)
	dfo.post[v] = dfo.postCounter
	dfo.postCounter++
}

// Returns the preorder number of vertex v
func (dfo DepthFirstOrder) Pre(v int) int {
	dfo.validateVertex(v)
	return dfo.pre[v]
}

// Returns the postorder number of vertex v
func (dfo DepthFirstOrder) Post(v int) int {
	dfo.validateVertex(v)
	return dfo.post[v]
}

// Return the vertices in postorder
func (dfo DepthFirstOrder) Postorder() fund.Iterator {
	return dfo.postorder.Iterator()
}

// Return the vertices in preorder
func (dfo DepthFirstOrder) Preorder() fund.Iterator {
	return dfo.preorder.Iterator()
}

// Return the vertices in reverse postorder
func (dfo DepthFirstOrder) ReversePostorder() fund.Iterator {
	reverse := fund.NewStack()
	for _, v := range dfo.Postorder() {
		reverse.Push(v)
	}
	return reverse.Iterator()
}

func (t DepthFirstOrder) validateVertex(v int) {
	if v < 0 || v >= len(t.marked) {
		panic("invalidate vertex")
	}
}