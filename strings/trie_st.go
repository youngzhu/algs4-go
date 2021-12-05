package strings

import (
	"github.com/youngzhu/algs4-go/fund"
)

// A string symbol table for extended ASCII strings, implemented using
// a 256-way trie

// R-way trie node
type node struct {
	value interface{}
	next  []*node
}

type TrieST struct {
	root *node // root of trie
	n    int   // number of keys in trie
}

func newNode() *node {
	next := make([]*node, R)
	return &node{next: next}
}

func NewTrieST() *TrieST {
	return &TrieST{}
}

// Returns the value associated with the given key
func (t *TrieST) Get(key string) interface{} {
	x := getNode(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.value
}
func getNode(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return getNode(x.next[c], key, d+1)
}

// Does this ST contain the given key?
func (t *TrieST) Contains(key string) bool {
	return t.Get(key) != nil
}

// Inserts the key-value pair into the symbol table,
// overwriting the old value with the new value if the key
// is already in the ST
func (t *TrieST) Put(key string, value interface{}) {
	t.root = t.put(t.root, key, value, 0)
}

func (t *TrieST) put(x *node, key string, value interface{}, d int) *node {
	if x == nil {
		x = newNode()
	}
	if d == len(key) {
		if x.value == nil {
			t.n++
		}
		x.value = value
		return x
	}
	c := key[d]
	x.next[c] = t.put(x.next[c], key, value, d+1)
	return x
}

// Returns the number of key-value pairs in the ST
func (t *TrieST) Size() int {
	return t.n
}

// Is this ST empty?
func (t *TrieST) IsEmpty() bool {
	return t.n == 0
}

// Returns all keys in the symbol table
func (t *TrieST) Keys() []string {
	return t.KeysWithPrefix("")
}

// Return all of the keys in the ST that start with prefix
func (t *TrieST) KeysWithPrefix(prefix string) []string {
	result := fund.NewQueue()
	x := getNode(t.root, prefix, 0)
	collect(x, prefix, result)

	return getSlice(*result)
}

func collect(x *node, prefix string, queue *fund.Queue) {
	if x == nil {
		return
	}
	if x.value != nil {
		queue.Enqueue(prefix)
	}

	prefixLen := len(prefix)
	for c := 0; c < R; c++ {
		prefix += string(rune(c))
		collect(x.next[c], prefix, queue)
		prefix = prefix[:prefixLen]
	}

	// for c := 0; c < R; c++ {
	// 	newPrefix := prefix + string(rune(c))
	// 	collect(x.next[c], newPrefix, queue)
	// }
}

// Returns all of the keys in the symbol table that match pattern,
// where the character '.' is interpreted as a wildcard character
func (t *TrieST) KeysThatMatch(pattern string) []string {
	result := fund.NewQueue()
	collectPattern(t.root, "", pattern, result)

	return getSlice(*result)
}

func collectPattern(x *node, prefix, pattern string, queue *fund.Queue) {
	if x == nil {
		return
	}

	d := len(prefix)
	if d == len(pattern) && x.value != nil {
		queue.Enqueue(prefix)
	}
	if d == len(pattern) {
		return
	}

	c := pattern[d]
	if c == '.' {
		for ch := 0; ch < R; ch++ {
			newPrefix := prefix + string(rune(ch))
			collectPattern(x.next[ch], newPrefix, pattern, queue)
		}
	} else {
		prefix += string(c)
		collectPattern(x.next[c], prefix, pattern, queue)
	}
}

// Returns the string in the ST that is the longest prefix of query
func (t *TrieST) LongestPrefixOf(query string) string {
	length := longestPrefixOf(t.root, query, 0, 0)

	return query[0:length]
}

// returns the length of the longest string key in the subtrie rooted
// at x that is a prefix of the query string, assuming the first d character
// match and we have already found a prefix match of given length
func longestPrefixOf(x *node, query string, d, length int) int {
	if x == nil {
		return length
	}
	if x.value != nil {
		length = d
	}
	if d == len(query) {
		return length
	}

	c := query[d]
	return longestPrefixOf(x.next[c], query, d+1, length)
}

// Removes the key from the ST if the key is present
func (t *TrieST) Delete(key string) {
	t.root = t.deleteNode(t.root, key, 0)
}

func (t *TrieST) deleteNode(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.value != nil {
			t.n--
			x.value = nil
		}
	} else {
		ch := key[d]
		x.next[ch] = t.deleteNode(x.next[ch], key, d+1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.value != nil {
		return x
	}
	for c := 0; c < R; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

func getSlice(queue fund.Queue) []string {
	slice := make([]string, queue.Size())
	for i, v := range queue.Iterator() {
		slice[i] = v.(string)
	}

	return slice
}
