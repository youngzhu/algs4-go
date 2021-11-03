package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/youngzhu/algs4-go/sorting"
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
type Buildin struct{}

func newBuildin() sorting.Sorter {
	return Buildin{}
}

func (s Buildin) SortInts(x []int) {
	sort.Ints(x)
}
func (s Buildin) SortFloat64s(x []float64) {
	sort.Float64s(x)
}
func (s Buildin) SortStrings(x []string) {
	sort.Strings(x)
}

var algs = map[string]sorting.Sorter{
	"selection": sorting.NewSelection(),
	"insertion": sorting.NewInsertion(),
	"buildin":   newBuildin(),
	"shell":     sorting.NewShell(),
	"merge":     sorting.NewMerge(),
	"mergex1":   sorting.NewMergeX1(),
	"mergex2":   sorting.NewMergeX2(),
	"mergex3":   sorting.NewMergeX3(),
	"mergex":    sorting.NewMergeX(),
	"mergebu":   sorting.NewMergeBU(),
	"quick":     sorting.NewQuick(),
	"quick3way": sorting.NewQuick3way(),
}

func init() {
	flag.StringVar(&alg1, "a1", "", "algorithm one")
	flag.StringVar(&alg2, "a2", "", "algorithm two")
	flag.BoolVar(&isSorted, "s", false, "is array sorted")
	flag.IntVar(&n, "n", 0, "array size")
	flag.IntVar(&trials, "t", 0, "run times")

	rand.Seed(time.Now().Unix())
}

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

	for t := 0; t < trials; t++ {
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = rand.Float64()
		}
		// fmt.Println("round:", t, "len:", len(a))
		total += timeElapsed(alg, a)
	}
	return total
}

