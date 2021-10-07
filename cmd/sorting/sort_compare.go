package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"strings"
	"os"
	"sort"

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

// Golang buildin sort
type Buildin struct {}
func (s Buildin) SortInts(x []int) {
	sort.Ints(x)
}
func (s Buildin) SortFloat64s(x []float64) {
	sort.Float64s(x)
}
func (s Buildin) SortStrings(x []string) {
	sort.Strings(x)
}

var algs = map[string] Sorter {
	"selection": Selection{},
	"insertion": Insertion{},
	"buildin": Buildin{},
	"shell": Shell{},
	"merge": Merge{},
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

// TEST RESULT
// insertion vs selection
// go run sort_compare.go -a1 insertion -a2 selection -n 100 -t 100 -s true
// result: insertion is 27.7 times faster than selection
// go run sort_compare.go -a1 insertion -a2 Selection -n 100 -t 100
// result: insertion is 42.4 times faster than Selection

// insertion vs buildin
// go run sort_compare.go -a1 insertion -a2 buildin -n 100 -t 100 -s true
// result: buildin is 9.8 times faster than insertion
// go run sort_compare.go -a1 insertion -a2 buildin -n 100 -t 100
// result: buildin is 5.9 times faster than insertion

// Selection vs Buildin
// go run sort_compare.go -a1 Selection -a2 Buildin -n 100 -t 100 -s true
// result: Buildin is 452.9 times faster than Selection
// go run sort_compare.go -a1 Selection -a2 Buildin -n 100 -t 100
// result: Buildin is 243.7 times faster than Selection

// Shell vs Selection
// go run sort_compare.go -a1 Shell -a2 Selection -n 100 -t 100 -s true
// got: Shell is 246.4 times faster than Selection
// go run sort_compare.go -a1 Shell -a2 Selection -n 100 -t 100
// got: Shell is 123.3 times faster than Selection

// Shell vs Insertion
// go run sort_compare.go -a1 Shell -a2 Insertion -n 100 -t 100 -s true
// got: Shell is 5.9 times faster than Insertion
// go run sort_compare.go -a1 Shell -a2 Insertion -n 100 -t 100
// got: Shell is 3.4 times faster than Insertion

// Shell vs Buildin
// go run sort_compare.go -a1 Shell -a2 Buildin -n 100 -t 100 -s true
// got: Buildin is 1.2 times faster than Shell
// go run sort_compare.go -a1 Shell -a2 Buildin -n 100 -t 100 
// got: Buildin is 1.8 times faster than Shell

// Merge vs Insetion
// go run sort_compare.go -a1 Merge -a2 Insertion -n 100 -t 100 -s
// got: Insertion is 4.9 times faster than Merge
// go run sort_compare.go -a1 Insertion -a2 Merge -n 100 -t 100 -s
// got: Insertion is 4.3 times faster than Merge
// go run sort_compare.go -a1 Insertion -a2 Merge -n 100 -t 100
// got: Insertion is 4.2 times faster than Merge

// Merge vs Selection
// go run sort_compare.go -a1 Merge -a2 Selection -n 100 -t 100 -s
// got: Merge is 8.4 times faster than Selection
// go run sort_compare.go -a1 Merge -a2 Selection -n 100 -t 100 
// got: Merge is 8.8 times faster than Selection

// Merge vs Shell
// go run sort_compare.go -a1 Merge -a2 Shell -n 100 -t 100 -s
// got: Shell is 42.2 times faster than Merge
// go run sort_compare.go -a1 Merge -a2 Shell -n 100 -t 100
// got: Shell is 17.3 times faster than Merge

// Merge vs Buildin
// go run sort_compare.go -a1 Merge -a2 Buildin -n 100 -t 100 -s
// got: Buildin is 44.3 times faster than Merge
// go run sort_compare.go -a1 Merge -a2 Buildin -n 100 -t 100
// got: Buildin is 20.1 times faster than Merge
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