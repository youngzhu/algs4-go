package pq

// We represent complete binary trees sequentially within an array by putting
// the nodes with level order, with the root at position 1, its children at
// positions 2 and 3, their children in positions 4, 5, 6 and 7, and so on.

// In a heap, the parent of the node in position k is in position k/2; and, conversely,
// the two children of the node in position k are in positions 2k and 2k+1. We
// can travel up and down by doing simple arithmetic on array indices: to move
// up the tree from a[k] we set k to k/2; to move down the tree we set k to 2*k or 2*k+1.

// binary heap using 1-based indexing
type BinaryHeapBased1 struct{}

func NewBinaryHeapBased1() BinaryHeapBased1 {
	return BinaryHeapBased1{}
}

// 1-based
func (bh BinaryHeapBased1) GetLeftChildIndex(p, n int) int {
	leftChild := 2 * p
	if leftChild > n {
		return -1 // no valid left child
	}

	return leftChild
}

func (bh BinaryHeapBased1) GetRightChildIndex(p, n int) int {
	rightChild := 2*p + 1
	if rightChild > n {
		return -1 // no valid right child
	}

	return rightChild
}

func (bh BinaryHeapBased1) GetParentIndex(c, n int) int {
	if c <= 1 || c > n {
		return -1
	}

	return c / 2
}

func (bh BinaryHeapBased1) GetRootIndex() int {
	return 1
}

func (bh BinaryHeapBased1) GetLastLeafIndex(n int) int {
	return n
}
