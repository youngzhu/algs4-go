package sorting_test

import (
	"fmt"

	. "github.com/youngzhu/algs4-go/sorting"
	"github.com/youngzhu/algs4-go/sorting/selection"
)

var (
	selectionAlg Sorter = selection.Selection{}
)

func ExampleSelectionInts() {
	ints := []int{5, 4, 5, 3, 1, 2}
	selectionAlg.SortInts(ints)
	fmt.Println(ints)
	// Output:
	// [1 2 3 4 5 5]
}
