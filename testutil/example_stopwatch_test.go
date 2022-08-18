package testutil_test

import (
	"fmt"
	"math"

	"github.com/youngzhu/algs4-go/testutil"
)

func ExampleStopwatch() {
	var n float64 = 1000
	timer := testutil.NewStopwatch()
	sum := 0.0
	for n > 0 {
		sum += math.Sqrt(n)
		n--
	}
	time := timer.ElapsedTime()
	fmt.Printf("%e (%.2f seconds)", sum, time)
}
