package searching

import "github.com/youngzhu/algs4-go/fund"

// 2-3 search tree, is a tree that either is empty or:
// 1) A 2-node, with one key (and associated value) and two links, a left link
// to a 2-3 search tree with smaller keys, and a right link to a 2-3 search 
// tree with larger keys
// 2) A 3-node, with two keys (and associated values) and three links, a left
// link to a 2-3 search tree with smaller keys, a middle link to a 2-3 search
// tree with keys between the node's keys and a right link to a 2-3 search tree
// with larger keys

// Red-black BST
// Encoding 3-nodes.
// The basic idea behind red-black BST is to encode 2-3 trees by starting with
// standard BST (which are made up of 2-nodes) and adding extra information to
// encode 3-nodes. We think of the links as being of two different types: red
// links, which bind together two 2-nodes to represent 3-nodes; black links, 
// which bind together the 2-3 tree. Specifically, we represent 3-nodes as two
// 2-nodes connected by a single red link that leans left. We refer to BST that
// represent 2-3 trees in this way as red-black BST.
// One advantage of using such a representation is that it allow us to use our
// get() code for standard BST search without modification.

// A 1-1 correspondence.
// Given any 2-3 tree, we can immdiately derive a corresponding red-black BST, 
// just by converting each node as specified. Conversely, if we draw the red
// links horizontally in a red-black BST, all of the null links are the same
// distance from the the root, and if we then collapse together the nodes 
// connected by red links, the result is a 2-3 tree.

// Color representation.
// Since each node is pointed to by precisely one link (from its parent), we
// encode the color of links in nodes, by adding a boolean variable color to
// our node, which is true if the link from the parent is red and false if it
// is black. By convention, null links are black.

// Rotations.
// The implementation that we will consider might allow right-leaning red links
// or two red-links in a row during an operation, but it always corrects these
// before completion, through judicious use of an operation called rotation that
// switches orientation of red links. First, suppose that we have a right-leaning
// red link that needs to be rotated to lean to the left. This operation is called
// a left rotation. Implementing a right rotation that converts a left-leaning red
// link to a right-leaning one amounts to the same code, with left and right 
// interchanged.

// Flipping colors.
// The implementation that we will consider might also allow a black parent to have
// two red children. The flip operation flips the colors of the two red children to
// black and the color of the black parent to red.

const (
	RED = true
	BLACK = false
)

// a left-leaning red-black BST
type RedBlackBST struct {
	root *RBNode // root of the BST
}

// 
type RBNode struct {
	key STKey // key
	value STValue // associated data
	left, right *RBNode // links to left and right subtrees
	color bool // color of parent link
	size int // subtree count
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func newRBNode(key STKey, value STValue, color bool, size int) *RBNode {
	return &RBNode{key: key, value: value, color: color, size: size}
}

// Returns the value associated with the given key
func (rb *RedBlackBST) Get(key STKey) STValue {
	if key == nil {
		panic("argument to Get() is nil")
	}
	return getRB(rb.root, key)
}

// value associated with the given key in subtree rooted at x
// return nil if no such key
func getRB(x *RBNode, key STKey) STValue {
	for x != nil {
		cmp := key.CompareTo(x.key)

		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.value
		}
	}
	return nil 
}

// Inserts the specified key-value pair into the ST, overwriting the old value
// with the new one if the ST already contains the specified key.
// Deletes the specified key (and its associated value) from this ST if the 
// specified value is nil
func (rb *RedBlackBST) Put(key STKey, value STValue) {
	if key == nil {
		panic("first arg to Put() is nil")
	}
	if value == nil {
		rb.Delete(key)
		return
	}
	rb.root = putRB(rb.root, key, value)
	rb.root.color = BLACK
}

// insert the key-value pair into the subtree rooted at h
func putRB(h *RBNode, key STKey, value STValue) *RBNode {
	if h == nil {
		return newRBNode(key, value, RED, 1)
	}

	cmp := key.CompareTo(h.key)
	if cmp < 0 {
		h.left = putRB(h.left, key, value)
	} else if cmp > 0 {
		h.right = putRB(h.right, key, value)
	} else {
		h.value = value
	}

	// fix-up any right-leaning links
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}

	h.size = sizeRB(h.left) + sizeRB(h.right) + 1

	return h
}

// Removes the specified key and its associated value from this ST
// (if the key is in the ST)
func (rb *RedBlackBST) Delete(key STKey) {
	if key == nil {
		panic("arg to Delete() is nil")
	}
	if !rb.Contains(key) {
		return
	}

	// if both children of root are black, set root to red
	root := rb.root
	if !isRed(root.left) && !isRed(root.right) {
		root.color = RED
	}

	root = deleteRB(root, key)
	if !rb.IsEmpty() {
		root.color = BLACK
	}
	rb.root = root
}

