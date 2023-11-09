package pq

import (
	"testing"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var arr = []string{"S", "I", "M", "-", "P", "L", "E", "H", "-", "-", "E", "A", "P"}

func testHeap(maxPQ *MaxPQ) {
	for _, item := range arr {
		if item == "-" {
			maxPQ.Delete()
		} else {
			maxPQ.Insert(StringItem(item))
		}
	}
}

func BenchmarkMaxPQ_based0(b *testing.B) {
	maxPQ := newMaxPQBased0()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(maxPQ)
	}
}

func BenchmarkMaxPQ_based1(b *testing.B) {
	maxPQ := newMaxPQBased1()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(maxPQ)
	}
}

func BenchmarkMinPQ(b *testing.B) {
	maxPQ := newMaxPQBased1()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(maxPQ)
	}
}

func newMaxPQBased0() *MaxPQ {
	items := make([]Item, 1)
	heap := NewBinaryHeapBased0()
	return &MaxPQ{items, 0, heap}
}
func newMaxPQBased1() *MaxPQ {
	items := make([]Item, 1)
	heap := NewBinaryHeapBased1()
	return &MaxPQ{items, 0, heap}
}
