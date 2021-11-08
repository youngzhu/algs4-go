package heap

type MaxHeap struct {
	maxN  int // maximum number of elements
	n     int // number of elements
	items []Comparable
	BinaryHeap
}

func NewMaxHeap() *MaxHeap {
	items := make([]Comparable, defaultMaxSize)

	return &MaxHeap{defaultMaxSize, 0, items, NewBinaryHeapBased1()}
}

func (h *MaxHeap) Remove() Comparable {
	if h.IsEmpty() {
		panic("heap is empty")
	}

	item := h.GetHighestPriority()
	root := h.GetRootIndex()
	lastLeaf := h.GetLastLeafIndex(h.n)

	h.swap(root, lastLeaf)
	h.Sink(root, lastLeaf-1)

	h.n--

	return item
}

func (h *MaxHeap) GetHighestPriority() Comparable {
	return h.items[0]
}

func (h *MaxHeap) Sink(p, max int) {

	elderChild := h.getElderChildIndex(p, max)

	// if the left and right child do not exist
	// stop sinking
	if elderChild == -1 {
		return
	}

	// compare the elder child with the current index to see
	// if swap and further sink is needed
	if h.less(p, elderChild) {
		h.swap(elderChild, p)
		h.Sink(elderChild, max)
	}

}

func (h *MaxHeap) Swim(c, max int) {
	parent := h.GetParentIndex(c, max)

	if parent != -1 && h.less(parent, c) {
		h.swap(c, parent)
		h.Swim(parent, max)
	}
}

func (h *MaxHeap) IsFull() bool {
	return h.n == h.maxN
}

func (h *MaxHeap) IsEmpty() bool {
	return h.n == 0
}

func (h *MaxHeap) GetItems() []Comparable {
	begin := h.GetRootIndex()
	end := h.GetLastLeafIndex(h.n)
	return h.items[begin : end+1]
}

// find the maximum of the left and right child elements
// if do not exist, return -1
func (h *MaxHeap) getElderChildIndex(p, max int) int {
	leftChild := h.GetLeftChildIndex(p, max)
	rightChild := h.GetRightChildIndex(p, max)

	if leftChild != -1 && rightChild != -1 {
		if h.less(leftChild, rightChild) {
			return rightChild
		} else {
			return leftChild
		}
	} else if leftChild != -1 {
		return leftChild
	} else if rightChild != -1 {
		return rightChild
	} else {
		return -1
	}
}

func (h *MaxHeap) getItemAtIndex(i int) Comparable {
	return h.items[i]
}

// if items[i] < items[j]
func (h *MaxHeap) less(i, j int) bool {
	one := h.getItemAtIndex(i)
	another := h.getItemAtIndex(j)
	return one.CompareTo(another) < 0
}

func (h *MaxHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
