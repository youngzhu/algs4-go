package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

var intsX2 = [...]int{74, 59, 238, -784, 98, 959, 999, 1000, 6666, 2333, 9999}

func TestMergesortInts(t *testing.T) {
	data := intsX2
	x := IntCompSlice(data[0:])
	Mergesort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}

func TestMergesortFloat64s(t *testing.T) {
	data := float64s
	x := Float64CompSlice(data[0:])
	Mergesort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}

func TestMergesortStrings(t *testing.T) {
	data := strings
	x := StringCompSlice(data[0:])
	Mergesort(x)
	if !IsSorted(x) {
		t.Errorf("sorting %v", strings)
		t.Errorf("    got %v", data)
	}
}
