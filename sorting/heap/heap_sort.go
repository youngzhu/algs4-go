package heap

import (
	"github.com/youngzhu/algs4-go/sorting"
)

// Heapsort.
// We can use any priority queue to develop a sorting method. We insert all
// the keys to be sorted into a minimum-oriented priority queue, then repeatedly
// use remove the minimum to remove them all in order. When using a heap for
// the priority queue, we obtain heapsort.

// Focusing on the task of sorting, we abandon the notion of hiding the heap
// representation of the priority queue and use swim() and sink() directly.
// Doing so allows us to sort an array without needing any extra space, by
// maintaining the heap within the array to be sorted. Heapsort breaks into two
// phases: heap construction, where we reorganzie the original array into a
// heap, and the sortdown, where we pull the items out of the heap in
// decreasing order to build the sorted result.
//
// Heap construction. We can accomplish this task in time proportional to NlgN,
// by proceeding from left to right through the array, using swim() to ensure
// that the entries to the left of the scanning pointer make up a heap-ordered
// complete tree, like successive priority queue insertions. A clever method that
// is much more efficient is to proceed from right to left, using sink() to make
// subheaps as we go. Every position in the array is the root of a small subheap;
// sink() works on such subheaps, as well. If the two children of a node are
// heaps, then calling sink() on that node makes the subtree rooted there a heap.
//
// Sortdown. Most of the work during heapsort is done during the second phase,
// where we remove the largest remaining items from the heap and put it into the
// array position vacated as the heap shrinks.

// sink-based heapsort
func Heapsort(x sorting.Sortable) {
	n := x.Len()

	// heapify phase
	for k := n / 2; k >= 1; k-- {
		sink(x, k, n)
	}

	// sortdown phase
	i := n
	for i > 1 {
		swap(x, 1, i)
		i--
		sink(x, 1, i)
	}
}

func sink(x sorting.Sortable, k, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && less(x, j, j+1) {
			j++
		}
		if !less(x, k, j) {
			break
		}
		swap(x, k, j)
		k = j
	}
}

// Helper functions for comparisons and swaps.
// Indices are "off-by-one" to support 1-based indexing.
func less(x sorting.Sortable, i, j int) bool {
	return x.Less(i-1, j-1)
}
func swap(x sorting.Sortable, i, j int) {
	x.Swap(i-1, j-1)
}

type Heapsorter struct{}

// Implements Sorter
func (s Heapsorter) SortInts(x []int) {
	Heapsort(sorting.IntSortSlice(x))
}
func (s Heapsorter) SortFloat64s(x []float64) {
	Heapsort(sorting.Float64SortSlice(x))
}
func (s Heapsorter) SortStrings(x []string) {
	Heapsort(sorting.StringSortSlice(x))
}
