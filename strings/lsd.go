package strings

// Least-Significant-Digit First (LSD) radix sort for fixed length strings.
// It includes a method for sorting 32-bit integers, treating each integer as
// a 4-byte string. When N is large, this algorithm is 2-3x faster than the system sort.

// - Sort a string[] array of n extended ASCII strings (R=256), each of length w
// - Sort an int[] array of n 32-bit integers, treating each integer as a sequence
//   of w=4 bytes

// Uses extra space proportional to n+R

// Rearranges the array of w-character strings in ascending order.
func LSDSort(a []string) {
	// check that strings have fixed length
	w := len(a[0])
	for _, s := range a {
		if len(s) != w {
			panic("strings must have fixed length")
		}
	}

	n := len(a)
	R := 256 // extend ASCII alphabet size

	aux := make([]string, n)

	// sort by key-indexed counting on dth character
	for d := w - 1; d >= 0; d-- {
		//compute frequency counts
		count := make([]int, R+1)
		for i := 0; i < n; i++ {
			count[a[i][d]+1]++
		}

		// compute cumulates
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// move data
		for i := 0; i < n; i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}

		// copy back
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}
