package strings

import "log"

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

/* Didn't work. Don't know why */

// Rearranges the array of 32-bit integers in ascending order
func LSDSortInts(a []int) {
	bits := 32 // each int is 32 bits
	bitsPerByte := 8
	R := 1 << bitsPerByte   // each bytes is between 0 and 255
	mask := R - 1           // 0xFF
	w := bits / bitsPerByte // each int is 4 byte

	log.Println(w)

	n := len(a)
	aux := make([]int, n)

	for d := 0; d < w; d++ {
		// compute frequency counts
		count := make([]int, R+1)
		for i := 0; i < n; i++ {
			c := (a[i] >> bitsPerByte * d) & mask
			count[c+1]++
		}

		// compute cumulates
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// for most significant byte, 0x80-0xFF come before 0x00-0x7F
		if d == w-1 {
			shift1 := count[R] - count[R/2]
			shift2 := count[R/2]
			for r := 0; r < R/2; r++ {
				count[r] += shift1
			}
			for r := R / 2; r < R; r++ {
				count[r] -= shift2
			}
		}

		// move data
		for i := 0; i < n; i++ {
			c := (a[i] >> bitsPerByte * d) & mask
			aux[count[c]] = a[i]
			count[c]++
		}

		// copy back
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}
