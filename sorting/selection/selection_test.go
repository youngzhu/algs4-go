package selection_test

import (
	"testing"
	"math"

	. "github.com/youngzhu/algs4-go/sorting/selection"
	"github.com/youngzhu/algs4-go/sorting"
)

// array
var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}

func TestInts(t *testing.T) {
	data := ints
	x := sorting.IntCompSlice(data[0:])
	Sort(x)
	if ! sorting.IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}

func TestFloat64s(t *testing.T) {
	data := float64s
	x := sorting.Float64CompSlice(data[0:])
	Sort(x)
	if ! sorting.IsSorted(x) {
		t.Errorf("sorting %v", float64s)
		t.Errorf("    got %v", data)
	}
}