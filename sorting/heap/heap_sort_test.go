package heap_test

import (
	"math"
	"testing"

	"github.com/youngzhu/algs4-go/sorting"
	. "github.com/youngzhu/algs4-go/sorting/heap"
)

var intsData = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64sData = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var stringsData = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestHeapsortInts(t *testing.T) {
	data := intsData
	x := sorting.IntSortSlice(data[0:])
	Heapsort(x)
	if !sorting.IsSorted(x) {
		t.Errorf("sorting %v", intsData)
		t.Errorf("    got %v", data)
	}
}

func TestHeapsortFloat64s(t *testing.T) {
	data := float64sData
	x := sorting.Float64SortSlice(data[0:])
	Heapsort(x)
	if !sorting.IsSorted(x) {
		t.Errorf("sorting %v", float64sData)
		t.Errorf("    got %v", data)
	}
}

func TestHeapsortStrings(t *testing.T) {
	data := stringsData
	x := sorting.StringSortSlice(data[0:])
	Heapsort(x)
	if !sorting.IsSorted(x) {
		t.Errorf("sorting %v", stringsData)
		t.Errorf("    got %v", data)
	}
}
