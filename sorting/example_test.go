package sorting_test

import (
	"fmt"
	"math"

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

func ExampleSelectionFloat64s() {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	selectionAlg.SortFloat64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	selectionAlg.SortFloat64s(s)
	fmt.Println(s)

	// Output: 
	// [-3.8 -1.3 0.7 2.6 5.2]
	// [NaN -Inf 0 +Inf]
}

func ExampleSelectionStrings() {
	s := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"} 
	selectionAlg.SortStrings(s)
	fmt.Println(s)

	// Output: 
	// [A E E L M O P R S T X]
}