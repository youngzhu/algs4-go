package util_test

import (
	"fmt"
	"math"

	"github.com/youngzhu/algs4-go/util"
)

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