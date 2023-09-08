package fund

type BinarySearch struct{}

func NewBinarySearch() BinarySearch {
	return BinarySearch{}
}

// Index return the index of the specified key in the specified array
// otherwise return -1
// (the array must be sorted in ascending order.)
func (bs BinarySearch) Index(a []int, key int) int {
	lo, hi := 0, len(a)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
