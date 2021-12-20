package strings

// Substring Search.

// Rabin-Karp randomized fingerprint algorithm
func RabinKarpSearch(pattern, txt string) int {
	search := NewRabinKarp(pattern)
	return search.Search(txt)
}

// Knuth-Morris-Pratt algorithm
func KMPSearch(pattern, txt string) int {
	search := NewKMP(pattern)
	return search.Search(txt)
}

// The bad-character rule part of the Boyer-Moore algorithm.
func BoyerMooreSearch(pattern, txt string) int {
	search := NewBoyerMoore(pattern)
	return search.Search(txt)
}
