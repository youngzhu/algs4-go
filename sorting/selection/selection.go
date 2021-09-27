package selection

import (
	"github.com/youngzhu/algs4-go/sorting"
)

type Selection struct {}

// Implements Sorter
func (s Selection) SortInts(x []int) {
	Sort(sorting.IntCompSlice(x))
}
func (s Selection) SortFloat64s(x []float64) {
	Sort(sorting.Float64CompSlice(x))
}
func (s Selection) SortStrings(x []string) {
	Sort(sorting.StringCompSlice(x))
}

func Sort(x sorting.Comparable) {
	n := x.Len()
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if x.Less(j, min) {
				min = j
			}
		}
		x.Swap(i, min)
	}
}