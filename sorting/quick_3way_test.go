package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

func TestQuicksort3wayInts(t *testing.T) {
	data := ints
	x := IntSortSlice(data[0:])
	Quicksort3way(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}

func TestQuicksort3wayFloat64s(t *testing.T) {
	data := float64s
	x := Float64SortSlice(data[0:])
	Quicksort3way(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}

func TestQuicksort3wayStrings(t *testing.T) {
	data := strings
	x := StringSortSlice(data[0:])
	Quicksort3way(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", strings)
		t.Errorf("    got %v", data)
	}
}
