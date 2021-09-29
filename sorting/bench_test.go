package sorting_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/sorting"
)

// go test -v -run="none" -bench="."

var a = [...]int {9, 10, 0, 7, 8, 4, 3, 6, 2, 1, 5, 5}

func BenchmarkSelection(b *testing.B) {
	soter := Selection{}
	s := a[0:]

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(s)
	}
}

func BenchmarkInsertion(b *testing.B) {
	soter := Insertion{}
	s := a[0:]

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(s)
	}
}

func BenchmarkShell(b *testing.B) {
	soter := Shell{}
	s := a[0:]

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		soter.SortInts(s)
	}
}