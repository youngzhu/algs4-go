package pq

type MinPQ struct {
	items []Item
	n     int
	BinaryHeap
}

// factory method
func NewMinPQN(n int) *MinPQ {
	items := make([]Item, n+1)
	heap := NewBinaryHeap()
	return &MinPQ{items, 0, heap}
}
func NewMinPQ() *MinPQ {
	items := make([]Item, 1)
	return &MinPQ{items, 0, NewBinaryHeap()}
}

// Implement interface PriorityQueue
func (pq *MinPQ) Insert(item Item) {
	// double size of array if necessary
	if pq.n == len(pq.items)-1 {
		pq.resize(2 * len(pq.items))
	}

	// add item, and percolate it up to maintain heap invariant
	pq.n++

	lastLeaf := pq.GetLastLeafIndex(pq.n)
	pq.items[lastLeaf] = item
	pq.swim(lastLeaf, lastLeaf)
}

func (pq *MinPQ) swim(child, max int) {
	parent := pq.GetParentIndex(child, max)

	if parent != -1 && pq.isHigherPriority(child, parent) {
		swap(pq.items, parent, child)
		pq.swim(parent, max)
	}
}

// if pq.items[i] higher than pq.items[j]
func (pq *MinPQ) isHigherPriority(i, j int) bool {
	i1 := pq.items[i]
	i2 := pq.items[j]
	return i1.CompareTo(i2) < 0
}

func (pq *MinPQ) Delete() Item {
	if pq.IsEmpty() {
		panic("The priority queue is empty")
	}

	rootIdx := pq.GetRootIndex()
	lastLeaf := pq.GetLastLeafIndex(pq.n)

	item := pq.items[rootIdx]

	swap(pq.items, rootIdx, lastLeaf)

	pq.sink(rootIdx, lastLeaf-1)

	pq.n--

	if pq.n > 0 && pq.n == (len(pq.items)-1)/4 {
		pq.resize(len(pq.items) / 2)
	}

	return item
}

func (pq *MinPQ) sink(parent, max int) {
	higherPriorityChild := pq.getHighPriorityChild(parent, max)

	// if the left and right child do not exist
	// stop sinking
	if higherPriorityChild == -1 {
		return
	}

	if pq.isHigherPriority(higherPriorityChild, parent) {
		swap(pq.items, higherPriorityChild, parent)
		pq.sink(higherPriorityChild, max)
	}

}

func (pq *MinPQ) getHighPriorityChild(parent, max int) int {
	leftChild := pq.GetLeftChildIndex(parent, max)
	rightChild := pq.GetRightChildIndex(parent, max)

	if leftChild != -1 && rightChild != -1 {
		if pq.isHigherPriority(leftChild, rightChild) {
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

func (pq *MinPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *MinPQ) Size() int {
	return pq.n
}

func (pq *MinPQ) resize(capacity int) {
	t := make([]Item, capacity)
	copy(t, pq.items)
	pq.items = t
}
