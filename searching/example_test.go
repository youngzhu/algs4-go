package searching_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/searching"
)

var tinyST = []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}

func ExampleSequentialSearchST() {
	st := searching.NewSequentialSearchST()
	for i, v := range tinyST {
		st.Put(searching.StringKey(v), i)
	}

	for _, k := range st.Keys() {
		fmt.Println(k, st.Get(k))
	}

	// Output:
	// L 11
	// P 10
	// M 9
	// X 7
	// H 5
	// C 4
	// R 3
	// A 8
	// E 12
	// S 0
}

func ExampleBinarySearchST() {
	st := searching.NewBinarySearchST()
	for i, v := range tinyST {
		st.Put(searching.StringKey(v), i)
	}

	for _, k := range st.Keys() {
		fmt.Println(k, st.Get(k))
	}

	// Output:
	// 	A 8
	// C 4
	// E 12
	// H 5
	// L 11
	// M 9
	// P 10
	// R 3
	// S 0
	// X 7
}

func ExampleBST() {
	st := searching.NewBST()
	for i, v := range tinyST {
		st.Put(searching.StringKey(v), i)
	}

	for _, k := range st.Keys() {
		fmt.Println(k, st.Get(k))
	}

	// Output:
	// A 8
	// C 4
	// E 12
	// H 5
	// L 11
	// M 9
	// P 10
	// R 3
	// S 0
	// X 7
}

func ExampleRedBlackBST() {
	st := searching.NewRedBlackBST()
	for i, v := range tinyST {
		st.Put(searching.StringKey(v), i)
	}

	for _, k := range st.Keys() {
		fmt.Println(k, st.Get(k))
	}

	// Output:
	// A 8
	// C 4
	// E 12
	// H 5
	// L 11
	// M 9
	// P 10
	// R 3
	// S 0
	// X 7
}