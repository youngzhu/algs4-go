package sorting

// Entropy-optimal sorting.
// Arrays with large numbers of duplicate sort keys arise frequently in applications.
// In such applications, there is potential to reduce the time of the sort from
// linearithmic to linear.

// One straightforward idea is to partition the array into three parts, one each
// for items with keys smaller than, equal to, and larger than the paritioning
// item's key. Accomplishing this partitioning was a classical programming exercise
// popularized by E.W.Dijkstra as the Dutch National Flag problem, because it is
// like sorting an array with three possible key values, which might correspond to
// the three colors on the flag.

// Dijkstra's solution is based on a single left-to-right pass through the array
// that maintains a pointer lt such that a[lo..lt-1] is less than v, a pointer gt
// such that a[gt+1..hi] is greater than v, and a pointer i such that a[lt..i-1]
// are equal to v, and a[i..gt] are not yet examined.

func Quicksort3way(x Sortable) {
	// for (mostly) ordered items, shuffle is important
	shuffle(x)

	quicksort3way(x, 0, x.Len()-1)
}

// quicksort the subarray from x[lo] to x[hi] using 3-way partitioning
func quicksort3way(x Sortable, lo, hi int) {
	if hi <= lo {
		return
	}

	lt, gt := partition3way(x, lo, hi)

	// x[lo..lt-1] < a[lt..gt] < a[gt+1..hi]
	quicksort3way(x, lo, lt-1)
	quicksort3way(x, gt+1, hi)
}

// Starting with i equal to lo+1 we process a[i] using the 3-way compare to handle
// the three possible cases:
// a[i] less than v: exchange a[lt] with a[i] and increment both lt and i
// a[i] greater then v: exchange a[i] with a[gt] and decrement gt
// a[i] equal to v: increment i
func partition3way(x Sortable, lo, hi int) (int, int) {
	i := lo + 1
	lt, gt := lo, hi

	// x[lt] === x[lo]
	for i <= gt {
		if x.Less(i, lt) {
			x.Swap(i, lt)
			lt++
			i++
		} else if x.Less(lt, i) {
			x.Swap(i, gt)
			gt--
		} else {
			i++
		}
	}

	return lt, gt
}

type Quick3way struct{}

// Implements Sorter
func (s Quick3way) SortInts(x []int) {
	Quicksort3way(IntCompSlice(x))
}
func (s Quick3way) SortFloat64s(x []float64) {
	Quicksort3way(Float64CompSlice(x))
}
func (s Quick3way) SortStrings(x []string) {
	Quicksort3way(StringCompSlice(x))
}