func timeSortedInput(alg string, n, trials int) int64 {
	var total int64

	for t := 0; t < trials; t++ {
		a := make([]float64, n)
		for i := 0; i < n; i++ {
			a[i] = float64(i)
		}

		// fmt.Println("round:", t, "len:", len(a))
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

// TEST RESULT
// insertion vs selection
// cmd: go run sort_compare.go -a1 insertion -a2 selection -n 100 -t 100 -s
// result: insertion is 27.7 times faster than selection
// cmd: go run sort_compare.go -a1 insertion -a2 Selection -n 100 -t 100
// result: insertion is 42.4 times faster than Selection

// insertion vs buildin
// cmd: go run sort_compare.go -a1 insertion -a2 buildin -n 100 -t 100 -s
// result: buildin is 9.8 times faster than insertion
// cmd: go run sort_compare.go -a1 insertion -a2 buildin -n 100 -t 100
// result: buildin is 5.9 times faster than insertion

// Selection vs Buildin
// cmd: go run sort_compare.go -a1 Selection -a2 Buildin -n 100 -t 100 -s
// result: Buildin is 452.9 times faster than Selection
// cmd: go run sort_compare.go -a1 Selection -a2 Buildin -n 100 -t 100
// result: Buildin is 243.7 times faster than Selection

// Shell vs Selection
// cmd: go run sort_compare.go -a1 Shell -a2 Selection -n 100 -t 100 -s
// got: Shell is 246.4 times faster than Selection
// cmd: go run sort_compare.go -a1 Shell -a2 Selection -n 100 -t 100
// got: Shell is 123.3 times faster than Selection

// Shell vs Insertion
// cmd: go run sort_compare.go -a1 Shell -a2 Insertion -n 100 -t 100 -s
// got: Shell is 5.9 times faster than Insertion
// cmd: go run sort_compare.go -a1 Shell -a2 Insertion -n 100 -t 100
// got: Shell is 3.4 times faster than Insertion

// Shell vs Buildin
// cmd: go run sort_compare.go -a1 Shell -a2 Buildin -n 100 -t 100 -s
// got: Buildin is 1.2 times faster than Shell
// cmd: go run sort_compare.go -a1 Shell -a2 Buildin -n 100 -t 100
// got: Buildin is 1.8 times faster than Shell

// Merge vs Insetion
// cmd: go run sort_compare.go -a1 Merge -a2 Insertion -n 100 -t 100 -s
// got: Insertion is 4.9 times faster than Merge
// cmd: go run sort_compare.go -a1 Insertion -a2 Merge -n 100 -t 100 -s
// got: Insertion is 4.3 times faster than Merge
// cmd: go run sort_compare.go -a1 Insertion -a2 Merge -n 100 -t 100
// got: Insertion is 4.2 times faster than Merge

// Merge vs Selection
// cmd: go run sort_compare.go -a1 Merge -a2 Selection -n 100 -t 100 -s
// got: Merge is 8.4 times faster than Selection
// cmd: go run sort_compare.go -a1 Merge -a2 Selection -n 100 -t 100
// got: Merge is 8.8 times faster than Selection

// Merge vs Shell
// cmd: go run sort_compare.go -a1 Merge -a2 Shell -n 100 -t 100 -s
// got: Shell is 42.2 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 Shell -n 100 -t 100
// got: Shell is 17.3 times faster than Merge

// Merge vs Buildin
// cmd: go run sort_compare.go -a1 Merge -a2 Buildin -n 100 -t 100 -s
// got: Buildin is 44.3 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 Buildin -n 100 -t 100
// got: Buildin is 20.1 times faster than Merge

// Merge vs MergeX1
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX1 -n 100 -t 100 -s
// got: MergeX1 is 2.31 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX1 -n 100 -t 100
// got: MergeX1 is 2.82 times faster than Merge

// Merge vs MergeX2
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX2 -n 100 -t 100 -s
// got: MergeX2 is 114.35 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX2 -n 100 -t 100
// got: MergeX2 is 54.61 times faster than Merge

// Merge vs MergeX3
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX3 -n 100 -t 100 -s
// got: MergeX3 is 48.5 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX3 -n 100 -t 100
// got: MergeX3 is 52.9 times faster than Merge

// Merge vs MergeX
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX -n 100 -t 100 -s
// got: MergeX is 2.6 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 MergeX -n 100 -t 100
// got: MergeX is 3.5 times faster than Merge

// Merge vs MergeBU
// cmd: go run sort_compare.go -a1 Merge -a2 MergeBU -n 100 -t 100 -s
// got: MergeBU is 54.6 times faster than Merge
// cmd: go run sort_compare.go -a1 Merge -a2 MergeBU -n 100 -t 100
// got: MergeBU is 55.9 times faster than Merge

// NOTE: following test result about Quicksort is no shuffle
// Quick vs Selection
// cmd: go run sort_compare.go -a1 Quick -a2 Selection -n 1000 -t 100 -s
// got: Quick is 1.3 times faster than Selection
// cmd: go run sort_compare.go -a1 Quick -a2 Selection -n 1000 -t 100
// got: Quick is 20.6 times faster than Selection

// Quick vs Insertion
// cmd: go run sort_compare.go -a1 Quick -a2 Insertion -n 1000 -t 100 -s
// got: Insertion is 337.0 times faster than Quick
// cmd: go run sort_compare.go -a1 Quick -a2 Insertion -n 1000 -t 100
// got: Quick is 16.6 times faster than Insertion

// Quick vs Shell
// cmd: go run sort_compare.go -a1 Quick -a2 Shell -n 1000 -t 100 -s
// got: Shell is 67.4 times faster than Quick
// cmd: go run sort_compare.go -a1 Quick -a2 Shell -n 1000 -t 100
// got: Quick is 1.4 times faster than Shell

// Quick vs Merge
// cmd: go run sort_compare.go -a1 Quick -a2 Merge -n 1000 -t 100 -s
// got: Merge is 11.5 times faster than Quick
// cmd: go run sort_compare.go -a1 Quick -a2 Merge -n 1000 -t 100
// got: Quick is 1.8 times faster than Merge

// Quick vs Buildin
// cmd: go run sort_compare.go -a1 Quick -a2 Buildin -n 1000 -t 100 -s
// got: Buildin is 50.0 times faster than Quick
// cmd: go run sort_compare.go -a1 Quick -a2 Buildin -n 1000 -t 100
// got: Quick is 1.2 times faster than Buildin

// Quick vs Quick3way
// cmd: go run sort_compare.go -a1 Quick -a2 Quick3way -n 1000 -t 100 -s
// got: Quick is 1.4 times faster than Quick3way
// cmd: go run sort_compare.go -a1 Quick -a2 Quick3way -n 1000 -t 100
// got: Quick is 2.1 times faster than Quick3way
