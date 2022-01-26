package suffix

func LongestRepeatedSubstring(text string) string {
	sa := NewSuffixArray(text)
	lrs := ""
	for i := 1; i < sa.Length(); i++ {
		longestLen := sa.LCP(i)
		if longestLen > len(lrs) {
			lrs = sa.Select(i)[0:longestLen]
		}
	}

	return lrs
}
