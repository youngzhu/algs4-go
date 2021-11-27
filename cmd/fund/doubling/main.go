package main

import (
	"flag"

	"github.com/youngzhu/algs4-go/fund/xsum"
	"github.com/youngzhu/algs4-go/util"
)

var (
	rand  util.Random
	ratio = flag.Bool("r", false, "doubling ratio test")
)

func init() {
	rand = *util.NewRandom()
}

func main() {
	flag.Parse()

	if *ratio {
		doublingRatio()
	} else {
		doublingTest()
	}
}

const maxInteger = 1000000

// Returns the amount of time to call ThreeSumCount() with n random
// 6-digit integers
func timeTrial(n int) float64 {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.UniformIntRange(-maxInteger, maxInteger)
	}
	timer := util.NewStopwatch()
	xsum.ThreeSumCount(a)
	return timer.ElapsedTime()
}
