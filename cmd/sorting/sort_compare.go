package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"strings"
	"os"

	. "github.com/youngzhu/algs4-go/sorting"
)
// Uses the sort() methods int the types named as command-line arguments 
// to perform the given number of experiments (sorting arrays of the given size)
// and prints the ration of the observed running times of the algorithms.

// Sort n random real numbers, trials times using the two algorithms specified
// on the command line.

var alg1, alg2 string
var isSorted bool
var n, trials int

var algs = map[string] Sorter {
	"selection": Selection{},
	"insertion": Insertion{},
}

func init() {
	flag.StringVar(&alg1, "a1", "", "algorithm one")
	flag.StringVar(&alg2, "a2", "", "algorithm two")
	flag.BoolVar(&isSorted, "s", false, "is array sorted")
	flag.IntVar(&n, "n", 0, "array size")
	flag.IntVar(&trials, "t", 0, "run times")

	rand.Seed(time.Now().Unix())
}

// PUT s THE LAST, OTHERWISE n,t GOT 0
// DO NOT
// go run sort_compare.go -a1 insertion -a2 selection -s true -n 1000 -t 100

// go run sort_compare.go -a1 insertion -a2 selection -n 100 -t 100 -s true
// result: insertion is 27.7 times faster than selection
// go run sort_compare.go -a1 insertion -a2 Selection -n 100 -t 100
// result: insertion is 42.4 times faster than Selection
func main() {
	flag.Parse() // parse the command line into the defined flags

	fmt.Println("alg1:", alg1, ",alg2:", alg2, ",n:", n, ",trials:", trials, ", isSorted:", isSorted)

	var time1, time2 int64

	if isSorted {
		time1 = timeSortedInput(alg1, n, trials)
		time2 = timeSortedInput(alg2, n, trials)
	} else {
		time1 = timeRandomInput(alg1, n, trials)
		time2 = timeRandomInput(alg2, n, trials)
	}

	fasterAlg, fasterTime := alg1, float64(time1)
	slowerAlg, slowerTime := alg2, float64(time2)

	if time1 > time2 {
		fasterAlg, fasterTime = alg2, float64(time2)
		slowerAlg, slowerTime = alg1, float64(time1)
	}

	if fasterTime == 0 {
		fmt.Println("The given number of experiments is too small.")
		os.Exit(1)
	}

	fmt.Printf("For %d random floats\n   %s is", n, fasterAlg)
	fmt.Printf(" %.1f times faster than %s\n", slowerTime/fasterTime, slowerAlg)
}

func timeRandomInput(alg string, n, trials int) int64 {
	var total int64
	a := make([]float64, n)
	for t := 0; t < trials; t++ {
		for i := 0; i < n; i++ {
			a = append(a, rand.Float64())
		}
		total += timeElapsed(alg, a)
	}
	return total 
}

func timeSortedInput(alg string, n, trials int) int64 {
	var total int64
	a := make([]float64, n)
	for t := 0; t < trials; t++ {
		for i := 0; i < n; i++ {
			a = append(a, float64(i))
		}
		total += timeElapsed(alg, a)
	}
	return total 
}

func timeElapsed(alg string, a []float64) int64 {
	var elapsed int64

	lowcase := strings.ToLower(alg)

	sorter, ok := algs[lowcase]
	if ok {
		start := time.Now()
		sorter.SortFloat64s(a)
		end := time.Now()
		elapsed = end.Sub(start).Milliseconds()
	} else {
		panic("Invalid algorithm: " + lowcase)
	}

	// fmt.Println("elapsed: ", elapsed)

	return elapsed
}