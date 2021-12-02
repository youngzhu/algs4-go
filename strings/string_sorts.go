package strings

const (
	R      = 256 // extended ASCII alphabet size
	cutoff = 15  // cutoff to insertion sort
)

// insertion sort a[lo..hi], starting at dth character
func insertion(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(a[j], a[j-1], d); j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

// is s1 less than s2, starting at dth character
func less(s1, s2 string, d int) bool {
	len1, len2 := len(s1), len(s2)
	for i := d; i < min(len1, len2); i++ {
		if s1[i] < s2[i] {
			return true
		}
		if s1[i] > s2[i] {
			return false
		}
	}
	return len1 < len2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
