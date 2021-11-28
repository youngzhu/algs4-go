package xsum

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/sorting"
)

var (
	sorter sorting.Sorter = sorting.Quick{}

	// fast way
	binarySearch = fund.NewBinarySearch()
)

// returns true if the sorted array a[]
// contains any duplicated integers
func containsDuplicates(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			return true
		}
	}

	return false
}

//
