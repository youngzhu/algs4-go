package util_test

import (
	"fmt"
	"math"

	"github.com/youngzhu/algs4-go/util"
)

func ExampleReadString() {
	in := util.NewInReadWords("testdata/in.txt")
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

func ExampleReadString_gz() {
	in := util.NewInReadWords("testdata/in.txt.gz")
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
	in := util.NewInReadWords("testdata/in.txt")
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [hello Gopher wating for you]
}

func ExampleStopwatch() {
	var n float64 = 1000
	timer := util.NewStopwatch()
	sum := 0.0
	for n > 0 {
		sum += math.Sqrt(n)
		n--
	}
	time := timer.ElapsedTime()
	fmt.Printf("%e (%.2f seconds)", sum, time)
}