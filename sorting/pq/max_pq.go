package pq

type MaxPQ struct {
	items []Item
	n     int
	BinaryHeap
}

// factory method
func NewMaxPQN(n int) *MaxPQ {
	items := make([]Item, n+1)
	heap := NewBinaryHeapBased1()
	return &MaxPQ{items, 0, heap}
}
func NewMaxPQ() *MaxPQ {
	items := make([]Item, 1)
	heap := NewBinaryHeapBased1()
	return &MaxPQ{items, 0, heap}
}

// Implement interface PriorityQueue
func (pq *MaxPQ) Insert(item Item) {
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

func (pq *MaxPQ) swim(child, max int) {
	parent := pq.GetParentIndex(child, max)

	if parent != -1 && pq.less(parent, child) {
		swap(pq.items, parent, child)
		pq.swim(parent, max)
	}
}

// if pq.items[i] < pq.items[j]
func (pq *MaxPQ) less(i, j int) bool {
	i1 := pq.items[i]
	i2 := pq.items[j]
	return i1.CompareTo(i2) < 0
}

func (pq *MaxPQ) Delete() Item {
	if pq.IsEmpty() {
		panic("The priority queue is empty")
	}

	root := pq.items[1]
	swap(pq.items, 1, pq.n)
	pq.n--
	maxSink(pq.items, 1, pq.n)
	if pq.n > 0 && pq.n == (len(pq.items)-1)/4 {
		pq.resize(len(pq.items) / 2)
	}
	return root
}

func (pq *MaxPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *MaxPQ) Size() int {
	return pq.n
}

func (pq *MaxPQ) resize(capacity int) {
	t := make([]Item, capacity)
	copy(t, pq.items)
	pq.items = t
}
