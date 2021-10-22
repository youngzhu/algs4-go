package pq

type MinPQ struct {
	items []Item
	n int
}

// Implement interface PriorityQueue
func (pq *MinPQ) Insert(item Item) {
	// double size of array if necessary
	if pq.n == len(pq.items)-1 {
		pq.resize(2*len(pq.items))
	}

	// add item, and percolate it up to maintain heap invariant
	pq.n++
	pq.items[pq.n]=item
	minSwim(pq.items, pq.n)
}

func (pq *MinPQ) Delete() Item {
	if pq.IsEmpty() {
		panic("The priority queue is empty")
	}

	root := pq.items[1]
	swap(pq.items, 1, pq.n)
	pq.n--
	minSink(pq.items, 1, pq.n)
	if  pq.n > 0 && pq.n == (len(pq.items)-1)/4 {
		pq.resize(len(pq.items)/2)
	}
	return root
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

// factory method
func NewMinPQN(n int) *MinPQ {
	items := make([]Item, n+1)
	return &MinPQ{items, 0}
}
func NewMinPQ() *MinPQ {
	items := make([]Item, 1)
	return &MinPQ{items, 0}
}