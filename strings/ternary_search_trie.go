package strings

import (
	"github.com/youngzhu/algs4-go/fund"
)

type TernarySearchTrie struct {
	root *tstNode // root of TST
	n    int      // size
}

type tstNode struct {
	char             byte        // character
	left, mid, right *tstNode    // left, middle, and right subtries
	value            interface{} // value associated with string
}

func NewTernarySearchTrie() *TernarySearchTrie {
	return &TernarySearchTrie{}
}

func (t *TernarySearchTrie) newNode(char byte) *tstNode {
	return &tstNode{char: char}
}

func (t *TernarySearchTrie) Get(key string) interface{} {
	if len(key) == 0 {
		panic("key must have length >= 1")
	}
	x := t.getNode(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.value
}

func (t *TernarySearchTrie) getNode(x *tstNode, key string, d int) *tstNode {
	if x == nil {
		return nil
	}
	c := key[d]
	if c < x.char {
		return t.getNode(x.left, key, d)
	} else if c > x.char {
		return t.getNode(x.right, key, d)
	} else if d < len(key)-1 {
		return t.getNode(x.mid, key, d+1)
	} else {
		return x
	}
}

func (t *TernarySearchTrie) Put(key string, value interface{}) {
	if !t.Contains(key) {
		t.n++
	} else if value == nil {
		t.n--
	}
	t.root = t.put(t.root, key, value, 0)
}
func (t *TernarySearchTrie) put(x *tstNode, key string, value interface{}, d int) *tstNode {
	c := key[d]
	if x == nil {
		x = t.newNode(c)
	}
	if c < x.char {
		x.left = t.put(x.left, key, value, d)
	} else if c > x.char {
		x.right = t.put(x.right, key, value, d)
	} else if d < len(key)-1 {
		x.mid = t.put(x.mid, key, value, d+1)
	} else {
		x.value = value
	}
	return x
}

func (t *TernarySearchTrie) LongestPrefixOf(prefix string) string {
	if len(prefix) == 0 {
		return ""
	}

	longestLen := 0
	x := t.root
	i := 0
	for x != nil && i < len(prefix) {
		c := prefix[i]
		if c < x.char {
			x = x.left
		} else if c > x.char {
			x = x.right
		} else {
			i++
			if x.value != nil {
				longestLen = i
			}
			x = x.mid
		}
	}
	return prefix[:longestLen]
}

func (t *TernarySearchTrie) Keys() []string {
	queue := fund.NewQueue()
	t.collect(t.root, "", queue)

	return getSlice(*queue)
}

func (t *TernarySearchTrie) collect(x *tstNode, prefix string, queue *fund.Queue) {
	if x == nil {
		return
	}
	t.collect(x.left, prefix, queue)
	if x.value != nil {
		queue.Enqueue(prefix + string(x.char))
	}
	t.collect(x.mid, prefix+string(x.char), queue)
	t.collect(x.right, prefix, queue)
}

func (t *TernarySearchTrie) KeysWithPrefix(prefix string) []string {
	queue := fund.NewQueue()
	x := t.getNode(t.root, prefix, 0)
	if x == nil {
		return getSlice(*queue)
	}
	if x.value != nil {
		queue.Enqueue(prefix)
	}
	t.collect(x.mid, prefix, queue)

	return getSlice(*queue)
}

func (t *TernarySearchTrie) KeysThatMatch(pattern string) []string {
	queue := fund.NewQueue()
	t.collectThatMatch(t.root, "", pattern, 0, queue)
	return getSlice(*queue)
}

func (t *TernarySearchTrie) collectThatMatch(x *tstNode, prefix, pattern string, i int, queue *fund.Queue) {
	if x == nil {
		return
	}

	c := pattern[i]
	if c == '.' || c < x.char {
		t.collectThatMatch(x.left, prefix, pattern, i, queue)
	}
	if c == '.' || c == x.char {
		if i == len(pattern)-1 && x.value != nil {
			queue.Enqueue(prefix + string(x.char))
		}
		if i < len(pattern)-1 {
			t.collectThatMatch(x.mid, prefix+string(x.char), pattern, i+1, queue)
		}
	}
	if c == '.' || c > x.char {
		t.collectThatMatch(x.right, prefix, pattern, i, queue)
	}
}

func (t *TernarySearchTrie) Contains(key string) bool {
	return t.Get(key) == nil
}

func (t *TernarySearchTrie) Size() int {
	return t.n
}
