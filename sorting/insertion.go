package sorting

// InsertionSort
// The algorithm that people often use to sort bridge hands is to consider
// the cards one at a time, inserting each into its peoper place among those
// already considered (keeping them sorted).
// In a computer implementation, we need to make space for the current item by moving
// larger items one position to the right, before inserting the current item into the
// vacated position
func InsertionSort(x Sortable) {
	n := x.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && x.Less(j, j-1); j-- {
			x.Swap(j, j-1)
		}
	}
}

type Insertion struct{}

func NewInsertion() Sorter {
	return Insertion{}
}

func (s Insertion) SortInts(x []int) {
	InsertionSort(IntSortSlice(x))
}
func (s Insertion) SortFloat64s(x []float64) {
	InsertionSort(Float64SortSlice(x))
}
func (s Insertion) SortStrings(x []string) {
	InsertionSort(StringSortSlice(x))
}
