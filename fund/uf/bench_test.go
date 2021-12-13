package uf_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/fund/uf"
	"github.com/youngzhu/algs4-go/util"
)

// go test -v -run="none" -bench="." -benchtime="3s"

type pair struct {
	p, q int
}

type testCase struct {
	n int
	pairs []pair
}

var testdata *testCase

func dataInit() {
	in := util.NewInReadWords("testdata/mediumUF.txt")

	n := in.ReadInt()
	pairs := make([]pair, 1)

	for !in.IsEmpty() {
		p := in.ReadInt()
		q := in.ReadInt()
		pairs = append(pairs, pair{p, q})
	}

	testdata = &testCase{n, pairs}
}

func benchUF(unionFind UnionFind, pairs []pair) {
	for _, pq := range pairs {
		if unionFind.Find(pq.p) == unionFind.Find(pq.q) {
			continue
		}
	
		unionFind.Union(pq.p, pq.q)
	}
}

func BenchmarkUF(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	unionFind := NewUF(testdata.n)

	for i := 0; i < b.N; i++ {
		benchUF(unionFind, testdata.pairs)
	}
}

func BenchmarkQuickFindUF(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	unionFind := NewQuickFindUF(testdata.n)

	for i := 0; i < b.N; i++ {
		benchUF(unionFind, testdata.pairs)
	}
}

func BenchmarkQuickUnionUF(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	unionFind := NewQuickUnionUF(testdata.n)

	for i := 0; i < b.N; i++ {
		benchUF(unionFind, testdata.pairs)
	}
}

func BenchmarkWeightedQuickUnionUF(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	unionFind := NewWeightedQuickUnionUF(testdata.n)

	for i := 0; i < b.N; i++ {
		benchUF(unionFind, testdata.pairs)
	}
}

/* Run Parallel */

func BenchmarkUFParallel(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	b.RunParallel(func(pb *testing.PB) {
		unionFind := NewUF(testdata.n)
		for pb.Next() {
			benchUF(unionFind, testdata.pairs)
		}
	})
}

func BenchmarkQuickFindUFParallel(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	b.RunParallel(func(pb *testing.PB) {
		unionFind := NewQuickFindUF(testdata.n)
		for pb.Next() {
			benchUF(unionFind, testdata.pairs)
		}
	})
}

func BenchmarkQuickUnionUFParallel(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	b.RunParallel(func(pb *testing.PB) {
		unionFind := NewQuickUnionUF(testdata.n)
		for pb.Next() {
			benchUF(unionFind, testdata.pairs)
		}
	})
}

func BenchmarkWeightedQuickUnionUFParallel(b *testing.B) {
	b.ReportAllocs()
	if testdata == nil {
		b.StopTimer()
		dataInit()
		b.StartTimer()
	}

	b.RunParallel(func(pb *testing.PB) {
		unionFind := NewWeightedQuickUnionUF(testdata.n)
		for pb.Next() {
			benchUF(unionFind, testdata.pairs)
		}
	})
}

// BenchmarkUF
// BenchmarkUF-4                                      56833             20741 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickFindUF
// BenchmarkQuickFindUF-4                            111165             10775 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickUnionUF
// BenchmarkQuickUnionUF-4                             4996            216117 ns/op               1 B/op          0 allocs/op
// BenchmarkWeightedQuickUnionUF
// BenchmarkWeightedQuickUnionUF-4                    41350             26305 ns/op               0 B/op          0 allocs/op
// BenchmarkUFParallel
// BenchmarkUFParallel-4                             116836             10029 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickFindUFParallel
// BenchmarkQuickFindUFParallel-4                    216003              5545 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickUnionUFParallel
// BenchmarkQuickUnionUFParallel-4                    18840             60550 ns/op               1 B/op          0 allocs/op
// BenchmarkWeightedQuickUnionUFParallel
// BenchmarkWeightedQuickUnionUFParallel-4           115634             10181 ns/op               0 B/op          0 allocs/op