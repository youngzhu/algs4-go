package selection_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting/selection"
	"github.com/youngzhu/algs4-go/sorting"
)

// array
var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInts(t *testing.T) {
	data := ints
	x := sorting.IntCompSlice(data[0:])
	Sort(x)
	if ! sorting.IsSorted(x) {
		t.Errorf("sorting %v", ints)
		t.Errorf("    got %v", data)
	}
}
