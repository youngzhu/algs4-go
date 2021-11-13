package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var a = [...]int{9, 10, 0, 7, 8, 4, 3, 6, 2, 1, 5, 5, -99}

func BenchmarkSelection(b *testing.B) {
	soter := NewSelection()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkInsertion(b *testing.B) {
	soter := NewInsertion()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkShell(b *testing.B) {
	soter := NewShell()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkMerge(b *testing.B) {
	soter := NewMerge()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkMergeX(b *testing.B) {
	soter := NewMergeX()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkMergeBU(b *testing.B) {
	soter := NewMergeBU()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkQuick(b *testing.B) {
	soter := NewQuick()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkQuick3way(b *testing.B) {
	soter := NewQuick3way()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkHeap(b *testing.B) {
	soter := NewHeap()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}
