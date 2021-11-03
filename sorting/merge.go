package sorting

import "reflect"

// import "log"

// Merging: combining two ordered arrays to make one larger ordered array.
// This operation immediately lends itself to a simple recursive sort method
// known as mergesort: to sort an array, divide it into two halves, sort the
// two halves (recursively), and then merge the results.

// Mergesort guarantees to sort an array of N items in time proportional to NlogN,
// no matter what the input. Its prime disadvantage is that it uses extra space
// proportional to N.

// Top-down mergesort. It is one of the best-known examples of utility of the
// divide-and-conquer paradigm for efficient algorithm design.
func Mergesort(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String() // sorting.IntSortSlice
	// t := (reflect.TypeOf(x)).Name() // IntSortSlice
	// log.Println(t)

	// convert type
	switch t {
	case "sorting.IntSortSlice":
		a := x.(IntSortSlice)
		aux := make(IntSortSlice, n)
		sortInts(a, aux, 0, n-1)
	case "sorting.Float64SortSlice":
		a := x.(Float64SortSlice)
		aux := make(Float64SortSlice, n)
		sortFloat64s(a, aux, 0, n-1)
	case "sorting.StringSortSlice":
		a := x.(StringSortSlice)
		aux := make(StringSortSlice, n)
		sortStrings(a, aux, 0, n-1)
	}

}

// mergesort x[lo..hi] using auxiliary array aux[lo..hi]
func sortInts(x, aux IntSortSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortInts(x, aux, lo, mid)
	sortInts(x, aux, mid+1, hi)
	mergeInts(x, aux, lo, mid, hi)
}
func sortFloat64s(x, aux Float64SortSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortFloat64s(x, aux, lo, mid)
	sortFloat64s(x, aux, mid+1, hi)
	mergeFloat64s(x, aux, lo, mid, hi)
}
func sortStrings(x, aux StringSortSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortStrings(x, aux, lo, mid)
	sortStrings(x, aux, mid+1, hi)
	mergeStrings(x, aux, lo, mid, hi)
}

// Abstract in-place merge
// The method puts the results of merging the subarrays a[lo..mid] with a[mid+1..hi]
// into a single ordered array, leaving the result in a[lo..hi]. While it would be
// desirable to implement this mothod without using a signfifcant amount of extra
// space, such solutions are remarkably complicated. Instead, merge() copies everything
// to an auxiliary array and then merges back to the original.

// stably merge x[lo..mid] with a[mid+1..hi] using aux[lo..hi]
func mergeInts(x, aux IntSortSlice, lo, mid, hi int) {
	// copy to aux[]
	copy(aux, x)

	// merge back to x[]
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			// not support aux[j++]
			x[k] = aux[j]
			j++
		} else if j > hi {
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
	}
}
func mergeFloat64s(x, aux Float64SortSlice, lo, mid, hi int) {
	// copy to aux[]
	copy(aux, x)

	// merge back to x[]
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			// not support aux[j++]
			x[k] = aux[j]
			j++
		} else if j > hi {
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
	}
}
func mergeStrings(x, aux StringSortSlice, lo, mid, hi int) {
	// copy to aux[]
	copy(aux, x)

	// merge back to x[]
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			// not support aux[j++]
			x[k] = aux[j]
			j++
		} else if j > hi {
			x[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
	}
}

type Merge struct{}

func NewMerge() Sorter {
	return Merge{}
}

// Implements Sorter
func (s Merge) SortInts(x []int) {
	Mergesort(IntSortSlice(x))
}
func (s Merge) SortFloat64s(x []float64) {
	Mergesort(Float64SortSlice(x))
}
func (s Merge) SortStrings(x []string) {
	Mergesort(StringSortSlice(x))
}
