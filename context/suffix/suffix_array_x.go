package suffix

import (
	"fmt"
)

const (
	cutoff = 5 // cutoff to insertion sort (any value between 0 an 12)
	// a sentinel and assumes that the character does not appear in the text
	sentinel = '\000'
)

// SuffixArrayX a data type that computes the suffix array of a string
// using 3-way radix quicksort.
type SuffixArrayX struct {
	text  []byte
	index []int // index[i]=j means text[j:] is ith largest suffix
	n     int   // number of characters in text
}

func NewSuffixArrayX(txt string) *SuffixArrayX {
	n := len(txt)
	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	txt = txt + string(sentinel)
	sa := &SuffixArrayX{text: []byte(txt), index: index, n: n}

	sa.sort(0, n-1, 0)

	return sa
}

// 3-way string quicksort lo...hi starting at dth character
func (sa *SuffixArrayX) sort(lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+cutoff {
		sa.insertion(lo, hi, d)
		return
	}

	lt, gt := lo, hi
	v := sa.text[sa.index[lo]+d]
	i := lo + 1
	for i <= gt {
		t := sa.text[sa.index[i]+d]
		if t < v {
			sa.exch(lt, i)
			lt++
			i++
		} else if t > v {
			sa.exch(i, gt)
			gt--
		} else {
			i++
		}
	}

	// a[lo...lt-1] < v = a[lt...gt] < a[gt+1...hi]
	sa.sort(lo, lt-1, d)
	if v != sentinel {
		sa.sort(lt, gt, d+1)
	}
	sa.sort(gt+1, hi, d)
}

// sort from a[lo] to a[hi], starting at the dth character
func (sa *SuffixArrayX) insertion(lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo; j-- {
			m, n := sa.index[j], sa.index[j-1]
			if !sa.less(m, n, d) {
				break
			}
			sa.exch(j, j-1)
		}
	}
}

// is text[i+d..n] < text[j+d..n]
func (sa *SuffixArrayX) less(i, j, d int) bool {
	if i == j {
		return false
	}
	i, j = i+d, j+d
	for i < sa.n && j < sa.n {
		if sa.text[i] < sa.text[j] {
			return true
		} else if sa.text[i] > sa.text[j] {
			return false
		} else {
			i++
			j++
		}
	}
	return i > j
}

// exchange index[i] and index[j]
func (sa *SuffixArrayX) exch(i, j int) {
	sa.index[i], sa.index[j] = sa.index[j], sa.index[i]
}

// Length returns the length of the input string
func (sa *SuffixArrayX) Length() int {
	return sa.n
}

// Index
// Returns the index into the original string of the ith smallest suffix.
func (sa *SuffixArrayX) Index(i int) int {
	if i < 0 || i >= sa.n {
		panic(`Illegal index: ` + fmt.Sprint(i))
	}
	return sa.index[i]
}

// LCP returns the length of the longest common prefix of the ith smallest suffix
// and the i-1 th smallest suffix.
func (sa *SuffixArrayX) LCP(i int) int {
	if i < 1 || i >= sa.n {
		panic("Illegal argument")
	}
	return sa.lcp(sa.index[i], sa.index[i-1])
}

// longest common prefix of text[i..n) and text[j..n)
func (sa *SuffixArrayX) lcp(i, j int) int {
	n := 0
	for i < sa.n && j < sa.n {
		if sa.text[i] != sa.text[j] {
			return n
		}
		i++
		j++
		n++
	}
	return n
}

// Select returns the ith smallest suffix as a string
func (sa *SuffixArrayX) Select(i int) string {
	if i < 0 || i >= sa.n {
		panic("Illegal argument")
	}
	m := sa.index[i]
	return string(sa.text[m:sa.n])
}

// Rank returns the number of suffixes strictly less than the query string.
// Note: rank(select(i)) == i
func (sa *SuffixArrayX) Rank(query string) int {
	lo, hi := 0, sa.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := sa.compare(query, sa.index[mid])
		if cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

// is query < text[i..n)
func (sa *SuffixArrayX) compare(query string, i int) int {
	m := len(query)
	j := 0
	for i < sa.n && j < m {
		if query[j] != sa.text[i] {
			return int(query[j]) - int(sa.text[i])
		}
		i++
		j++
	}
	if i < sa.n {
		return -1
	}
	if j < m {
		return +1
	}
	return 0
}
