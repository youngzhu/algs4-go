package sorting

// See Go src sort.Interface
// An implementation of Sortable can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Sortable interface {
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
	SortFloat64s(x []float64)
	SortStrings(x []string)
}

// Reports whether data is sorted
func IsSorted(data Sortable) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
func IsSortedInts(data []int) bool {
	return IsSorted(IntCompSlice(data[0:]))
}

// Convenience types for common cases

type (
	// Attaches the methods of Sortable to []int
	// sorting in increasing order
	IntCompSlice []int

	// Implements Sortable for a []folat64
	// sorting in increasing order
	// with not-a-number (NaN) values ordered before other values
	Float64CompSlice []float64

	// Attaches the methods of Sortable to []string
	// sorting in increasing order
	StringCompSlice []string
)

// IntCompSlice implements
func (x IntCompSlice) Len() int {
	return len(x)
}
func (x IntCompSlice) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x IntCompSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Float64CompSlice implements
func (x Float64CompSlice) Len() int {
	return len(x)
}

// Note that floating-point comparison by itself is not a transitive relation:
// it does not report a consistent ordering for not-a-number (NaN) values.
// This implementation places NaN values before others, by using:
//
// x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
//
func (x Float64CompSlice) Less(i, j int) bool {
	return x[i] < x[j] || (isNaN(x[i]) && !isNaN(x[j]))
}
func (x Float64CompSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// A copy of math.IsNaN to avoid a dependency on the math package
func isNaN(f float64) bool {
	return f != f
}

// StringCompSlice implements
func (x StringCompSlice) Len() int {
	return len(x)
}
func (x StringCompSlice) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x StringCompSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
