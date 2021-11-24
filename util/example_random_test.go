package util_test

import (
	"fmt"

	"github.com/youngzhu/algs4-go/util"
)

func ExampleRandom_UniformIntN() {
	r := util.NewRandom()

	// Seeding with the same value results in the same random
	// sequence each run.
	r.Seed(100)

	fmt.Println(r.UniformIntN(100))
	fmt.Println(r.UniformIntN(100))
	fmt.Println(r.UniformIntN(100))

	// Output:
	// 83
	// 68
	// 80
}

func ExampleRandom_UniformIntRange() {
	r := util.NewRandom()

	// Seeding with the same value results in the same random
	// sequence each run.
	r.Seed(9999)

	a, b := -10000, 10000

	fmt.Println(r.UniformIntRange(a, b))
	fmt.Println(r.UniformIntRange(a, b))
	fmt.Println(r.UniformIntRange(a, b))

	// Output:
	// -1219
	// 8260
	// -8752
}
