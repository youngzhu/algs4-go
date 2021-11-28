package xsum

// Returns the number of distinct pairs (i, j)
// such that a[i]+a[j]=0
func TwoSumCount(a []int) int {
	n := len(a)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i]+a[j] == 0 {
				count++
			}

		}
	}

	return count
}

func TwoSumCountFast(a []int) int {

	sorter.SortInts(a)

	if containsDuplicates(a) {
		panic("contains duplicate integers")
	}

	count := 0

	n := len(a)
	for i := 0; i < n; i++ {
		if j := binarySearch.Index(a, -a[i]); j > i {
			count++
		}
	}

	return count
}
