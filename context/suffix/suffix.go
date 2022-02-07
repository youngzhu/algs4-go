package suffix

func LongestRepeatedSubstring(text string) (lrs string) {
	sa := NewSuffixArray(text)
	lrs = ""
	for i := 1; i < sa.Length(); i++ {
		longestLen := sa.LCP(i)
		if longestLen > len(lrs) {
			lrs = sa.Select(i)[0:longestLen]
			//lrs = text[sa.Index(i) : sa.Index(i)+longestLen]
		}
	}

	return "\"" + lrs + "\""
}

// LongestCommonSubstring returns the longest common string of the two specified strings
func LongestCommonSubstring(s, t string) (lcs string) {
	suffix1 := NewSuffixArray(s)
	suffix2 := NewSuffixArray(t)

	// find the longest common substring by "merging" sorted suffixes
	lcs = ""
	i, j := 0, 0
	for i < suffix1.Length() && j < suffix2.Length() {
		p := suffix1.Index(i)
		q := suffix2.Index(j)
		x := lcpFrom(s, t, p, q)
		if len(x) > len(lcs) {
			lcs = x
		}
		if compare(s, t, p, q) < 0 {
			i++
		} else {
			j++
		}
	}

	return "\"" + lcs + "\""
}

// return the longest common prefix of suffix s[p...] and suffix t[q...]
func lcpFrom(s, t string, p, q int) string {
	n := min(len(s)-p, len(t)-q)
	for i := 0; i < n; i++ {
		if s[p+i] != t[q+i] {
			return s[p : p+i]
		}
	}
	return s[p : p+n]
}

// compare suffix s[p...] and suffix t[q...]
func compare(s, t string, p, q int) int {
	pp, qq := len(s)-p, len(t)-q
	n := min(pp, qq)
	for i := 0; i < n; i++ {
		if s[p+i] != t[q+i] {
			return int(s[p+i]) - int(t[q+i])
		}
	}

	if pp < qq {
		return -1
	} else if pp > qq {
		return +1
	} else {
		return 0
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
