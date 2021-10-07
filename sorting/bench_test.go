package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var a = [...]int{9, 10, 0, 7, 8, 4, 3, 6, 2, 1, 5, 5, -99}

func BenchmarkSelection(b *testing.B) {
	soter := Selection{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkInsertion(b *testing.B) {
	soter := Insertion{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkShell(b *testing.B) {
	soter := Shell{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}

func BenchmarkMerge(b *testing.B) {
	soter := Merge{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(a[0:])
	}
}