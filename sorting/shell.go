package sorting

// Shellsort
// Shellsort is simple extension of insertion sort that gains speed by allowing exchanges
// of entries that are far apart, to produce partially sorted arrays that can be efficientlly
// sorted, eventually by insertion sort.
// The idea is to rearrange the array to give it the property that taking every hth entry
// (starting anywhere) yields a sorted sequence. Such an array is said to be h-sorted.
// By h-sorting for some large values of h, we can move entries in the array long distances
// and thus make it easier to h-sort for smaller values of h. Using such a procedure for any
// increment sequence of values of h that ends in 1 will produce a sorted array: that is shellsort.
func Shellsort(x Sortable) {
	n := x.Len()

	// 3h+1 increment sequence: 1, 4, 13, 40, 121...
	h := 1
	for {
		if h >= n/3 {
			break
		}
		h = 3*h + 1
	}

	for {
		if h < 1 {
			break
		}
		// h-sort the array
		for i := h; i < n; i++ {
			for j := i; j >= h && x.Less(j, j-h); j -= h {
				x.Swap(j, j-h)
			}
		}
		h /= 3
	}

}

type Shell struct{}

func NewShell() Sorter {
	return Shell{}
}

// Implements Sorter
func (s Shell) SortInts(x []int) {
	Shellsort(IntSortSlice(x))
}
func (s Shell) SortFloat64s(x []float64) {
	Shellsort(Float64SortSlice(x))
}
func (s Shell) SortStrings(x []string) {
	Shellsort(StringSortSlice(x))
}
