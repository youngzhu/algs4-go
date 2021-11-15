package fund_test

import (
	"fmt"
	"math"

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

func ExampleQueue_Iterator() {
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

func ExampleBag() {
	bag := fund.NewBag()

	in := util.NewInReadWords("testdata/tobe.txt")

	for !in.IsEmpty() {
		item := in.ReadString()
		bag.Add(item)
	}

	fmt.Printf("size of bag: %d\n", bag.Size())

	for _, v := range bag.Iterator() {
		fmt.Print(v, " ")
	}

	// Output:
	// size of bag: 14
	// is - - - that - - be - to not or be to
}

// Read a sequence of numbers and computes their mean and standard deviation
func ExampleBag_stats() {
	bag := fund.NewBag()

	numbers := []float64{100, 99, 101, 120, 98, 107, 109, 81, 101, 90}

	for _, v := range numbers {
		bag.Add(v)
	}

	n := float64(bag.Size())

	// compute sample mean
	sum := 0.0
	for _, v := range bag.Iterator() {
		sum += v.(float64)
	}
	mean := sum / n
	fmt.Printf("Mean: %.2f\n", mean)

	// compute sample standard deviation
	sum = 0
	for _, v := range bag.Iterator() {
		t := v.(float64)
		sum += (t - mean) * (t - mean)
	}
	stdDev := math.Sqrt(sum / (n - 1))
	fmt.Printf("Std dev: %.2f\n", stdDev)

	// Output:
	// Mean: 100.60
	// Std dev: 10.51

}
