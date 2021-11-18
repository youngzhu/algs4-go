package util_test

import (
	"fmt"
	"math"

	"github.com/youngzhu/algs4-go/util"
)

func ExampleIn_ReadString() {
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

func ExampleIn_ReadString_gz() {
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

func ExampleIn_ReadInt() {
	in := util.NewInReadWords("testdata/ints.txt")
	for !in.IsEmpty() {
		fmt.Println(in.ReadInt())
	}
	// Output:
	// 12
	// 3
	// -1
	// 5
	// 6
}

func ExampleIn_ReadAllStrings() {
	in := util.NewInReadWords("testdata/in.txt")
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [hello Gopher wating for you]
}

func ExampleIn_ReadAllStrings_http() {
	const url = "https://algs4.cs.princeton.edu/24pq/tiny.txt"
	in := util.NewInReadWords(url)
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [S O R T E X A M P L E]
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