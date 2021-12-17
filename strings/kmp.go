package strings

// Knuth-Morris-Pratt algorithm.
type KMP struct {
	radix int     // the radix
	m     int     // length of pattern
	dfa   [][]int // the KMP automoton
}

func NewKMP(pattern string) *KMP {
	radix := 256
	m := len(pattern)

	// build DFA from pattern
	dfa := make([][]int, radix)
	for i := range dfa {
		dfa[i] = make([]int, m)
	}
	dfa[pattern[0]][0] = 1
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < radix; c++ {
			dfa[c][j] = dfa[c][x] // copy mismatch cases
		}
		dfa[pattern[j]][j] = j + 1 // set match case
		x = dfa[pattern[j]][x]     // update restart state
	}

	return &KMP{radix, m, dfa}
}

// Returns the index of the first occurrence of the pattern sting
// in the text string.
// If no such match, renturn -1
func (s *KMP) Search(txt string) int {
	// simulate operation of DFA on text
	n := len(txt)
	i, j := 0, 0
	for ; i < n && j < s.m; i++ {
		j = s.dfa[txt[i]][j]
	}

	if j == s.m {
		return i - s.m // found
	} else {
		return -1 // not found
	}
}
