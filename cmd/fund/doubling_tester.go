package main

import (
	"fmt"

	"github.com/youngzhu/algs4-go/util"
)

var rand util.Random

func init() {
	rand = *util.NewRandom()
}

// Generates a sequence of random input arrays, doubling the array size
// at each step, and prints the running times of ThreeSumCount for each input size.

// cmd: go run *.go
func main() {
	for n := 250; true; n += n {
		time := timeTrial(n)
		fmt.Printf("%7d %7.1f\n", n, time)
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
	ThreeSumCount(a)
	return timer.ElapsedTime()
}
