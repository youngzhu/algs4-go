package heap_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting/heap"
)

var arr = []string{"S", "I", "M", "-", "P", "L", "E", "H", "-", "-", "E", "A", "P"}

func testHeap(h *MinHeap) {
	for _, item := range arr {
		if item == "-" {
			h.Remove()
		} else {
			if h.IsFull() {
				h.Remove()
			}
			h.Insert(StringItem(item))
		}

	}
}

func BenchmarkBased0(b *testing.B) {
	heap := NewMinHeapBased0()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(heap)
	}
}

func BenchmarkBased1(b *testing.B) {
	heap := NewMinHeapBased1()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(heap)
	}
}
