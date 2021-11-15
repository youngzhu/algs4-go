package pq_test

import (
	"testing"

	"github.com/youngzhu/algs4-go/sorting/pq"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var arr = []string{"S", "I", "M", "-", "P", "L", "E", "H", "-", "-", "E", "A", "P"}

func testHeap(maxPQ *pq.MaxPQ) {
	for _, item := range arr {
		if item == "-" {
			maxPQ.Delete()
		} else {
			maxPQ.Insert(pq.StringItem(item))
		}
	}
}

func BenchmarkMaxPQ_based0(b *testing.B) {
	maxPQ := pq.NewMaxPQ_based0()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(maxPQ)
	}
}

func BenchmarkMaxPQ_based1(b *testing.B) {
	maxPQ := pq.NewMaxPQ_based1()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(maxPQ)
	}
}

func BenchmarkMinPQ(b *testing.B) {
	maxPQ := pq.NewMaxPQ_based1()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testHeap(maxPQ)
	}
}
