package sorting

// Optimized Mergesort

import "reflect"

// Improvements
// We can cut the running time of mergesort substantially with some
// carefully considered modifications to the implementation.

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

func MergesortX3(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String() // sorting.IntCompSlice
	// t := (reflect.TypeOf(x)).Name() // IntCompSlice
	// log.Println(t)

	// convert type
	switch t {
	case "sorting.IntCompSlice":
		a := x.(IntCompSlice)
		aux := make(IntCompSlice, n)
		copy(aux, a)
		sortIntsX3(aux, a, 0, n-1)
	case "sorting.Float64CompSlice":
		a := x.(Float64CompSlice)
		aux := make(Float64CompSlice, n)
		copy(aux, a)
		sortFloat64sX3(aux, a, 0, n-1)
	case "sorting.StringCompSlice":
		a := x.(StringCompSlice)
		aux := make(StringCompSlice, n)
		copy(aux, a)
		sortStringsX3(aux, a, 0, n-1)
	}

}

func sortIntsX3(src, dst IntCompSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortIntsX3(dst, src, lo, mid)
	sortIntsX3(dst, src, mid+1, hi)

	mergeIntsX3(src, dst, lo, mid, hi)
}
func sortFloat64sX3(src, dst Float64CompSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortFloat64sX3(dst, src, lo, mid)
	sortFloat64sX3(dst, src, mid+1, hi)

	mergeFloat64sX3(src, dst, lo, mid, hi)
}
func sortStringsX3(src, dst StringCompSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortStringsX3(dst, src, lo, mid)
	sortStringsX3(dst, src, mid+1, hi)

	mergeStringsX3(src, dst, lo, mid, hi)
}

func mergeIntsX3(src, dst IntCompSlice, lo, mid, hi int) {

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
func mergeFloat64sX3(src, dst Float64CompSlice, lo, mid, hi int) {

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
func mergeStringsX3(src, dst StringCompSlice, lo, mid, hi int) {

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

type MergeX3 struct{}

// Implements Sorter
func (s MergeX3) SortInts(x []int) {
	MergesortX3(IntCompSlice(x))
}
func (s MergeX3) SortFloat64s(x []float64) {
	MergesortX3(Float64CompSlice(x))
}
func (s MergeX3) SortStrings(x []string) {
	MergesortX3(StringCompSlice(x))
}
