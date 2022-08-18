package sort_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/strings/sort"
	"github.com/youngzhu/algs4-go/testutil"
)

func ExampleLSDSort() {
	in := testutil.NewInReadWords("testdata/words3.txt")
	a := in.ReadAllStrings()
	sort.LSDSort(a)

	fmt.Println(a)
}

func ExampleLSDSortInts() {
	ints := []int{1, 22, 55, 3, 2, 44, 0}
	sort.LSDSortInts(ints[:])

	fmt.Println(ints)

	// --
	// Output:
	//
}

func ExampleMSDSort() {
	in := testutil.NewInReadWords("testdata/shells.txt")
	a := in.ReadAllStrings()
	sort.MSDSort(a)

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
	in := testutil.NewInReadWords("testdata/shells.txt")
	a := in.ReadAllStrings()
	sort.Quicksort(a)

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
