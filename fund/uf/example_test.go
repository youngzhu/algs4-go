package uf_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/fund/uf"
	"github.com/youngzhu/algs4-go/util"
)

func ExampleUF() {

	in := util.NewInReadWords("testdata/tinyUF.txt")

	n := in.ReadInt()
	unionFind := uf.NewUF(n)

	for !in.IsEmpty() {
		p := in.ReadInt()
		q := in.ReadInt()

		if unionFind.Find(p) == unionFind.Find(q) {
			continue
		}

		unionFind.Union(p, q)
		fmt.Printf("%d %d\n", p, q)
	}

	fmt.Printf("%d components", unionFind.Count())

	// Output:
	// 4 3
	// 3 8
	// 6 5
	// 9 4
	// 2 1
	// 5 0
	// 7 2
	// 6 1
	// 2 components
}
