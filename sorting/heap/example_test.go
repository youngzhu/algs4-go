package heap_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/sorting/heap"
)

var ints = []int{2, 4, 7, 6, 8, 11, 17, 15}

func ExampleMinHeap_Insert() {
	mh := heap.NewMinHeap()
	for _, v := range ints {
		mh.Insert(heap.IntItem(v))
	}
	fmt.Println(mh.GetItems())

	mh.Insert(heap.IntItem(3))
	fmt.Println(mh.GetItems())

	// Output:
	// [2 4 7 6 8 11 17 15]
	// [2 3 7 4 8 11 17 15 6]
}

func ExampleMinHeap_Remove() {
	mh := heap.NewMinHeap()
	for _, v := range ints {
		mh.Insert(heap.IntItem(v))
	}
	fmt.Println(mh.GetItems())

	mh.Remove()
	fmt.Println(mh.GetItems())

	// Output:
	// [2 4 7 6 8 11 17 15]
	// [4 6 7 15 8 11 17]
}

func ExampleMaxHeap_Insert() {
	mh := heap.NewMaxHeap()
	for _, v := range ints {
		mh.Insert(heap.IntItem(v))
	}
	fmt.Println(mh.GetItems())

	mh.Insert(heap.IntItem(3))
	fmt.Println(mh.GetItems())

	// Output:
	// [17 15 11 7 6 4 8 2]
	// [17 15 11 7 6 4 8 2 3]
}

func ExampleMaxHeap_Remove() {
	mh := heap.NewMaxHeap()
	for _, v := range ints {
		mh.Insert(heap.IntItem(v))
	}
	fmt.Println(mh.GetItems())

	mh.Remove()
	fmt.Println(mh.GetItems())

	// Output:
	// [17 15 11 7 6 4 8 2]
	// [15 7 11 2 6 4 8]
}
