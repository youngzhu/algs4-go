package digraph

import (
	"strings"

	"github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/util"
)

type SymbolDigraph struct {
	st searching.RedBlackBST // string -> index
	keys []string // index -> string
	digraph Digraph // the underlying digraph
}

// New a Symbol Digraph from a file using the specified delimiter.
// Each line in the file contains the name of a vertex,
// followed by a list of the names of the vertices adjacent to that vertex,
// separated by the delimiter.
func NewSymbolDigraph(filepath, delimiter string) SymbolDigraph {
	st := searching.NewRedBlackBST()

	// first pass builds the index by reading strings to associate
	// distinct strings with an index
	in := util.NewInReadLines(filepath)
	for in.HasNext() {
		slice := strings.Split(in.ReadLine(), delimiter)
		for _, v := range slice {
			key := searching.StringKey(v)
			if !st.Contains(key) {
				st.Put(key, st.Size())
			}
		}
	}

	// inverted index to get string keys in an array
	keys := make([]string, st.Size())
	for _, k := range st.Keys() {
		i := st.Get(k).(int)
		keys[i] = string(k.(searching.StringKey))
	}

	// second pass builds the digraph by connecting first vertex 
	// on each line to all others
	digraph := NewDigraphN(st.Size())
	in = util.NewIn(filepath)
	for in.HasNext() {
		slice := strings.Split(in.ReadLine(), delimiter)
		key := searching.StringKey(slice[0])
		v := st.Get(key).(int) // first vertex
		var w int
		for _, str := range slice[1:] {
			key = searching.StringKey(str)
			w = st.Get(key).(int)
			digraph.AddEdge(v, w)
		}
	}

	return SymbolDigraph{*st, keys, *digraph}
}

// Does the digraph contain the vertex named s
func (sg SymbolDigraph) Contains(s string) bool {
	key := searching.StringKey(s)
	return sg.st.Contains(key)
}

// Returns the integer associated with the vertex named s
func (sg SymbolDigraph) Index(s string) int {
	key := searching.StringKey(s)
	return sg.st.Get(key).(int)
}

// Returns the name of the vertex associated with the integer v
func (sg SymbolDigraph) Name(v int) string {
	sg.digraph.validateVertex(v)
	return sg.keys[v]
}

// Returns the digraph associated with the symbol digraph
func (sg SymbolDigraph) Digraph() Digraph {
	return sg.digraph
}