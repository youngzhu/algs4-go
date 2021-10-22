package pq_test

import (
	"fmt"

	. "github.com/youngzhu/algs4-go/sorting/pq"
)

var a = [...]string{"P", "Q", "E", "-", "X", "A", "M", "-", "P", "L", "E", "-"}


func ExampleMaxPQ() {
	pq := NewMaxPQ()

	for _, item := range a {
		if item != "-" {
			pq.Insert(StringItem(item))
		} else if !pq.IsEmpty() {
			fmt.Print(pq.Delete(), " ")
		}
		
	}

	fmt.Printf("(%d left on pq)", pq.Size())

	// Output: Q X P (6 left on pq)
}