package pq

type MaxPQ struct {
	items []Item
	n int
}

// Implement interface PriorityQueue
func (pq *MaxPQ) Insert(item Item) {
	// double size of array if necessary
	if pq.n == len(pq.items)-1 {
		pq.resize(2*len(pq.items))
	}

	// add item, and percolate it up to maintain heap invariant
	pq.n++
	pq.items[pq.n]=item
	maxSwim(pq.items, pq.n)
}

func (pq *MaxPQ) Delete() Item {
	// if empty ..

	root := pq.items[1]
	swap(pq.items, 1, pq.n)
	pq.n--
	maxSink(pq.items, 1, pq.n)
	if  pq.n > 0 && pq.n == (len(pq.items)-1)/4 {
		pq.resize(len(pq.items)/2)
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

// factory method
func NewMaxPQN(n int) *MaxPQ {
	items := make([]Item, n+1)
	return &MaxPQ{items, 0}
}
func NewMaxPQ() *MaxPQ {
	items := make([]Item, 1)
	return &MaxPQ{items, 0}
}