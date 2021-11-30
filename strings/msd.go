package strings

// Sort an array of strings using MSD radix sort

// Rearranges the array of extended ASCII strings in ascending order
func MSDSort(a []string) {
	n := len(a)
	aux := make([]string, n)
	sort(a, aux, 0, n-1, 0)
}

const (
	R      = 256 // extended ADCII alphabet size
	cutoff = 15  // cutoff to insertion sort
)

// sort from a[lo] to a[hi], starting at the dth character
func sort(a, aux []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi < lo+cutoff {
		insertion(a, lo, hi, d)
		return
	}

	// compute frequency counts
	count := make([]int, R+2)
	for i := lo; i <= hi; i++ {
		c := a[i][d]
		count[c+2]++
	}

	// transform counts to indicies
	for r := 0; r < R; r++ {
		count[r+1] += count[r]
	}

	// distribute
	for i := lo; i <= hi; i++ {
		c := a[i][d]
		aux[count[c+1]] = a[i]
		count[c+1]++
	}

	// copy back
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// recursively sort for each character (excludes sentinel -1)
	for r := 0; r < R; r++ {
		sort(a, aux, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

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
