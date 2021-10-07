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

// 2. Test whether array is already in order.
// We can reduce the running time to be linear for arrays that 
// already in order by adding a test to skip call to merge() 
// if a[mid] is less than or equal to a[mid+1]. With this change,
// we still do all the recursive calls, but the running time for
// any sorted subarray is linear.

// 3. Eliminate the copy to the auxiliary array.
// It is possible to eliminate the time (but not the space) taken
// to copy to the auxiliary array used for merging. To do so, we
// use two invocations of the sort method, one that takes its input
// from the given array and puts the sorted output in the auxiliary
// array; the other takes its input from the auxiliary array and
// puts the sorted output in the given array. With this approach, in
// a bit of mindbending recursive trickery, we can arrange the
// recursive calls such that the computation switchs the roles of
// the input array and the auxiliary at each level.

const CUTOFF int = 14 // cutoff to insertion sort

func MergesortX2(x Comparable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String() // sorting.IntCompSlice
	// t := (reflect.TypeOf(x)).Name() // IntCompSlice
	// log.Println(t)

	// convert type
	switch t {
	case "sorting.IntCompSlice":
		a := x.(IntCompSlice)
		aux := make(IntCompSlice, n)
		sortIntsX2(a, aux, 0, n-1)
	case "sorting.Float64CompSlice":
		a := x.(Float64CompSlice)
		aux := make(Float64CompSlice, n)
		sortFloat64sX2(a, aux, 0, n-1)
	case "sorting.StringCompSlice":
		a := x.(StringCompSlice)
		aux := make(StringCompSlice, n)
		sortStringsX2(a, aux, 0, n-1)
	}

}

// mergesort x[lo..hi] using auxiliary array aux[lo..hi]
func sortIntsX2(x, aux IntCompSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortIntsX2(x, aux, lo, mid)
	sortIntsX2(x, aux, mid+1, hi)

	// improvement 2: Test whether array is already in order.
	// x[mid+1]>=x[mid]
	if !x.Less(mid+1, mid) {
		return
	}
	
	mergeIntsX2(x, aux, lo, mid, hi)
}
func sortFloat64sX2(x, aux Float64CompSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortFloat64sX2(x, aux, lo, mid)
	sortFloat64sX2(x, aux, mid+1, hi)

	// improvement 2: Test whether array is already in order.
	// x[mid+1]>=x[mid]
	if !x.Less(mid+1, mid) {
		return
	}

	mergeFloat64sX2(x, aux, lo, mid, hi)
}
func sortStringsX2(x, aux StringCompSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortStringsX2(x, aux, lo, mid)
	sortStringsX2(x, aux, mid+1, hi)

	// improvement 2: Test whether array is already in order.
	// x[mid+1]>=x[mid]
	if !x.Less(mid+1, mid) {
		return
	}

	mergeStringsX2(x, aux, lo, mid, hi)
}

// stably merge x[lo..mid] with x[mid+1..hi] using aux[lo..hi]
func mergeIntsX2(x, aux IntCompSlice, lo, mid, hi int) {
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
func mergeFloat64sX2(x, aux Float64CompSlice, lo, mid, hi int) {
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
func mergeStringsX2(x, aux StringCompSlice, lo, mid, hi int) {
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

// insertion srot
func insertionSort(x Comparable, lo, hi int) {
	for i := lo; i <=hi; i++ {
		for j := i; j > 0 && x.Less(j, j-1); j-- {
			x.Swap(j, j-1)
		}
	}
}

type MergeX2 struct{}

// Implements Sorter
func (s MergeX2) SortInts(x []int) {
	MergesortX2(IntCompSlice(x))
}
func (s MergeX2) SortFloat64s(x []float64) {
	MergesortX2(Float64CompSlice(x))
}
func (s MergeX2) SortStrings(x []string) {
	MergesortX2(StringCompSlice(x))
}
