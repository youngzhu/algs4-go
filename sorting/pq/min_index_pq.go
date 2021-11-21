package pq

import (
	"fmt"
)

// Minimum-oriented indexed PQ implementation using a binary heap.
type MinIndexPQ struct {
	maxN  int    // maximum number of elements on PQ
	n     int    // number of elements on PQ
	pq    []int  // binary heap using 1-based indexing
	qp    []int  // inverse of pq: qp[pq[i]] = pq[qp[i]] = i
	items []Item // items[i] = priority of i
	BinaryHeap
}

func NewMinIndexPQ(maxN int) *MinIndexPQ {
	heap := NewBinaryHeap()
	size := heap.GetLastLeafIndex(maxN) + 1

	pq := make([]int, size)
	qp := make([]int, size)
	items := make([]Item, size)

	for i := 0; i < size; i++ {
		qp[i] = -1
	}

	return &MinIndexPQ{maxN, 0, pq, qp, items, heap}
}

// Associates item with index i
func (ipq *MinIndexPQ) Insert(i int, item Item) {
	ipq.validateIndex(i)
	if ipq.Contains(i) {
		panic("index is already in the priority queue")
	}

	ipq.n++

	idx := ipq.getLastIndex()
	ipq.qp[i] = idx
	ipq.items[i] = item
	ipq.pq[idx] = i
	ipq.swim(idx, idx)
}

func (ipq *MinIndexPQ) swim(child, max int) {
	root := ipq.GetRootIndex()

	for child > root {
		parent := ipq.GetParentIndex(child, max)
		if ipq.isHigherPriority(parent, child) {
			break
		}
		ipq.swap(parent, child)
		child = parent
	}
}

// if pq.items[i] priority higher than pq.items[j]
func (ipq *MinIndexPQ) isHigherPriority(i, j int) bool {
	ii, jj := ipq.pq[i], ipq.pq[j]
	i1 := ipq.items[ii]
	i2 := ipq.items[jj]
	// different from MaxPQ
	return i1.CompareTo(i2) <= 0
}

func (ipq *MinIndexPQ) swap(i, j int) {
	ipq.pq[i], ipq.pq[j] = ipq.pq[j], ipq.pq[i]
	ii, jj := ipq.pq[i], ipq.pq[j]
	ipq.qp[ii], ipq.qp[jj] = i, j
}

// Returns an index associated with a minimum key.
func (ipq *MinIndexPQ) MinIndex() int {
	if ipq.n == 0 {
		panic("Priority queue is empty")
	}
	return ipq.pq[ipq.GetRootIndex()]
}

// Returns a minimum item
func (ipq *MinIndexPQ) MinItem() Item {
	return ipq.items[ipq.MinIndex()]
}

func (ipq *MinIndexPQ) HighestPriorityItem() Item {
	return ipq.MinItem()
}

// Removes a minimum item and returns its associated index
func (ipq *MinIndexPQ) Delete() int {
	minIdx := ipq.MinIndex()

	rootIdx := ipq.GetRootIndex()
	lastLeaf := ipq.GetLastLeafIndex(ipq.n)

	ipq.swap(rootIdx, lastLeaf)
	ipq.sink(rootIdx, lastLeaf-1)

	ipq.qp[minIdx] = -1 // delete

	ipq.n--

	return minIdx
}

func (ipq *MinIndexPQ) sink(parent, max int) {
	for {
		higherPriorityChild := ipq.getHigherPriorityChild(parent, max)

		// if the left and right child do not exist
		// stop sinking
		if higherPriorityChild == -1 {
			break
		}

		if ipq.isHigherPriority(parent, higherPriorityChild) {
			break
		}

		ipq.swap(higherPriorityChild, parent)
		parent = higherPriorityChild
	}
}

func (ipq *MinIndexPQ) getHigherPriorityChild(parent, max int) int {
	leftChild := ipq.GetLeftChildIndex(parent, max)
	rightChild := ipq.GetRightChildIndex(parent, max)

	if leftChild != -1 && rightChild != -1 {
		if ipq.isHigherPriority(leftChild, rightChild) {
			return leftChild
		} else {
			return rightChild
		}
	} else if leftChild != -1 {
		return leftChild
	} else if rightChild != -1 {
		return rightChild
	} else {
		return -1
	}

}

func (ipq *MinIndexPQ) Update(i int, item Item) {
	ipq.validateIndex(i)
	if !ipq.Contains(i) {
		panic("index is not in the priority queue")
	}

	ipq.items[i] = item
	ipq.swim(i, ipq.n)
	ipq.sink(i, ipq.n)
}

func (ipq *MinIndexPQ) Contains(i int) bool {
	ipq.validateIndex(i)
	return ipq.qp[i] != -1
}

func (ipq *MinIndexPQ) IsEmpty() bool {
	return ipq.n == 0
}

func (ipq *MinIndexPQ) Size() int {
	return ipq.n
}

func (ipq *MinIndexPQ) String() string {
	qp := fmt.Sprintf("qp:%v", ipq.qp)
	items := fmt.Sprintf("items:%v", ipq.items)
	pq := fmt.Sprintf("pq:%v", ipq.pq)
	return fmt.Sprintf("%v, %v, %v", qp, items, pq)
}

func (ipq MinIndexPQ) validateIndex(idx int) {
	if idx < 0 {
		panic("index is negative")
	}
	if idx >= ipq.maxN {
		panic("index >= capacity")
	}
}

func (ipq MinIndexPQ) getLastIndex() int {
	return ipq.GetLastLeafIndex(ipq.n)
}
