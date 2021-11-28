package xsum_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/fund/xsum"
	"github.com/youngzhu/algs4-go/util"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var a = util.NewIn("testdata/1Kints.txt").ReadAllInts()

func BenchmarkTwoSumCount(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TwoSumCount(a[0:])
	}
}

func BenchmarkTwoSumCountFast(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TwoSumCountFast(a[0:])
	}
}
