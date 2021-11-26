package util_test

import (
	"fmt"

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

func ExampleIn_ReadLine() {
	in := util.NewIn("testdata/in.txt")
	for in.HasNext() {
		fmt.Println(in.ReadLine())
	}
	// Output:
	// hello Gopher
	// wating for  you
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

func ExampleIn_ReadInt_part() {
	in := util.NewInReadWords("testdata/ints.txt")
	
	fmt.Println(in.ReadInt())
	fmt.Println(in.ReadInt())
	fmt.Println(in.ReadInt())
	
	// Output:
	// 12
	// 3
	// -1
}

func ExampleIn_ReadAllStrings() {
	in := util.NewInReadWords("testdata/in.txt")
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [hello Gopher wating for you]
}

func ExampleIn_ReadAllInts() {
	in := util.NewInReadWords("testdata/ints.txt")
	s := in.ReadAllInts()
	fmt.Println(s)
	// Output:
	// [12 3 -1 5 6]
}

func ExampleIn_ReadAllStrings_http() {
	const url = "https://algs4.cs.princeton.edu/24pq/tiny.txt"
	in := util.NewInReadWords(url)
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [S O R T E X A M P L E]
}
