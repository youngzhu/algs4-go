package sorting

// Optimized Mergesort

import "reflect"

// Improvements
// We can cut the running time of mergesort substantially with some
// carefully considered modifications to the implementation.

// 1. Use insertion sort for small subarrays.
// We can improve most recursive algorithms by handling small cases
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
// recursive calls such that the computation switches the roles of
// the input array and the auxiliary at each level.

const CUTOFF int = 14 // cutoff to insertion sort

func MergesortX(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String() // sorting.IntSortSlice
	// t := (reflect.TypeOf(x)).Name() // IntSortSlice
	// log.Println(t)

	// convert type
	switch t {
	case "sorting.IntSortSlice":
		a := x.(IntSortSlice)
		aux := make(IntSortSlice, n)
		copy(aux, a)
		sortIntsX(aux, a, 0, n-1)
	case "sorting.Float64SortSlice":
		a := x.(Float64SortSlice)
		aux := make(Float64SortSlice, n)
		copy(aux, a)
		sortFloat64sX(aux, a, 0, n-1)
	case "sorting.StringSortSlice":
		a := x.(StringSortSlice)
		aux := make(StringSortSlice, n)
		copy(aux, a)
		sortStringsX(aux, a, 0, n-1)
	}

}

func sortIntsX(src, dst IntSortSlice, lo, hi int) {
	// improvment 1. Use insertion sort for small subarrays.
	if hi <= lo+CUTOFF {
		insertionSort(dst, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	sortIntsX(dst, src, lo, mid)
	sortIntsX(dst, src, mid+1, hi)

	// improvement 2: Test whether array is already in order.
	// x[mid+1]>=x[mid]
	if !src.Less(mid+1, mid) {
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}

	mergeIntsX(src, dst, lo, mid, hi)
}
func sortFloat64sX(src, dst Float64SortSlice, lo, hi int) {
	// improvment 1. Use insertion sort for small subarrays.
	if hi <= lo+CUTOFF {
		insertionSort(dst, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	sortFloat64sX(dst, src, lo, mid)
	sortFloat64sX(dst, src, mid+1, hi)

	// improvement 2: Test whether array is already in order.
	// x[mid+1]>=x[mid]
	if !src.Less(mid+1, mid) {
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}

	mergeFloat64sX(src, dst, lo, mid, hi)
}
func sortStringsX(src, dst StringSortSlice, lo, hi int) {
	// improvement 1. Use insertion sort for small subarrays.
	if hi <= lo+CUTOFF {
		insertionSort(dst, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	sortStringsX(dst, src, lo, mid)
	sortStringsX(dst, src, mid+1, hi)

	// improvement 2: Test whether array is already in order.
	// x[mid+1]>=x[mid]
	if !src.Less(mid+1, mid) {
		copy(dst[lo:hi+1], src[lo:hi+1])
		return
	}

	mergeStringsX(src, dst, lo, mid, hi)
}

func mergeIntsX(src, dst IntSortSlice, lo, mid, hi int) {

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
		} else if j > hi {
			dst[k] = src[i]
			i++
		} else if src.Less(j, i) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}
func mergeFloat64sX(src, dst Float64SortSlice, lo, mid, hi int) {

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
		} else if j > hi {
			dst[k] = src[i]
			i++
		} else if src.Less(j, i) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}
func mergeStringsX(src, dst StringSortSlice, lo, mid, hi int) {

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
		} else if j > hi {
			dst[k] = src[i]
			i++
		} else if src.Less(j, i) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

// insertion sort
func insertionSort(x Sortable, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && x.Less(j, j-1); j-- {
			x.Swap(j, j-1)
		}
	}
}

type MergeX struct{}

func NewMergeX() Sorter {
	return MergeX{}
}

// Implements Sorter

func (s MergeX) SortInts(x []int) {
	MergesortX(IntSortSlice(x))
}
func (s MergeX) SortFloat64s(x []float64) {
	MergesortX(Float64SortSlice(x))
}
func (s MergeX) SortStrings(x []string) {
	MergesortX(StringSortSlice(x))
}
