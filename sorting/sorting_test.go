package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

func TestIsSortedInt(t *testing.T) {
	sorted := IntCompSlice([]int{-99, 0, 99, 100})
	if !IsSorted(sorted) {
		t.Errorf("excepted true, got false")
	}

	unsorted := IntCompSlice([]int{1000, 88888, -1, 9, 66})
	if IsSorted(unsorted) {
		t.Errorf("excepted false, got true")
	}
}