// delete the key-value with the given key rooted at h
func deleteRB(h *RBNode, key STKey) *RBNode {
	if key.CompareTo(h.key) < 0 {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = deleteRB(h.left, key)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if key.CompareTo(h.key) == 0 && h.right == nil {
			return nil
		}
		if !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		if key.CompareTo(h.key) == 0 {
			x := minRB(h.right)
			h.key = x.key
			h.value = x.value
			h.right = deleteMinRB(h.right)
		} else {
			h.right = deleteRB(h.right, key)
		}
	}

	return balance(h)
}

// Returns the smallest key in the ST
func (rb *RedBlackBST) Min() STKey {
	if rb.IsEmpty() {
		panic("calls Min() with empty ST")
	}
	return minRB(rb.root).key
}

// the smallest key in subtree rooted at x
// nil if no such key
func minRB(x *RBNode) *RBNode {
	if x.left == nil {
		return x
	} else {
		return minRB(x.left)
	}
}

// Retruns the largest key in the ST
func (rb *RedBlackBST) Max() STKey {
	if rb.IsEmpty() {
		panic("calls Max() with empty ST")
	}
	return maxRB(rb.root).key
}

// the largest key in subtree rooted at x
// nil if no such key
func maxRB(x *RBNode) *RBNode {
	if x.right == nil {
		return x
	} else {
		return maxRB(x.right)
	}
}

// Removes the smallest key and associated value from the ST
func (rb *RedBlackBST) DeleteMin() {
	if rb.IsEmpty() {
		panic("BST is empty")
	}

	// if both children of root are balck, set root to red
	root := rb.root
	if !isRed(root.left) && !isRed(root.right) {
		root.color = RED
	}

	root = deleteMinRB(root)
	rb.root = root
	if !rb.IsEmpty() {
		rb.root.color = BLACK
	}
}

// delete the key-value pair with the minimum key rooted at x
func deleteMinRB(x *RBNode) *RBNode {
	if x.left == nil {
		return nil
	}

	if !isRed(x.left) && !isRed(x.left.left) {
		x = moveRedLeft(x)
	}

	x.left = deleteMinRB(x.left)

	return balance(x)
}

// Return all keys in the ST
func (rb *RedBlackBST) Keys() []STKey {
	if rb.IsEmpty() {
		panic("The BST is empty")
	}
	return rb.rangeKeys(rb.Min(), rb.Max())
}

// returns all keys in the symbol table in the given range
func (rb *RedBlackBST) rangeKeys(lo, hi STKey) []STKey {
	if lo == nil {
		panic("first argument to rangeKeys() is nil")
	}
	if hi == nil {
		panic("second argument to rangeKeys() is nil")
	}

	queue := fund.NewQueue()
	keysRB(rb.root, queue, lo, hi)

	keySliece := make([]STKey, queue.Size())

	for i := 0; !queue.IsEmpty(); i++ {
		keySliece[i] = queue.Dequeue().(STKey)
	}

	return keySliece
}

// add the keys between lo and hi in the subtree rooted at x to the queue
func keysRB(x *RBNode, queue *fund.Queue, lo, hi STKey) {
	if x == nil {
		return 
	}

	cmpLo := lo.CompareTo(x.key)
	cmpHi := hi.CompareTo(x.key)
	if cmpLo < 0 {
		keysRB(x.left, queue, lo, hi)
	}
	if cmpLo <= 0 && cmpHi >= 0 {
		queue.Enqueue(x.key)
	}
	if cmpHi > 0 {
		keysRB(x.right, queue, lo, hi)
	}
}

/**** Helper func ****/

// is node x red
func isRed(x *RBNode) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

// make a left-leaning link lean to the right
func rotateRight(h *RBNode) *RBNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = x.right.color
	x.right.color = RED
	x.size = h.size
	h.size = sizeRB(h.left) + sizeRB(h.right) + 1
	return x
}

// make a right-leaning link lean to the left
func rotateLeft(h *RBNode) *RBNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = x.left.color
	x.left.color = RED
	x.size = h.size
	h.size = sizeRB(h.left) + sizeRB(h.right) + 1
	return x
}

// flip the colors of a node and its two children
func flipColors(h *RBNode) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// assuming that h is red and both h.left and h.left.left are black
// make h.left or one of its chidren red
func moveRedLeft(h *RBNode) *RBNode {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

// assuming that h is red and both h.right and h.right.left are black
// make h.right or one of its children red
func moveRedRight(h *RBNode) *RBNode {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

// restore red-black tree invariant
func balance(h *RBNode) *RBNode {
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}

	h.size = sizeRB(h.left) + sizeRB(h.right) + 1

	return h
}

// number of node in subtree rooted at x
func sizeRB(x *RBNode) int {
	if x == nil  {
		return 0
	}
	return x.size
}

// Returns the nubmer of key-value pairs in this BST
func (rb *RedBlackBST) Size() int {
	return sizeRB(rb.root)
}

// Is this BST empty?
func (rb *RedBlackBST) IsEmpty() bool {
	return rb.root == nil
}

// Does this ST contain the given key?
func (rb *RedBlackBST) Contains(key STKey) bool {
	return rb.Get(key) == nil
}