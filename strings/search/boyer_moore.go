package search

// The bad-character rule part of the Boyer-Moore algorithm.
// It does not implement the strong good suffix rule.
type BoyerMoore struct {
	radix   int
	right   []int // the bad-character skip array
	pattern string
}

func NewBoyerMoore(pattern string) *BoyerMoore {
	radix := 256

	// position of rightmos occurrence of c in the pettern
	right := make([]int, radix)
	for i := 0; i < radix; i++ {
		right[i] = -1
	}
	for j := 0; j < len(pattern); j++ {
		right[pattern[j]] = j
	}

	return &BoyerMoore{radix, right, pattern}
}

// Returns the index of the first occurrence of the pattern sting
// in the text string.
// If no such match, renturn -1
func (s *BoyerMoore) Search(txt string) int {
	m, n := len(s.pattern), len(txt)
	var skip int
	for i := 0; i <= n-m; i += skip {
		skip = 0
		for j := m - 1; j >= 0; j-- {
			if s.pattern[j] != txt[i+j] {
				k := int(txt[i+j])
				skip = max(1, j-s.right[k])
				break
			}
		}
		if skip == 0 {
			return i // found
		}
	}
	return -1 // not found
}

func max(i, j int) int {
	if j > i {
		return j
	}
	return i
}
