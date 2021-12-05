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

var trieST *strings.TrieST

func initTrieST() {
	// build ST from file
	trieST = strings.NewTrieST()

	in := util.NewInReadWords("testdata/shellsST.txt")
	a := in.ReadAllStrings()

	for i, s := range a {
		trieST.Put(s, i)
	}
}

func ExampleTrieST_Keys() {
	if trieST == nil {
		initTrieST()
	}

	fmt.Println("keys(\"\"):")
	for _, v := range trieST.Keys() {
		fmt.Printf("%s %d\n", v, trieST.Get(v))
	}

	// Output:
	// keys(""):
	// by 4
	// sea 6
	// sells 1
	// she 0
	// shells 3
	// shore 7
	// the 5
}

func ExampleTrieST_LongestPrefixOf() {
	if trieST == nil {
		initTrieST()
	}

	prefix := "shellsort"
	fmt.Printf("LongestPrefixOf(\"%s\"):\n", prefix)
	fmt.Println(trieST.LongestPrefixOf(prefix))
	fmt.Println()

	prefix = "quicksort"
	fmt.Printf("LongestPrefixOf(\"%s\"):\n", prefix)
	fmt.Println(trieST.LongestPrefixOf(prefix))
	fmt.Println()

	// Output:
	// LongestPrefixOf("shellsort"):
	// shells
	//
	// LongestPrefixOf("quicksort"):
	//
	//
}

func ExampleTrieST_KeysWithPrefix() {
	if trieST == nil {
		initTrieST()
	}

	prefix := "shor"
	fmt.Printf("KeysWithPrefix(\"%s\"):\n", prefix)
	for _, s := range trieST.KeysWithPrefix(prefix) {
		fmt.Println(s)
	}

	// Output:
	// KeysWithPrefix("shor"):
	// shore
}

func ExampleTrieST_KeysThatMatch() {
	if trieST == nil {
		initTrieST()
	}

	pattern := "shor"
	fmt.Printf("KeysThatMatch(\"%s\"):\n", pattern)
	for _, s := range trieST.KeysThatMatch(pattern) {
		fmt.Println(s)
	}

	// Output:
	// KeysThatMatch("shor"):
	//
}

func ExampleTernarySearchTrie() {
	tst := strings.NewTernarySearchTrie()

	in := util.NewInReadWords("testdata/shellsST.txt")
	a := in.ReadAllStrings()

	for i, s := range a {
		tst.Put(s, i)
	}

	fmt.Println("Keys:")
	for _, v := range tst.Keys() {
		fmt.Printf("%s %d\n", v, tst.Get(v))
	}
	fmt.Println()

	prefix := "shellsort"
	fmt.Printf("LongestPrefixOf(\"%s\"):\n", prefix)
	fmt.Println(tst.LongestPrefixOf(prefix))
	fmt.Println()

	prefix = "shor"
	fmt.Printf("KeysWithPrefix(\"%s\"):\n", prefix)
	for _, s := range tst.KeysWithPrefix(prefix) {
		fmt.Println(s)
	}
	fmt.Println()

	pattern := ".he.l."
	fmt.Printf("KeysThatMatch(\"%s\"):\n", pattern)
	for _, s := range tst.KeysThatMatch(pattern) {
		fmt.Println(s)
	}

	// Output:
	// Keys:
	// by 4
	// sea 6
	// sells 1
	// she 0
	// shells 3
	// shore 7
	// the 5
	//
	// LongestPrefixOf("shellsort"):
	// shells
	//
	// KeysWithPrefix("shor"):
	// shore
	//
	// KeysThatMatch(".he.l."):
	// shells
}
