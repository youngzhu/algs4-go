package graphs

import (
	"strings"

	"github.com/youngzhu/algs4-go/searching"
	"github.com/youngzhu/algs4-go/util"
)

// Symbol graphs.
// Typical applications involve processing graphs using strings, not integer 
// indices, to difine and refer to vertices. To accommodate such applications,
// we define an input format with the following properties:
// 1. Vertex names are strings.
// 2. A specified delimiter separates vertex names (to allow for the possibility
//    of spaces in names).
// 3. Each line represents a set of edges, connecting the first vertex name on 
//    the line to each of the other vertices named on the line

// Three data structures:
// 1. A symbol table st with string keys (vertex names) and int values (indices)
// 2. An array keys[] that serves as an inverted index, giving the vertex name
//    associated with each integer index
// 3. A graph built using the indices to refer to vertices
type SymbolGraph struct {
	st searching.RedBlackBST // string -> index
	keys []string // index -> string
	graph Graph // the underlying graph
}

// New a Symbol Graph from a file using the specified delimiter.
// Each line in the file contains the name of a vertex,
// followed by a list of the names of the vertices adjacent to that vertex,
// separated by the delimiter.
func NewSymbolGraph(filepath, delimiter string) SymbolGraph {
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

	// second pass builds the graph by connecting first vertex 
	// on each line to all others
	graph := NewGraphN(st.Size())
	in = util.NewIn(filepath)
	for in.HasNext() {
		slice := strings.Split(in.ReadLine(), delimiter)
		key := searching.StringKey(slice[0])
		v := st.Get(key).(int) // first vertex
		var w int
		for _, str := range slice[1:] {
			key = searching.StringKey(str)
			w = st.Get(key).(int)
			graph.AddEdge(v, w)
		}
	}

	return SymbolGraph{*st, keys, *graph}
}

// Does the graph contain the vertex named s
func (sg SymbolGraph) Contains(s string) bool {
	key := searching.StringKey(s)
	return sg.st.Contains(key)
}

// Returns the integer associated with the vertex named s
func (sg SymbolGraph) Index(s string) int {
	key := searching.StringKey(s)
	return sg.st.Get(key).(int)
}

// Returns the name of the vertex associated with the integer v
func (sg SymbolGraph) Name(v int) string {
	sg.graph.validateVertex(v)
	return sg.keys[v]
}

// Returns the graph associated with the symbol graph
func (sg SymbolGraph) Graph() Graph {
	return sg.graph
}