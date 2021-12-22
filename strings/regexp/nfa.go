package regexp

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/graphs/digraph"
)

type NFA struct {
	graph digraph.Digraph // digraph of epsilon transition
	regExp string // regular expression
	m int // number of characters in regular expression
}

func NewNFA(regExp string) *NFA {
	m := len(regExp)
	ops := fund.NewStack()
	graph := digraph.NewDigraphN(m+1)
	for i := 0; i < m; i++ {
		lp := i
		if regExp[i] == '(' || regExp[i] == '|' {
			ops.Push(i)
		} else if regExp[i] == ')' {
			or := ops.Pop().(int)

			// 2-way or operator
			if regExp[or] == '|' {
				lp = ops.Pop().(int)
				graph.AddEdge(lp, or+1)
				graph.AddEdge(or, i)
			} else {
				lp = or
			}
		}

		// closure operator (uses 1-character lookahead)
		if i < m-1 && regExp[i+1] == '*' {
			graph.AddEdge(lp, i+1)
			graph.AddEdge(i+1, lp)
		}
		if regExp[i] == '(' || regExp[i] == '*' || regExp[i] == ')' {
			graph.AddEdge(i, i+1)
		}
	}

	if ops.Size() != 0 {
		panic("invalid regular expression")
	}

	return &NFA{*graph, regExp, m}
}

// Returns true if the text is matched by the regular expression
func (p *NFA) Recognizes(txt string) bool {
	dfs := digraph.NewDirectedDFS(p.graph, 0)
	pc := fund.NewBag()
	for v := 0; v < p.graph.V(); v++ {
		if dfs.Marked(v) {
			pc.Add(v)
		}
	}

	// Compute possible NFA states for txt[i+1]
	for i := 0; i < len(txt); i++ {
		cur := txt[i]
		if cur == '*' || cur == '|' || cur == '(' || cur == ')' {
			panic("the text contains the metacharacter")
		}

		match := fund.NewBag()
		for _, val := range pc.Iterator() {
			v := val.(int)
			if v == p.m {
				continue
			}
			if p.regExp[v] == cur || p.regExp[v] == '.' {
				match.Add(v+1)
			}
		}
		sources := make([]int, match.Size())
		for i, v := range match.Iterator() {
			sources[i] = v.(int)
		}
		dfs = digraph.NewDirectedDFSN(p.graph, sources)
		pc = fund.NewBag()
		for v := 0; v < p.graph.V(); v++ {
			if dfs.Marked(v) {
				pc.Add(v)
			}
		}

		// optimization if no states reachable
		if pc.Size() == 0 {
			return false
		}
	}

	// check for accept state
	for _, v := range pc.Iterator() {
		if v == p.m {
			return true
		}
	}
	return false
}