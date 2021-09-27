package util_test

import (
	"fmt"

	. "github.com/youngzhu/algs4-go/util"
)

func ExampleReadAllStrings() {
	s := ReadAllStrings("testdata/intest.txt")
	fmt.Println(s)
	// Output:
	// [hello Gopher wating for you]
}