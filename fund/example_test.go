package fund_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/util"
)

func ExampleQueue() {
	queue := fund.NewQueue()

	in := util.NewInReadWords("testdata/tobe.txt")

	for !in.IsEmpty() {
		item := in.ReadString()
		if item != "-" {
			queue.Enqueue(item)
		} else {
			fmt.Print(queue.Dequeue(), " ")
		}
	}

	fmt.Print("(", queue.Size(), " left on queue)")

	// Output:
	// to be or not to be (2 left on queue)
}

func ExampleQueue_iterator() {
	queue := fund.NewQueue()

	items := []string{"to", "be", "or", "not", "to", "be"}

	for _, v := range items {
		queue.Enqueue(v)
	}

	for _, v := range queue.Iterator() {
		fmt.Print(v, " ")
	}
	
	// Output:
	// to be or not to be
}