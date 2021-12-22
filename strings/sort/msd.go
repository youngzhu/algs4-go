package sort

import "github.com/youngzhu/algs4-go/strings"

// Sort an array of strings using MSD radix sort

// Rearranges the array of extended ASCII strings in ascending order
func MSDSort(a []string) {
	n := len(a)
	aux := make([]string, n)
	sort(a, aux, 0, n-1, 0)
}

// sort from a[lo] to a[hi], starting at the dth character
func sort(a, aux []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi < lo+cutoff {
		insertion(a, lo, hi, d)
		return
	}

	// compute frequency counts
	count := make([]int, strings.R+2)
	for i := lo; i <= hi; i++ {
		c := a[i][d]
		count[c+2]++
	}

	// transform counts to indicies
	for r := 0; r < strings.R; r++ {
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
	for r := 0; r < strings.R; r++ {
		sort(a, aux, lo+count[r], lo+count[r+1]-1, d+1)
	}
}
