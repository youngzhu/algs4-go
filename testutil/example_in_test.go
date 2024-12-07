package testutil_test

import (
	"errors"
	"fmt"

	"github.com/youngzhu/algs4-go/testutil"
)

func ExampleIn_ReadString() {
	in := testutil.NewInReadWords("testdata/in.txt")
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
	in := testutil.NewIn("testdata/in.txt")
	for in.HasNext() {
		fmt.Println(in.ReadLine())
	}
	// Output:
	// hello Gopher
	// wating for  you
}

func ExampleIn_ReadString_gz() {
	in := testutil.NewInReadWords("testdata/in.txt.gz")
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
	in := testutil.NewInReadWords("testdata/ints.txt")
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
	in := testutil.NewInReadWords("testdata/ints.txt")

	fmt.Println(in.ReadInt())
	fmt.Println(in.ReadInt())
	fmt.Println(in.ReadInt())

	// Output:
	// 12
	// 3
	// -1
}

func ExampleIn_ReadFloat() {
	in := testutil.NewInReadWords("testdata/floats.txt")
	for !in.IsEmpty() {
		fmt.Printf("%.6f\n", in.ReadFloat())
	}

	// Output:
	// 0.000000
	// 0.100000
	// 99.990000
	// 100.000000
	// 0.000009
}

func ExampleIn_ReadAllStrings() {
	in := testutil.NewInReadWords("testdata/in.txt")
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [hello Gopher wating for you]
}

func ExampleIn_ReadAllInts() {
	in := testutil.NewInReadWords("testdata/ints.txt")
	s := in.ReadAllInts()
	fmt.Println(s)
	// Output:
	// [12 3 -1 5 6]
}

func ExampleIn_ReadAllStrings_http() {
	const url = "https://algs4.cs.princeton.edu/24pq/tiny.txt"
	in := testutil.NewInReadWords(url)
	s := in.ReadAllStrings()
	fmt.Println(s)
	// Output:
	// [S O R T E X A M P L E]
}

func ExampleIn_ReadAll_http() {
	const url = "https://algs4.cs.princeton.edu/21elementary/words3.txt"
	in := testutil.NewIn(url)
	s := in.ReadAll()
	fmt.Println(s)

	// Output:

	// bed bug dad yes zoo
	// now for tip ilk dim
	// tag jot sob nob sky
	// hut men egg few jay
	// owl joy rap gig wee
	// was wad fee tap tar
	// dug jam all bad yet
}

func ExampleIn_ReadAllLines() {
	in := testutil.NewIn("testdata/in.txt")
	lines := in.ReadAllLines()

	for _, line := range lines {
		fmt.Println(line)
	}

	// Output:
	// hello Gopher
	// wating for  you
}

func ExampleIn_error() {
	_, err := testutil.NewInWithError("")

	fmt.Println(errors.Is(err, testutil.ErrEmpty))

	// Output:
	// true
}
