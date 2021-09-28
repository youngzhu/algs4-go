package sorting_test

import (
	"fmt"
	"math"

	. "github.com/youngzhu/algs4-go/sorting"
	"github.com/youngzhu/algs4-go/util"
)

var (
	selectionAlg Sorter = Selection{}
	insertionAlg Sorter = Insertion{}
)

// ints
func ExampleSelectionInts() {
	ints := []int{5, 4, 5, 3, 1, 2}
	selectionAlg.SortInts(ints)
	fmt.Println(ints)
	// Output:
	// [1 2 3 4 5 5]
}

func ExampleInsertionInts() {
	ints := []int{5, 4, 5, 3, 1, 2}
	insertionAlg.SortInts(ints)
	fmt.Println(ints)
	// Output:
	// [1 2 3 4 5 5]
}

// float64s
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
func ExampleInsertionFloat64s() {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	insertionAlg.SortFloat64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	insertionAlg.SortFloat64s(s)
	fmt.Println(s)

	// Output: 
	// [-3.8 -1.3 0.7 2.6 5.2]
	// [NaN -Inf 0 +Inf]
}

// strings
func ExampleSelectionStrings() {
	s := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"} 
	selectionAlg.SortStrings(s)
	fmt.Println(s)

	// Output: 
	// [A E E L M O P R S T X]
}
func ExampleInsertionStrings() {
	s := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"} 
	insertionAlg.SortStrings(s)
	fmt.Println(s)

	// Output: 
	// [A E E L M O P R S T X]
}

// from file
func ExampleSelectionFromFile() {
	s := util.ReadAllStrings("testdata/words3.txt") 
	selectionAlg.SortStrings(s)
	fmt.Println(s)

	// Output: 
	// [all bad bed bug dad dim dug egg fee few for gig hut ilk jam jay jot joy men nob now owl rap sky sob tag tap tar tip wad was wee yes yet zoo]
}
func ExampleInsertionFromFile() {
	s := util.ReadAllStrings("testdata/words3.txt") 
	insertionAlg.SortStrings(s)
	fmt.Println(s)

	// Output: 
	// [all bad bed bug dad dim dug egg fee few for gig hut ilk jam jay jot joy men nob now owl rap sky sob tag tap tar tip wad was wee yes yet zoo]
}