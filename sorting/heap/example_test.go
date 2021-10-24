package heap_test

import (
	"fmt"

	. "github.com/youngzhu/algs4-go/sorting/heap"
)

var ints = []int{2, 4, 7, 6, 8, 11, 17, 15}

func ExampleMinHeap() {
	heap := NewMinHeap()
	for _, v := range ints {
		heap.Insert(IntItem(v))
	}
	fmt.Println(heap.GetItems())

	heap.Insert(IntItem(3))
	fmt.Println(heap.GetItems())

	// Output:
	// [2 4 7 6 8 11 17 15]
	// [2 3 7 4 8 11 17 15 6]

}
