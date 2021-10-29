package searching

// Binary search in an ordered array.
// BinarySearchST implements the ordered symbol table API. The underlying
// data structure is two parallel array, with the keys kept in order. The heart
// of the implemention is the rank() method, which returns the number
// of keys smaller than a given key. For get(), the rank tells us precisely
// where the key is to be found if it is in the table (and, if it is not
// there, that it is not in the table). For put(), the rank tells us precisely
// where to update the value when the key is in the table, and precisely where
// to put the key when the key is not in the table. We move all large keys
// over one position to make room (working from back to front) and insert the 
// given key and value into the proper position in their respective arrays.

// The reason that we keep keys in an ordered array is so that we can use 
// array indexing to dramatically reduce the number of compares required 
// for each search, using a venerable classic algorithm known as binary search.
// The basic idea is simple: we maintain indices into the sorted key
// array that delimit the subarrsy that might contain the search key.
// To search, we compare the search key against the in the middle of the
// subarray. If the search key is less than the key in the middle, we 
// search in the left half of the subarray; if the search key is greater
// than the key in the middle we search in the right half of the subarray;
// otherwise the key in the middle is equal to the search key.

const initCapacity = 2

type BinarySearchST struct {
	keys []STKey
	values []STValue
	n int
}

func NewBinarySearchST() *BinarySearchST {
	return NewBinarySearchSTN(initCapacity)
}

// Initializes an empty symbol table with the specified initial capacity
func NewBinarySearchSTN(n int) *BinarySearchST {
	keys := make([]STKey, n)
	values := make([]STValue, n)
	return &BinarySearchST{keys, values, 0}
}

func (st *BinarySearchST) resize(newCap int) {
	if newCap < st.n {
		return
	}

	newKeys := make([]STKey, newCap)
	copy(newKeys, st.keys)
	st.keys = newKeys

	newValues := make([]STValue, newCap)
	copy(newValues, st.values)
	st.values = newValues
} 

// The heart of the implemention
// Returns the number of keys strictly less than the given key
func (st *BinarySearchST) Rank(key STKey) int {
	if key == nil {
		panic("argument to Rank() is nil")
	}

	lo, hi := 0, st.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := key.CompareTo(st.keys[mid])
		if cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return lo
}

// Returns the value associated with the given key 
func (st *BinarySearchST) Get(key STKey) STValue {
	if key == nil {
		panic("argument to Get() is nil")
	}
	if st.IsEmpty() {
		return nil
	}
	r := st.Rank(key)
	if r < st.n && key.Equals(st.keys[r]) {
		return st.values[r]
	}
	return nil
}

// Inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol
// already contains the specified key.
// Deletes the specified key (and its associated value) from 
// this symbol table if the specified value is nil
func (st *BinarySearchST) Put(key STKey, value STValue) {
	if key == nil {
		panic("first argument to Put() is nil")
	}
	if value == nil {
		st.Delete(key)
		return
	}

	r := st.Rank(key)

	// key is alread in table
	if r < st.n && key.Equals(st.keys[r]) {
		st.values[r] = value
		return
	}

	// insert new key-value pair
	if st.n == len(st.keys) {
		st.resize(2*st.n)
	}

	for i := st.n; i > r; i-- {
		st.keys[i] = st.keys[i-1]
		st.values[i] = st.values[i-1]
	}
	st.keys[r], st.values[r] = key, value
	st.n++
}

// Removes the specified key and associated value
// (if the key is in the symbol table)
func (st *BinarySearchST) Delete(key STKey) {
	if key == nil {
		panic("argument to Delete() is nil")
	}
	if st.IsEmpty() {
		return
	}

	// compute rank
	r := st.Rank(key)

	// key not in table
	if r == st.n || !key.Equals(st.keys[r]) {
		return
	}

	for i := r; i < st.n-1; i++ {
		st.keys[i] = st.keys[i+1]
		st.values[i] = st.values[i+1]
	}
	st.n--

	// resize if 1/4 full
	if st.n > 0 && st.n == len(st.keys)/4 {
		st.resize(len(st.keys)/2)
	}
}

// Returns all keys in the symbol table
func (st BinarySearchST) Keys() []STKey {
	return st.keys[:st.n]
}

func (st BinarySearchST) IsEmpty() bool {
	return st.n == 0
}