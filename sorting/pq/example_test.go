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