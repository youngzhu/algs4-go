package strings

// Brute force string search.
// Return offset of first match or -1 if no match
func BruteSearch1(pattern, text string) int {
	m, n := len(pattern), len(text)

	for i := 0; i <= n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if text[i+j] != pattern[j] {
				break
			}
		}
		if j == m {
			return i // found at offset i
		}
	}

	return -1 // not found
}

func BruteSearch2(pattern, text string) int {
	m, n := len(pattern), len(text)

	i, j := 0, 0
	for ; i < n && j < m; i++ {
		if text[i] == pattern[j] {
			j++
		} else {
			i -= j
			j = 0
		}
	}

	if j == m {
		return i - m // found
	} else {
		return -1 // not found
	}
}
