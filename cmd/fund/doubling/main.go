package main

import (
	"flag"

	"github.com/youngzhu/algs4-go/fund/xsum"
	"github.com/youngzhu/algs4-go/testutil"
)

var (
	rand  testutil.Random
	ratio = flag.Bool("r", false, "doubling ratio test")
)

func init() {
	rand = *testutil.NewRandom()
}

func main() {
	flag.Parse()

	if *ratio {
		doublingRatio()
	} else {
		doublingTest()
	}
}

const (
	min = -1000000
	max = 1000000
)

// Returns the amount of time to call ThreeSumCount() with n random
// 6-digit integers
func timeTrial(n int) float64 {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.UniformIntRange(min, max)
	}
	timer := testutil.NewStopwatch()

	xsum.ThreeSumCount(a)

	// xsum.TwoSumCount(a)
	// xsum.TwoSumCountFast(a)

	return timer.ElapsedTime()
}
