package pq
/*
// Minimum-oriented indexed PQ implementation using a binary heap.
type MinIndexPQ struct {
	maxN int // maximum number of elements on PQ
	n int // number of elements on PQ
	pq []int // binary heap
	qp []int // inverse of pq: qp[pq[i]] = pq[qp[i]] = i
	items []Item // items[i] = priority of i
}

func NewMinIndexPQ(maxN int) *MinIndexPQ {
	pq := make([]int, maxN+1)
	// qp := make([]int, maxN, 1)
	items := make([]Item, maxN)
	return &MinIndexPQ{maxN, 0, pq, qp, items}
}

// Associates item with index i
func (pq *MinIndexPQ) Insert(i int, item Item) {
	validateIndex(i, pq.maxN)
	if pq.Contains(i) {
		panic("index is already in the priority queue")
	}

	pq.n++
	qp.qp[i]=pq.n
	pq.pq[pq.n]=i
	pq.items[i]=item
	minSwim(pq.items, pq.n)
}

// Returns an index associated with a minimum key.
func (pq *MinIndexPQ) MinIndex() int {
	if pq.n == 0 {
		panic("Priority queue is empty")
	}
	return pq.pq[1]
}

// Returns a minimum item
func (pq *MinIndexPQ) MinItem() Item {
	return pq.items[pq.MinIndex()]
}

// Removes a minimum item and returns its associated index
func (pq *MinIndexPQ) DeleteMin() int {
	minIdx := pq.MinIndex()
	sawp(pq.items, 1, pq.n)
	pq.n--
	minSink(1)
	pq.qp[minIdx]=-1 // delete
	return minIdx
}

func (pq *MinIndexPQ) Update(i int, item Item) {
	validateIndex(i, pq.maxN)
	if !pq.Contains(i) {
		panic("index is not in the priority queue")
	}

	pq.items[i] = item
	minSwim(pq.items, pq.qp[i])
	minSink(pq.items, pq.qp[i])
}



func (pq *MinIndexPQ) Contains(i int) bool {
	validateIndex(i, pq.maxN)
	return pq.qp[i] != -1
}

func (pq *MinIndexPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *MinIndexPQ) Size() int {
	return pq.n
}

func validateIndex(idx, max int) {
	if idx < 0 {
		panic("index is negative: " + idx)
	}
	if idx >= max {
		panic("index >= capacity:" + idx)
	}
}
*/