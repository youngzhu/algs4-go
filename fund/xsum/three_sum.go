package xsum

// Counts the number of triples in an array of N integers that sums to 0

// Returns the number of triples (i, j, k) with i<j<k
// such that a[i]+a[j]+a[k]==0
func ThreeSumCount(a []int) int {
	n := len(a)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}

func ThreeSumCountFast(a []int) int {

	sorter.SortInts(a)

	if containsDuplicates(a) {
		panic("contains duplicate integers")
	}

	n := len(a)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if k := binarySearch.Index(a, -a[i]-a[j]); k > j {
				count++
			}
		}
	}

	return count
}
