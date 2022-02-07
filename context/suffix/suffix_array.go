package suffix

import (
	"github.com/youngzhu/algs4-go/strings/sort"
	"regexp"
)

// Suffix sorting: given a string, sort the suffixes of that string in ascending order.
// Resulting sorted list is called a suffix array.

type SuffixArray struct {
	suffixes []string
	n        int
}

func NewSuffixArray(text string) SuffixArray {
	// not work
	//txt := strings.ReplaceAll(text, "\\s+", " ")

	space := regexp.MustCompile("\\s+")
	txt := string(space.ReplaceAll([]byte(text), []byte(" ")))
	//log.Println(txt)

	n := len(txt)
	suffixes := make([]string, n)
	for i := 0; i < n; i++ {
		suffixes[i] = txt[i:]
	}
	sort.Quicksort(suffixes)

	return SuffixArray{suffixes: suffixes, n: n}
}

// Length
// Returns the length of the input string
func (sa SuffixArray) Length() int {
	return sa.n
}

// Index
// Returns the index into the original string of the ith smallest suffix.
func (sa SuffixArray) Index(i int) int {
	if i < 0 || i >= sa.n {
		panic("Illegal index")
	}
	return sa.n - len(sa.suffixes[i])
}

// LCP returns the length of the longest common prefix of the ith smallest suffix
// and the i-1 th smallest suffix.
func (sa SuffixArray) LCP(i int) int {
	if i < 1 || i >= sa.n {
		panic("Illegal argument")
	}
	return lcp(sa.suffixes[i], sa.suffixes[i-1])
}

// longest common prefix of s and t
func lcp(s, t string) int {
	n := min(len(s), len(t))
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			return i
		}
	}
	return n
}

// Select
// Returns the ith smallest suffix as a string
func (sa SuffixArray) Select(i int) string {
	if i < 0 || i >= sa.n {
		panic("Illegal argument")
	}
	return sa.suffixes[i]
}

// Rank
// Returns the number of suffixes strictly less than the query string.
// Note: rank(select(i)) == i
func (sa SuffixArray) Rank(query string) int {
	lo, hi := 0, sa.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midSuffix := sa.suffixes[mid]
		switch {
		case query < midSuffix:
			hi = mid - 1
		case query > midSuffix:
			lo = mid + 1
		default:
			return mid
		}
	}
	return lo
}
