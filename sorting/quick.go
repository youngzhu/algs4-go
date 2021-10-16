package sorting

import "math/rand"

// Quicksort is popular because it is not difficult to implement, works well for
// a variety of different kinds of input data, and is substantially faster than
// an other sorting method in typical applications. It is in-place (uses only a
// small auxiliary stack), requires time proportional to NlogN on the average
// to sort N items, and has an extremely short inner loop.

// The basic algorithm
// Quicksort is a divide-and-conquer method for soring. It works by partitioning
// an array into two parts, then soring the parts independently.

// The crux of the method is the partitioning process, which rearranges the 
// array to make the following three conditions hold:
// 1. The entry a[j] is in its final place in the array, for some j
// 2. No entry in a[lo] through a[j-1] is greater than a[j]
// 3. No entry in a[j+1] through a[hi] is less than a[j]
// We achieve a complete sort by partitioning, then recursively applying the
// method to the subarrays. It is a randomized algorithm, because it randomly
// shuffles the array before sorting it.

// Partitioning.
// To complete the implementation, we need to implement the partitioning. We
// use the following general strategy:
// First, we arbitrarily choose a[lo] to be the partitioning item, the one that
// will go into its final position.
// Next, we scan from the left end of the array until we find an entry that is
// greater than (or equal to) the partitioning item, and we scan from the right
// end of the array util we find an entry less than (or equal to) the partitioning
// item.
// The two items that stopped the scans are out of place in the final partitioned
// array, so we exchange them. When the scan indices cross, all that we need to
// do to complete the partitioning process is to exchange the partitioning item
// a[lo] with the rightmost entry of the left subarray (a[j]) and return its index j.

func Quicksort(x Comparable) {
	// for (mostly) ordered items, shuffle is important
	// see cmd/sorting/sort_compare.go
	shuffle(x) 

	quicksort(x, 0, x.Len()-1)
}

func shuffle(x Comparable) {
	rand.Shuffle(x.Len(), func (i, j int) {
		x.Swap(i, j)
	})
}

// quicksort the subarray from x[lo] to x[hi]
func quicksort(x Comparable, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(x, lo, hi)
	quicksort(x, lo, j-1)
	quicksort(x, j+1, hi)
}

// partition the subarray x[lo..hi]
// so that x[lo..j-1] <= x[j] <= x[j+1..hi]
// and return the index j
func partition(x Comparable, lo, hi int) int {
	i, j := lo+1, hi

	for {
		// find item on lo to swap
		for ; x.Less(i, lo); i++ {
			if i == hi {
				break
			}
		}
		// find item on hi to swap
		for ; x.Less(lo, j); j-- {
			if j == lo {
				break //redundant since x[lo] acts as sentinel
			}
		}

		// check if pointers cross
		if i >= j {
			break
		}

		x.Swap(i, j)
	}

	// put partitioning item at x[j]
	x.Swap(lo, j)

	// now, x[lo..j-1] <= x[j] <= x[j+1..hi]
	return j
}

type Quick struct{}

// Implements Sorter
func (s Quick) SortInts(x []int) {
	Quicksort(IntCompSlice(x))
}
func (s Quick) SortFloat64s(x []float64) {
	Quicksort(Float64CompSlice(x))
}
func (s Quick) SortStrings(x []string) {
	Quicksort(StringCompSlice(x))
}

// Implementation details. There are several subtle issues with respect to
// implementing quicksort that are reflected in this code and worthy of mention.

// 1. Partitioning inplace.
// If we use an extra array, partitioning is easy to implement, but not so much
// easier that it is worth the extra cost of copying the partitioned version
// back into the original.

// 2. Staying in bounds.
// If the smallest item or the largest item in the array is the partitioning item,
// we have to take care that the pointers do not run off the left or right ends
// of the array, respctively.

// 3. Preserving randomness.
// The random shuffle put the array in random order. Since it treats all items in
// the subarrays uniformly, this implemention has the property that its two 
// subarrays are also in random order. This fact is crucial to the algorithm's
// predictability. An alternate way to preserve randomness is to choose a random
// item for partitioning within partition().

// 4. Terminating the loop.
// Properly testing whether the pointers have crossed is a bit tricker than it
// might seem at first glance. A common error is to fail to take into account
// that the array might contain other keys with the same value as the partitioning item.

// 5. Handling items with keys equal to the partitioning item's key.
// It is best to stop the left scan of items with keys greater than or equal to
// the partitioning item's key and the right scan for items less than or equal to
// the partitioning item's key. Even though this policy might seem to create
// unnecessary exchanges involving items wiht keys equal to the partitioning 
// item's key, it is crucial to avoiding quadratic running time in certain
// typical applications.

// 6. Terminating the recursion.
// A commmon mistake in implementing quicksort involves not ensuring that one
// item is always put into position, then falling into an infinite recursive
// loop when the partitioning item happens to be the largest or smallest item
// in the array.
