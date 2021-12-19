package strings_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/strings"
)

var testCases = []struct {
	pattern, text string
	index         int
}{
	{"abracadabra", "abacadabrabracabracadabrabrabracad", 14},
	{"rab", "abacadabrabracabracadabrabrabracad", 8},
	{"rabrabracad", "abacadabrabracabracadabrabrabracad", 23},
	{"bcara", "abacadabrabracabracadabrabrabracad", -1},
	{"abacad", "abacadabrabracabracadabrabrabracad", 0},
}

func TestBruteForceSearch1(t *testing.T) {
	for _, tc := range testCases {
		got := BruteForceSearch1(tc.pattern, tc.text)
		want := tc.index

		if got != want {
			t.Errorf("BruteSearch1(%q, %q) got %d, want %d", tc.pattern, tc.text, got, want)
		}
	}
}

func TestBruteForceSearch2(t *testing.T) {
	for _, tc := range testCases {
		got := BruteForceSearch2(tc.pattern, tc.text)
		want := tc.index

		if got != want {
			t.Errorf("BruteSearch1(%q, %q) got %d, want %d", tc.pattern, tc.text, got, want)
		}
	}
}

func TestRabinKarp(t *testing.T) {
	for _, tc := range testCases {
		rk := NewRabinKarp(tc.pattern)
		got := rk.Search(tc.text)
		want := tc.index

		if got != want {
			t.Errorf("BruteSearch1(%q, %q) got %d, want %d", tc.pattern, tc.text, got, want)
		}
	}
}

func TestKMP(t *testing.T) {
	for _, tc := range testCases {
		rk := NewKMP(tc.pattern)
		got := rk.Search(tc.text)
		want := tc.index

		if got != want {
			t.Errorf("BruteSearch1(%q, %q) got %d, want %d", tc.pattern, tc.text, got, want)
		}
	}
}
