package strings_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/strings"
)

// go test -v -run="none" -bench="." -benchtime="3s"

var benchCases = []struct {
	pattern, text string
}{
	{"abracadabra", "abacadabrabracabracadabrabrabracad"},
	{"rab", "abacadabrabracabracadabrabrabracad"},
	{"rabrabracad", "abacadabrabracabracadabrabrabracad"},
	{"bcara", "abacadabrabracabracadabrabrabracad"},
	{"abacad", "abacadabrabracabracadabrabrabracad"},
}

func BenchmarkBruteForceSearch1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range benchCases {
			BruteForceSearch1(tc.pattern, tc.text)
		}
	}
}

func BenchmarkBruteForceSearch2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range benchCases {
			BruteForceSearch2(tc.pattern, tc.text)
		}
	}
}

func BenchmarkRabinKarpSearch(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range benchCases {
			RabinKarpSearch(tc.pattern, tc.text)
		}
	}
}

func BenchmarkKMPSearch(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range benchCases {
			KMPSearch(tc.pattern, tc.text)
		}
	}
}

func BenchmarkBoyerMooreSearch(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range benchCases {
			BoyerMooreSearch(tc.pattern, tc.text)
		}
	}
}
