package strings_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/strings"
	"github.com/youngzhu/algs4-go/util"
)

func ExampleLSDSort() {
	in := util.NewInReadWords("testdata/words3.txt")
	a := in.ReadAllStrings()
	strings.LSDSort(a)

	fmt.Println(a)
}

func ExampleLSDSortInts() {
	ints := []int{1, 22, 55, 3, 2, 44, 0}
	strings.LSDSortInts(ints[:])

	fmt.Println(ints)

	// --
	// Output:
	//
}

func ExampleMSDSort() {
	in := util.NewInReadWords("testdata/shells.txt")
	a := in.ReadAllStrings()
	strings.MSDSort(a)

	for _, s := range a {
		fmt.Println(s)
	}

	// Output:
	// are
	// by
	// sea
	// seashells
	// seashells
	// sells
	// sells
	// she
	// she
	// shells
	// shore
	// surely
	// the
	// the
}

func ExampleQuicksort() {
	in := util.NewInReadWords("testdata/shells.txt")
	a := in.ReadAllStrings()
	strings.Quicksort(a)

	for _, s := range a {
		fmt.Println(s)
	}

	// Output:
	// are
	// by
	// sea
	// seashells
	// seashells
	// sells
	// sells
	// she
	// she
	// shells
	// shore
	// surely
	// the
	// the
}
