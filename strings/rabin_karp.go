package strings

import "crypto/rand"

// Rabin-Karp randomized fingerprint algorithm.
type RabinKarp struct {
	patternHash int // pattern hash value
	m           int // pattern length
	q           int // a large prime, small enough to avoid overflow
	r           int // radix
	rm          int // r^(m-1)%q
}

func NewRabinKarp(pattern string) *RabinKarp {
	r := 256
	m := len(pattern)
	q := longRandomPrime()

	// precompute r^(m-1)%q for use in removing leading digit
	rm := 1
	for i := 1; i <= m-1; i++ {
		rm = (r * rm) % q
	}

	rk := &RabinKarp{m: m, q: q, r: r, rm: rm}

	rk.patternHash = rk.hash(pattern)

	return rk
}

// compute hash for key[0..m-1]
func (rk *RabinKarp) hash(key string) int {
	h := 0
	for j := 0; j < rk.m; j++ {
		h = (rk.r*h + int(key[j])) % rk.q
	}
	return h
}

// Returns the index of the first occurrence of the pattern sting
// in the text string.
// If no such match, renturn -1
func (rk *RabinKarp) Search(txt string) int {
	n := len(txt)
	if n < rk.m {
		return -1
	}

	txtHash := rk.hash(txt)

	// check for match at offset 0
	if txtHash == rk.patternHash {
		return 0
	}

	// check for hash match
	// if hash match, check for exact match
	for i := rk.m; i < n; i++ {
		// remove leading digit, add trailing digit, check for match
		k := int(txt[i-rk.m])
		txtHash = (txtHash + rk.q - rk.rm*k%rk.q) % rk.q
		txtHash = (txtHash*rk.r + int(txt[i])) % rk.q

		// match
		if txtHash == rk.patternHash {
			return i - rk.m + 1
		}
	}

	// no match
	return -1
}

// a random 31-bit prime
func longRandomPrime() int {
	prime, _ := rand.Prime(rand.Reader, 31)
	return int(prime.Int64())
}
