package pq_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/sorting/pq"
)

var a = [...]string{"P", "Q", "E", "-", "X", "A", "M", "-", "P", "L", "E", "-"}

func ExampleMaxPQ() {
	p := pq.NewMaxPQ()

	for _, item := range a {
		if item != "-" {
			p.Insert(pq.StringItem(item))
		} else if !p.IsEmpty() {
			fmt.Print(p.Delete(), " ")
		}

	}

	fmt.Printf("(%d left on pq)", p.Size())

	// Output: Q X P (6 left on pq)
}

func ExampleMinPQ() {
	p := pq.NewMinPQ()

	for _, item := range a {
		if item != "-" {
			p.Insert(pq.StringItem(item))
		} else if !p.IsEmpty() {
			fmt.Print(p.Delete(), " ")
		}

	}

	fmt.Printf("(%d left on pq)", p.Size())

	// Output: E A E (6 left on pq)
}

var ints = []int{2, 4, 7, 6, 8, 11, 17, 15}

func ExampleMinPQ_Insert() {
	q := pq.NewMinPQ()
	for _, v := range ints {
		q.Insert(pq.IntItem(v))
	}
	fmt.Println(q.GetItems())

	q.Insert(pq.IntItem(3))
	fmt.Println(q.GetItems())

	// Output:
	// [2 4 7 6 8 11 17 15]
	// [2 3 7 4 8 11 17 15 6]
}

func ExampleMinPQ_Delete() {
	q := pq.NewMinPQ()
	for _, v := range ints {
		q.Insert(pq.IntItem(v))
	}
	fmt.Println(q.GetItems())

	q.Delete()
	fmt.Println(q.GetItems())

	// Output:
	// [2 4 7 6 8 11 17 15]
	// [4 6 7 15 8 11 17]
}

func ExampleMaxPQ_Insert() {
	q := pq.NewMaxPQ()
	for _, v := range ints {
		q.Insert(pq.IntItem(v))
	}
	fmt.Println(q.GetItems())

	q.Insert(pq.IntItem(3))
	fmt.Println(q.GetItems())

	// Output:
	// [17 15 11 7 6 4 8 2]
	// [17 15 11 7 6 4 8 2 3]
}

func ExampleMaxPQ_Delete() {
	q := pq.NewMaxPQ()
	for _, v := range ints {
		q.Insert(pq.IntItem(v))
	}
	fmt.Println(q.GetItems())

	q.Delete()
	fmt.Println(q.GetItems())

	// Output:
	// [17 15 11 7 6 4 8 2]
	// [15 7 11 2 6 4 8]
}
