package main

import (
	"fmt"
)

func doublingRatio() {
	prev := timeTrial(125)
	for n := 250; true; n += n {
		time := timeTrial(n)
		fmt.Printf("%7d %7.1f %5.1f\n", n, time, time/prev)
		prev = time
	}
}
