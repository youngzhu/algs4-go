package sorting_test

import (
	"testing"
	"math"

	. "github.com/youngzhu/algs4-go/sorting"
)

// test data
// array
var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

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