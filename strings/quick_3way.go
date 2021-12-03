package strings

// Sorting an array of strings using 3-way radix quicksort.
func Quicksort(a []string) {
	// random.shuffle(a)
	quicksort(a, 0, len(a)-1, 0)
}

// 3-way string quicksort a[lo..hi] starting at dth character
func quicksort(a []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi < lo+cutoff {
		insertion(a, lo, hi, d)
		return
	}

	lt, gt := lo, hi
	v := a[lo][d]
	i := lo + 1
	for i <= gt {
		t := a[i][d]
		if t < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if t > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}

	// a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi]
	quicksort(a, lo, lt-1, d)
	quicksort(a, lt, gt, d+1)
	quicksort(a, gt+1, hi, d)
}
