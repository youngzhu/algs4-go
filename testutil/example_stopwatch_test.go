package testutil_test

import (
	"fmt"
	"time"

	"github.com/youngzhu/algs4-go/testutil"
)

func ExampleStopwatch() {
	timer := testutil.NewStopwatch()
	time.Sleep(5 * time.Second)
	elapsedTime := timer.ElapsedTime()
	fmt.Printf("elapsed %.0f+ seconds", elapsedTime)

	// Output:
	// elapsed 5+ seconds

}
