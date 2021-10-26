package util_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/util"
)

func ExampleReadString() {
	in := util.NewInReadWords("testdata/intest.txt")
	for !in.IsEmpty() {
		fmt.Println(in.ReadString())
	}
	// Output:
	// hello
	// Gopher
	// wating
	// for
	// you
}

func ExampleReadAllStrings() {
	in := util.NewInReadWords("testdata/intest.txt")
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [hello Gopher wating for you]
}