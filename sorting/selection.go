package sorting

// Selection Sort
// First, find the smallest item int the slice, and swap it with the first entry
// Then, find the next smallest item and swap it with the second entry
// Continue in this way until the entire slice is sorted

// This method is called selection sort because it works by repeatedly selecting
// the smallest remaining items
func SelectionSort(x Sortable) {
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

type Selection struct{}

// Implements Sorter
func (s Selection) SortInts(x []int) {
	SelectionSort(IntCompSlice(x))
}
func (s Selection) SortFloat64s(x []float64) {
	SelectionSort(Float64CompSlice(x))
}
func (s Selection) SortStrings(x []string) {
	SelectionSort(StringCompSlice(x))
}
