package pq

// binary heap using 0-based indexing
type BinaryHeapBased0 struct{}

// factory method
func NewBinaryHeapBased0() BinaryHeapBased0 {
	return BinaryHeapBased0{}
}

// 0-based
func (bh BinaryHeapBased0) GetLeftChildIndex(p, n int) int {
	leftChild := 2*p + 1
	if leftChild >= n {
		return -1 // no valid left child
	}

	return leftChild
}

func (bh BinaryHeapBased0) GetRightChildIndex(p, n int) int {
	rightChild := 2*p + 2
	if rightChild >= n {
		return -1 // no valid right child
	}

	return rightChild
}

func (bh BinaryHeapBased0) GetParentIndex(c, n int) int {
	if c < 0 || c > n {
		return -1
	}

	return (c - 1) / 2
}

func (bh BinaryHeapBased0) GetRootIndex() int {
	return 0
}

func (bh BinaryHeapBased0) GetLastLeafIndex(n int) int {
	return n - 1
}
