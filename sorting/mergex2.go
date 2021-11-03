package sorting

// Optimized Mergesort

import "reflect"

// Improvements
// We can cut the running time of mergesort substantially with some
// carefully considered modifications to the implementation.

// 2. Test whether array is already in order.
// We can reduce the running time to be linear for arrays that
// already in order by adding a test to skip call to merge()
// if a[mid] is less than or equal to a[mid+1]. With this change,
// we still do all the recursive calls, but the running time for
// any sorted subarray is linear.

func MergesortX2(x Sortable) {
	n := x.Len()

	t := (reflect.TypeOf(x)).String() // sorting.IntSortSlice
	// t := (reflect.TypeOf(x)).Name() // IntSortSlice
	// log.Println(t)

	// convert type
	switch t {
	case "sorting.IntSortSlice":
		a := x.(IntSortSlice)
		aux := make(IntSortSlice, n)
		sortIntsX2(a, aux, 0, n-1)
	case "sorting.Float64SortSlice":
		a := x.(Float64SortSlice)
		aux := make(Float64SortSlice, n)
		sortFloat64sX2(a, aux, 0, n-1)
	case "sorting.StringSortSlice":
		a := x.(StringSortSlice)
		aux := make(StringSortSlice, n)
		sortStringsX2(a, aux, 0, n-1)
	}

}

// mergesort x[lo..hi] using auxiliary array aux[lo..hi]
func sortIntsX2(x, aux IntSortSlice, lo, hi int) {
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
func sortFloat64sX2(x, aux Float64SortSlice, lo, hi int) {
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
func sortStringsX2(x, aux StringSortSlice, lo, hi int) {
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
func mergeIntsX2(x, aux IntSortSlice, lo, mid, hi int) {
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
func mergeFloat64sX2(x, aux Float64SortSlice, lo, mid, hi int) {
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
func mergeStringsX2(x, aux StringSortSlice, lo, mid, hi int) {
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

type MergeX2 struct{}

func NewMergeX2() Sorter {
	return MergeX2{}
}

// Implements Sorter
func (s MergeX2) SortInts(x []int) {
	MergesortX2(IntSortSlice(x))
}
func (s MergeX2) SortFloat64s(x []float64) {
	MergesortX2(Float64SortSlice(x))
}
func (s MergeX2) SortStrings(x []string) {
	MergesortX2(StringSortSlice(x))
}
