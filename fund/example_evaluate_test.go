package fund_test

import (
	"fmt"
	"github.com/youngzhu/algs4-go/fund"
)

func ExampleEvaluate() {
	exp := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"

	fmt.Printf("%.1f", fund.Evaluate(exp))

	// Output:
	// 101.0
}