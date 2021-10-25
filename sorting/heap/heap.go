package heap

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

const defaultMaxSize = 40

type (
	IntItem    int
	StringItem string
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

func (s StringItem) CompareTo(other Comparable) int {
	ss := other.(StringItem)
	if s < ss {
		return -1
	} else if s > ss {
		return 1
	} else {
		return 0
	}
}
