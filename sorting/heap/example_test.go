package heap_test

import (
	"fmt"

	. "github.com/youngzhu/algs4-go/sorting/heap"
)

var ints = []int{2, 4, 7, 6, 8, 11, 17, 15}

func ExampleMinHeap_Insert() {
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

func ExampleMinHeap_Remove() {
	heap := NewMinHeap()
	for _, v := range ints {
		heap.Insert(IntItem(v))
	}
	fmt.Println(heap.GetItems())

	heap.Remove()
	fmt.Println(heap.GetItems())

	// Output:
	// [2 4 7 6 8 11 17 15]
	// [4 6 7 15 8 11 17]

}
