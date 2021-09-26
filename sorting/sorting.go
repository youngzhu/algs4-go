package sorting

// See Go src sort.Interface
// An implementation of Comparable can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Comparable interface {
	// The number of elements in the collection
	Len() int

	// Reports whether the element with index i
	// must sort before the element with index j
	Less(i, j int) bool

	// Swap the elements with indexes i and j
	Swap(i, j int)
}

type Sorter interface {
	SortInts(x []int)
	// SortFloat64s(x []float64)
	// SortStrings(x []string)
}

// Reports whether data is sorted
func IsSorted(data Comparable) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type (
	IntCompSlice []int
)

func (x IntCompSlice) Len() int {
	return len(x)
}
func (x IntCompSlice) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x IntCompSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
