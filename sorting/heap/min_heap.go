package heap

import "log"

type MinHeap struct {
	maxN  int // maximum number of elements
	n     int // number of elements
	items []Comparable
	BinaryHeap
}

func (h *MinHeap) Insert(item Comparable) {
	if h.IsFull() {
		panic("heap is full")
	}

	h.items[h.n] = item
	h.Swim(h.n)
	h.n++ // after swim
}

func (h *MinHeap) Remove() Comparable {
	if h.IsEmpty() {
		panic("heap is empty")
	}

	item := h.GetHighestPriority()
	h.swap(0, h.n-1)
	h.Sink(0)
	h.n--

	return item
}

func (h *MinHeap) GetHighestPriority() Comparable {
	return h.items[0]
}

func (h *MinHeap) Sink(p int) {

	smallerChild := h.getSmallerChildIndex(p)

	// if the left and right child do not exist
	// stop sinking
	if smallerChild == -1 {
		return
	}

	// compare the smaller child with the current index to see
	// if swap and further sink is needed
	if h.less(smallerChild, p) {
		h.swap(smallerChild, p)
		h.Sink(p)
	}

}

func (h *MinHeap) Swim(c int) {
	parent := h.GetParentIndex(c, h.n)
	log.Printf("%v, value:%v", parent, h.items[c])
	if parent != -1 && h.less(c, parent) {
		h.swap(c, parent)
		h.Swim(parent)
	}
}

func (h *MinHeap) IsFull() bool {
	return h.n == h.maxN
}

func (h *MinHeap) IsEmpty() bool {
	return h.n == 0
}

func (h *MinHeap) GetItems() []Comparable {
	return h.items[0:h.n]
}

// find the minimum of the left and right child elements
// if do not exist, return -1
func (h *MinHeap) getSmallerChildIndex(p int) int {
	leftChild := h.GetLeftChildIndex(p, h.n)
	rightChild := h.GetRightChildIndex(p, h.n)

	if leftChild != -1 && rightChild != -1 {
		if h.less(leftChild, rightChild) {
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

func (h *MinHeap) getItemAtIndex(i int) Comparable {
	return h.items[i]
}

// if items[i] < items[j]
func (h *MinHeap) less(i, j int) bool {
	one := h.getItemAtIndex(i)
	another := h.getItemAtIndex(j)
	return one.CompareTo(another) < 0
}

func (h *MinHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func NewMinHeap() *MinHeap {
	items := make([]Comparable, DefaultMaxSize)

	return &MinHeap{DefaultMaxSize, 0, items, NewBinaryHeap()}
}
