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

func ExampleStack() {
	stack := fund.NewStack()

	in := util.NewInReadWords("testdata/tobe.txt")

	for !in.IsEmpty() {
		item := in.ReadString()
		if item != "-" {
			stack.Push(item)
		} else {
			fmt.Print(stack.Pop(), " ")
		}
	}

	fmt.Print("(", stack.Size(), " left on stack)")

	// Output:
	// to be not that or be (2 left on stack)
}

// Read a sequence of integers and print them in reverse order
func ExampleStack_reverse() {
	stack := fund.NewStack()

	ints := []int{1, 2, 3, 4, 5}

	for _, v := range ints {
		stack.Push(v)
	}

	for _, v := range stack.Iterator() {
		fmt.Print(v, " ")
	}

	// Output:
	// 5 4 3 2 1
}