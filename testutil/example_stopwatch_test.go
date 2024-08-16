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
	fmt.Printf("elapsed %.2f seconds", elapsedTime)

	// Output:

}
