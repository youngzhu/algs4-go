package sorting

// Optimized Mergesort

import "reflect"

// Improvements
// We can cut the running time of mergesort substantially with some
// carefully considered modifications to the implementation.

// 1. Use insertion sort for small subarrays.
// We can improve most recursive algorithms by handling samll cases
// differently. Switching to insertion sort for small subarrays will
// improve the running time of a typical mergesort implementation by
// 10 to 15 percent.

// See sort_compare.go

func MergesortX1(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String() // sorting.IntCompSlice
	// t := (reflect.TypeOf(x)).Name() // IntCompSlice
	// log.Println(t)

	// convert type
	switch t {
	case "sorting.IntCompSlice":
		a := x.(IntCompSlice)
		aux := make(IntCompSlice, n)
		sortIntsX1(a, aux, 0, n-1)
	case "sorting.Float64CompSlice":
		a := x.(Float64CompSlice)
		aux := make(Float64CompSlice, n)
		sortFloat64sX1(a, aux, 0, n-1)
	case "sorting.StringCompSlice":
		a := x.(StringCompSlice)
		aux := make(StringCompSlice, n)
		sortStringsX1(a, aux, 0, n-1)
	}

}

// mergesort x[lo..hi] using auxiliary array aux[lo..hi]
func sortIntsX1(x, aux IntCompSlice, lo, hi int) {
	// improvment 1. Use insertion sort for small subarrays.
	if hi <= lo+CUTOFF {
		insertionSort(x, lo, hi)
		return
	}
	mid := lo + (hi-lo)/2
	sortIntsX1(x, aux, lo, mid)
	sortIntsX1(x, aux, mid+1, hi)
	mergeIntsX1(x, aux, lo, mid, hi)
}
func sortFloat64sX1(x, aux Float64CompSlice, lo, hi int) {
	if hi <= lo+CUTOFF {
		insertionSort(x, lo, hi)
		return
	}
	mid := lo + (hi-lo)/2
	sortFloat64sX1(x, aux, lo, mid)
	sortFloat64sX1(x, aux, mid+1, hi)
	mergeFloat64sX1(x, aux, lo, mid, hi)
}
func sortStringsX1(x, aux StringCompSlice, lo, hi int) {
	if hi <= lo+CUTOFF {
		insertionSort(x, lo, hi)
		return
	}
	mid := lo + (hi-lo)/2
	sortStringsX1(x, aux, lo, mid)
	sortStringsX1(x, aux, mid+1, hi)
	mergeStringsX1(x, aux, lo, mid, hi)
}

// stably merge x[lo..mid] with a[mid+1..hi] using aux[lo..hi]
func mergeIntsX1(x, aux IntCompSlice, lo, mid, hi int) {
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
func mergeFloat64sX1(x, aux Float64CompSlice, lo, mid, hi int) {
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
func mergeStringsX1(x, aux StringCompSlice, lo, mid, hi int) {
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

type MergeX1 struct{}

// Implements Sorter
func (s MergeX1) SortInts(x []int) {
	MergesortX1(IntCompSlice(x))
}
func (s MergeX1) SortFloat64s(x []float64) {
	MergesortX1(Float64CompSlice(x))
}
func (s MergeX1) SortStrings(x []string) {
	MergesortX1(StringCompSlice(x))
}
