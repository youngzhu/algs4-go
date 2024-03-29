package pq

// Many applications require that we process items having keys in order, but not
// necessarily in full sorted order and not necessarily all at once. Often, we
// collect a set of items, then process the one with the largest key, then perhaps
// collect more items, then process the one with the current largest key, and so forth.
// An appropriate data type in such an environment supports two operations: remove
// and insert. Such a data type is called a priority queue.

type PriorityQueue interface {
	Insert(item Item)
	Delete() Item
	IsEmpty() bool
	Size() int
}

// IndexPriorityQueue
// In many applications, it makes sense to allow client refer to items that
// are already on the priority queue. One easy way to do so is to associate
// a unique integer index with each item.
type IndexPriorityQueue interface {
	// Insert inserts item, associate it with k
	Insert(k int, item Item)

	// Update updates the item associated with k to item
	Update(k int, item Item)

	// Contains report is k associated with any item?
	Contains(k int) bool

	// Delete remove the minimal/maximal item and return its index
	Delete() int

	// IsEmpty report is the priority queue empty?
	IsEmpty() bool

	// Size return number of items in the priority queue
	Size() int

	// HighestPriorityItem get the minimal/maximal item
	HighestPriorityItem() Item
}

type Item interface {
	CompareTo(x Item) int
}

type StringItem string

func (s StringItem) CompareTo(x Item) int {
	ss := x.(StringItem)
	if s < ss {
		return -1
	} else if s > ss {
		return 1
	} else {
		return 0
	}
}

type IntItem int

func (i IntItem) CompareTo(x Item) int {
	ii := x.(IntItem)
	if i < ii {
		return -1
	} else if i > ii {
		return 1
	} else {
		return 0
	}
}
