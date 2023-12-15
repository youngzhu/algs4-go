package searching

import "github.com/youngzhu/algs4-go/fund"

// We examine a symbol-table implementation that combines the flexibility of
// insertion in linked lists with the efficiency of search in an ordered array.
// Specifically, using two links per node leads to an efficient symbol-table
// implementation based on the binary search tree data structure, which qualifies
// as one of the most fundamental algorithms in computer science.

// Binary Search Tree (BST)
// BST is a binary tree where each node has a Comparable key (and an associated value)
// and satisfies the restriction that the key in any node is larger than the keys
// in all nodes in that node's left subtree and smaller than the keys in all nodes
// in that node's right subtree.

// BST implements the ordered symbol-table API using a binary search tree. We
// define a type to define nodes in BST. Each node contains a key, a value,
// a left link, a right link, and a node count. The left link points to a BST
// for items with smaller keys, and the right link points to a BST for items with
// larger keys. The variable N (size) gives the node count in the subtree rooted
// at the node. This field facilitates the implementation of various ordered
// symobl-table operations, as you will see.

// Search.
// A recursive algorithm to search for a key in a BST follows immediately from the
// recursive structure: If the tree is empty, we have a search miss; if the search
// key is equal to the key at the root, we have a serarch hit. Otherwise, we search
// (recursively) in the appropriate subtree. The recursive get() method implements
// this algorithm directly. It takes a node (root of a subtree) as first argument
// and a key as second argument, starting with the root of the tree and the search key.

// Insert.
// Insert is not much more difficult to implement htan search. Indeed, a search for
// a key not in the tree ends a null link, and all that we need to do is replace that
// link with a new node containing the key. The recursive put() method accomplishes
// this task using logic similar to that we used for recursive search: If the tree is
// empty, we return a new node containing the key and value; if the search key is less
// than the key at the root, we set the left link to the result of inserting the key
// into the left subtree; otherwise, we set the right link to the result of inserting
// the key into the right subtree.

type BST struct {
	root *Node // root of BST
}

// must be `Node`, not `node`
// otherwise got error: cannot use b.root (type *node) as type *node in argument to get
type Node struct {
	key         OSTKey  // sorted by key
	value       STValue // associated data
	left, right *Node   // left and right subtrees
	size        int     // number of nodes in subtree
}

func NewBST() *BST {
	return &BST{}
}

func newNode(key OSTKey, value STValue) *Node {
	return &Node{key: key, value: value, size: 1}
}

// Returns the value associated with the given key
func (b *BST) Get(key OSTKey) STValue {
	return get(b.root, key)
}

func get(x *Node, key OSTKey) STValue {
	if key == nil {
		panic("calls get() with a nil key")
	}
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		return get(x.left, key)
	} else if cmp > 0 {
		return get(x.right, key)
	} else {
		return x.value
	}
}

// Inserts the specified key-value pair into the symbol table, overwriting the
// old value with the new value if the symbol table already contains the specified key.
//
// Deletes the specified key (and its associated value) from the symbol table if
// the specified value is nil
func (b *BST) Put(key OSTKey, value STValue) {
	if key == nil {
		panic("calls put() with a nil key")
	}
	if value == nil {
		b.Delete(key)
		return
	}
	b.root = put(b.root, key, value)
}

func put(x *Node, key OSTKey, value STValue) *Node {
	if x == nil {
		return newNode(key, value)
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = put(x.left, key, value)
	} else if cmp > 0 {
		x.right = put(x.right, key, value)
	} else {
		x.value = value
	}

	x.size = 1 + size(x.left) + size(x.right)

	return x
}

// return number of key-value pairs in BST rooted at x
func size(x *Node) int {
	if x == nil {
		return 0
	} else {
		return x.size
	}
}

// Reomoves the specified key and its associated value from the symbol table
// (if the key is in this symbol table)
func (b *BST) Delete(key OSTKey) {
	if key == nil {
		panic("calls Delete() with a nil key")
	}
	b.root = delete(b.root, key)
}

func delete(x *Node, key OSTKey) *Node {
	if x == nil {
		return nil
	}

	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = delete(x.left, key)
	} else if cmp > 0 {
		x.right = delete(x.right, key)
	} else {
		if x.left == nil {
			return x.right
		}
		if x.right == nil {
			return x.left
		}
		t := x
		x = min(t.right)
		x.right = deleteMin(t.right)
		x.left = t.left
	}

	x.size = 1 + size(x.left) + size(x.right)

	return x
}

func deleteMin(x *Node) *Node {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.size = 1 + size(x.left) + size(x.right)
	return x
}

func min(x *Node) *Node {
	if x.left == nil {
		return x
	} else {
		return min(x.left)
	}
}

// Returns all keys in the symbol table
func (b *BST) Keys() []OSTKey {
	if b.IsEmpty() {
		panic("The BST is empty")
	}

	return b.rangeKeys(b.Min(), b.Max())
}

// Returns all keys in the symbol table in the given range
func (b *BST) rangeKeys(lo, hi OSTKey) []OSTKey {
	if lo == nil {
		panic("first argument to rangeKeys() is nil")
	}
	if hi == nil {
		panic("second argument to rangeKeys() is nil")
	}

	queue := fund.NewQueue()
	keys(b.root, queue, lo, hi)

	keySliece := make([]OSTKey, queue.Size())
	i := 0
	for _, v := range queue.Iterator() {
		keySliece[i] = v.(OSTKey)
		i++
	}

	return keySliece
}

func keys(x *Node, queue *fund.Queue, lo, hi OSTKey) {
	if x == nil {
		return
	}

	cmpLo := lo.CompareTo(x.key)
	cmpHi := hi.CompareTo(x.key)
	if cmpLo < 0 {
		keys(x.left, queue, lo, hi)
	}
	if cmpLo <= 0 && cmpHi >= 0 {
		queue.Enqueue(x.key)
	}
	if cmpHi > 0 {
		keys(x.right, queue, lo, hi)
	}
}

// Returns the smallest key in the BST
func (b *BST) Min() OSTKey {
	if b.IsEmpty() {
		panic("The BST is empty")
	}

	return min(b.root).key
}

// Returns the largest key in the BST
func (b *BST) Max() OSTKey {
	if b.IsEmpty() {
		panic("The BST is empty")
	}
	return max(b.root).key
}

func max(x *Node) *Node {
	if x.right == nil {
		return x
	} else {
		return max(x.right)
	}
}

// Returns true if this symbol table is empty
func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

// Returns the nubmer of key-value pairs in this symbol table
func (b *BST) Size() int {
	return size(b.root)
}

// Does this BST contains the given key?
func (b *BST) Contains(key OSTKey) bool {
	if key == nil {
		panic("argument to Contains() is nil")
	}
	return b.Get(key) != nil
}
