package sorting

import "reflect"

// Bottom-up mergesort
// Even though we are thinking in terms of merging together two large subarrays,
// the fact is that most merges are merging together tiny subarrays. Another
// way to implement mergesort is to organize so that wo do all the merges of
// tiny arrays on one pass, then do a second pass to merge those arrays in pairs,
// and so forth, continuing util we do a merge that encompasses the whole array.
// This method requires less code than the standard recursive implementation.
// We start by doing a pass of 1-by-1 merges (considering individual items as
// subarrays of size 1), then a pass of 2-by-2 merges (merge subarrays of size 2
// to make subarrays of size 4), then 4-by-4 merges, and so forth.

func MergesortBU(x Sortable) {

	t := (reflect.TypeOf(x)).String() // sorting.IntSortSlice
	// t := (reflect.TypeOf(x)).Name() // IntSortSlice

	// convert type
	switch t {
	case "sorting.IntSortSlice":
		a := x.(IntSortSlice)
		sortIntsBU(a)
	case "sorting.Float64SortSlice":
		a := x.(Float64SortSlice)
		sortFloat64sBU(a)
	case "sorting.StringSortSlice":
		a := x.(StringSortSlice)
		sortStringsBU(a)
	}
}

func sortIntsBU(x IntSortSlice) {
	n := x.Len()
	aux := make(IntSortSlice, n)
	for len := 1; len < n; len *= 2 {
		for lo := 0; lo < n-len; lo += len * 2 {
			mid := lo + len - 1
			hi := min(lo+len*2-1, n-1)
			mergeIntsBU(x, aux, lo, mid, hi)
		}
	}
}

func mergeIntsBU(x, aux IntSortSlice, lo, mid, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x
	i, j, k := lo, mid+1, lo

	for ; k <= hi; k++ {
		if i > mid {
			break
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

func sortFloat64sBU(x Float64SortSlice) {
	n := x.Len()
	aux := make(Float64SortSlice, n)
	for len := 1; len < n; len *= 2 {
		for lo := 0; lo < n-len; lo += len * 2 {
			mid := lo + len - 1
			hi := min(lo+len*2-1, n-1)
			mergeFloat64sBU(x, aux, lo, mid, hi)
		}
	}
}

func mergeFloat64sBU(x, aux Float64SortSlice, lo, mid, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x
	i, j, k := lo, mid+1, lo

	for ; k <= hi; k++ {
		if i > mid {
			break
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

func sortStringsBU(x StringSortSlice) {
	n := x.Len()
	aux := make(StringSortSlice, n)
	for len := 1; len < n; len *= 2 {
		for lo := 0; lo < n-len; lo += len * 2 {
			mid := lo + len - 1
			hi := min(lo+len*2-1, n-1)
			mergeStringsBU(x, aux, lo, mid, hi)
		}
	}
}

func mergeStringsBU(x, aux StringSortSlice, lo, mid, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x
	i, j, k := lo, mid+1, lo

	for ; k <= hi; k++ {
		if i > mid {
			break
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type MergeBU struct{}

func NewMergeBU() Sorter {
	return MergeBU{}
}

// Implements Sorter

func (s MergeBU) SortInts(x []int) {
	MergesortBU(IntSortSlice(x))
}
func (s MergeBU) SortFloat64s(x []float64) {
	MergesortBU(Float64SortSlice(x))
}
func (s MergeBU) SortStrings(x []string) {
	MergesortBU(StringSortSlice(x))
}
