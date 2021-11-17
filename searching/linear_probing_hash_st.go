package searching

// Hashing with linear probing.
// Another approach to implementing hashing is to store N key-value pairs in a
// hash table of size M > N, relying on empty entries in the table to help with
// collision resolution. Such methods are called open-addresding hashing methods.
// The simplest open-addressing method is called linear probing: when there is
// a collision (when we hash to a table index that is already occupied with a key
// different from the search key), then we just check the next entry in the table,
// by incrementing the index.
// There are three possible outcomes:
// 1. key equal to search key: search hit
// 2. empty position (null key at indexed position): search miss
// 3. key not equal to search key: try next entry

// As with separate chaining, the performance of open-addressing methods is
// dependent on the ratio α = N/M, but we interpret it differently. For
// separate chaining α is the average number of items per list and is generally
// larger than 1. For open addressing, α is the percentage of table positions 
// that are occupied; it must be less than 1. We refer to α as the *load factor*
// of the hash table.

type LinearProbingHashST struct {
	n int // number of key-value pairs in the symbol table
	m int // size of linear probing table
	keys []HashSTKey
	values []STValue
}

// Initializes an empty symbol table with defalut capacity
func NewLinearProbingHashST() *LinearProbingHashST {
	return NewLinearProbingHashSTN(initHashCapacity)
}

// Initializes an empty symbol table with the specified initial capacity
func NewLinearProbingHashSTN(capacity int) *LinearProbingHashST {
	keys := make([]HashSTKey, capacity)
	values := make([]STValue, capacity)

	return &LinearProbingHashST{m: capacity, keys: keys, values: values}
}

// Inserts the specified key-value pair into the symbol table, overwriting the
// old value with the new value if the symbol table already contains the specified key.
// 
// Deletes the specified key (and its associated value) from the symbol table if
// the specified value is nil
func (h *LinearProbingHashST) Put(key HashSTKey, value STValue) {
	if key == nil {
		panic("calls Put() with a nil key")
	}
	if value == nil {
		h.Delete(key)
		return
	}
	
	// double table size if 50% full
	if h.n >= h.m / 2 {
		h.resize(2*h.m)
	}

	i := h.hash(key)
	for ; h.keys[i] != nil; i = (i+1)%h.m {
		if h.keys[i] == key {
			h.values[i] = value
			return
		}
	}

	h.keys[i] = key
	h.values[i] = value
	h.n++
}

// Returns the value associated with the given key
func (h *LinearProbingHashST) Get(key HashSTKey) STValue {
	if key == nil {
		panic("calls get() with a nil key")
	}
	for i := h.hash(key); h.keys[i] != nil; i = (i+1)%h.m {
		if h.keys[i] == key {
			return h.values[i]
		}
	}
	return nil
}

// Reomoves the specified key and its associated value from the symbol table
// (if the key is in this symbol table)
func (h *LinearProbingHashST) Delete(key HashSTKey) {
	if key == nil {
		panic("calls Delete() with a nil key")
	}

	if !h.Contains(key) {
		return
	}

	// find position i of key
	i := h.hash(key)
	for h.keys[i] != key {
		i = (i+1)%h.m
	}
	
	// delete key and associated value
	h.keys[i] = nil
	h.values[i] = nil

	// rehash all keys in same cluster
	i = (i+1)%h.m
	for h.keys[i] != nil {
		// delete keys[i] and values[i] and reinsert
		k, v := h.keys[i], h.values[i]
		h.keys[i], h.values[i] = nil , nil
		h.n--
		h.Put(k, v)

		i = (i+1)%h.m
	}
	 
	h.n--

	// halve table size if it's 1/8 full or less
	if h.n > 0 && h.n <= h.m/8 {
		h.resize(h.m/2)
	}
}

// Returns all keys in the symbol table
func (h *LinearProbingHashST) Keys() []HashSTKey {
	if h.IsEmpty() {
		panic("The ST is empty")
	}

	keys := make([]HashSTKey, h.Size())
	
	for i, n := 0, 0; i < h.m; i++ {
		if (h.keys[i] != nil) {
			keys[n] = h.keys[i]
			n++
		}
	}

	return keys
}

// Does this ST contains the given key?
func (h *LinearProbingHashST) Contains(key HashSTKey) bool {
	if key == nil {
		panic("argument to Contains() is nil")
	}
	return h.Get(key) != nil
}

// Returns the number of key-value pairs in this symbol table
func (h *LinearProbingHashST) Size() int {
	return h.n
}

// Returns ture if this ST is empty
func (h *LinearProbingHashST) IsEmpty() bool {
	return h.Size() == 0
}

// resize the hash table to have the given number of chains,
// rehashing all of the keys
func (h *LinearProbingHashST) resize(capacity int) {
	temp := NewLinearProbingHashSTN(capacity)

	for i := 0; i < h.m; i++ {
		if h.keys[i] != nil {
			temp.Put(h.keys[i], h.values[i])
		}
	}

	h.m = temp.m
	h.keys = temp.keys
	h.values = temp.values
}

// hash function for keys
// returns value between 0 and m-1
func (h *LinearProbingHashST) hash(key HashSTKey) int {
	return (key.hashCode() & 0x7fffffff) % h.m
}