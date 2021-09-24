package selection_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/sorting"
)

func ExampleInts() {
	ints := []int{5, 4, 5, 3, 1, 2}
	sorter.Sort(sorting.IntCompSlice(ints))
	fmt.Println(ints)
	// Output:
	// [1 2 3 4 5 5]
}
