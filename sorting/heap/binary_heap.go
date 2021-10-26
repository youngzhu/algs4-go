package heap

type BinaryHeap interface {
	GetLeftChildIndex(p, n int) int
	GetRightChildIndex(p, n int) int
	GetParentIndex(c, n int) int
	GetRootIndex() int
	GetLastLeafIndex(n int) int
}

// binary heap using 0-based indexing
type BinaryHeapBased0 struct{}

// binary heap using 1-based indexing
type BinaryHeapBased1 struct{}

// factory method
func NewBinaryHeap() BinaryHeap {
	return BinaryHeapBased0{}
}

func NewBinaryHeapBased0() BinaryHeapBased0 {
	return BinaryHeapBased0{}
}

func NewBinaryHeapBased1() BinaryHeapBased1 {
	return BinaryHeapBased1{}
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
