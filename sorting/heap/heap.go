package heap

import "log"

type Comparable interface {
	CompareTo(x Comparable) int
}

type Heap interface {
	Insert(x Comparable)
	Remove() Comparable
	GetHighestPriority() Comparable
	Sink(i int)
	Swim(i int)
	IsEmpty() bool
	IsFull() bool
	Size() int
}

type BinaryHeap interface {
	GetLeftChildIndex(p, n int) int
	GetRightChildIndex(p, n int) int
	GetParentIndex(c, n int) int
}

// binary heap using 0-based indexing
type BinaryHeapBased0 struct{}

const DefaultMaxSize = 40

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
	log.Println(c, n)
	if c < 0 || c > n {
		return -1
	}

	return (c - 1) / 2
}

type BaseHeap struct {
}

type (
	IntItem int
)

func (i IntItem) CompareTo(x Comparable) int {
	ii := x.(IntItem)
	if i < ii {
		return -1
	} else if i > ii {
		return 1
	} else {
		return 0
	}
}

// factory method
func NewBinaryHeap() BinaryHeapBased0 {
	return BinaryHeapBased0{}
}
