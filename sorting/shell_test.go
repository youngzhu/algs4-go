package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

func TestShellsortInts(t *testing.T) {
	data := ints
	x := IntCompSlice(data[0:])
	Shellsort(x)
	if ! IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}

func TestShellsortFloat64s(t *testing.T) {
	data := float64s
	x := Float64CompSlice(data[0:])
	Shellsort(x)
	if ! IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}

func TestShellsortStrings(t *testing.T) {
	data := strings
	x := StringCompSlice(data[0:])
	Shellsort(x)
	if ! IsSorted(x) {
		t.Errorf("sorting %v", strings)
		t.Errorf("    got %v", data)
	}
}