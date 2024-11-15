package searching

// Hashing with separate chaining.
// A hash function converts keys into array indices. The second component of a
// hashing algorithm is collision resolution: a strategy for handling the case
// when two or more keys to be inserted hash to the same index. A straightforward
// approach to collision resolution is to build, for each of the M array indices,
// a linked list of the key-value pairs whose keys hash to that index. The basic
// idea is to choose M to be sufficiently large that the lists are sufficiently
// short to enable efficient search through a two-step process: hash to find the
// list that could contain the key, then sequentially search through that list
// for the key.

// SeparateChainingHashST implements a symbol table with a separate-chaining hash
// table. It maintains an array of SequentialSearchST and implements Get() and Put()
// by computing a hash function to choose which SequentialSearchST can contain the
// key and then using Get() and Put() from SequentialSearchST to complete either job.

type SeparateChainingHashST struct {
	n  int                   // number of key-value pairs
	m  int                   // hash table size
	st []*SequentialSearchST // array of linked-list symbol tables
}

func NewSeparateChainingHashST() *SeparateChainingHashST {
	return NewSeparateChainingHashSTN(initHashCapacity)
}

func NewSeparateChainingHashSTN(m int) *SeparateChainingHashST {
	st := make([]*SequentialSearchST, m)
	for i := 0; i < m; i++ {
		st[i] = NewSequentialSearchST()
	}

	return &SeparateChainingHashST{0, m, st}
}

// Returns the value associated with the given key
func (h *SeparateChainingHashST) Get(key HashSTKey) STValue {
	if key == nil {
		panic("calls get() with a nil key")
	}
	i := h.hash(key)
	return h.st[i].Get(key)
}

// Inserts the specified key-value pair into the symbol table, overwriting the
// old value with the new value if the symbol table already contains the specified key.
//
// Deletes the specified key (and its associated value) from the symbol table if
// the specified value is nil
func (h *SeparateChainingHashST) Put(key HashSTKey, value STValue) {
	if key == nil {
		panic("calls Put() with a nil key")
	}
	if value == nil {
		h.Delete(key)
		return
	}

	// double table size if average length of list >= 10
	if h.n >= 10*h.m {
		h.resize(2 * h.m)
	}

	i := h.hash(key)
	if !h.st[i].Contains(key) {
		h.n++
	}
	h.st[i].Put(key, value)
}

// Reomoves the specified key and its associated value from the symbol table
// (if the key is in this symbol table)
func (h *SeparateChainingHashST) Delete(key HashSTKey) {
	if key == nil {
		panic("calls Delete() with a nil key")
	}

	i := h.hash(key)
	if h.st[i].Contains(key) {
		h.n--
		h.st[i].Delete(key)
	}

	// halve table size if average length of list <= 2
	if h.m > initHashCapacity && h.n <= 2*h.m {
		h.resize(h.m / 2)
	}
}

// Returns all keys in the symbol table
func (h *SeparateChainingHashST) Keys() []HashSTKey {
	if h.IsEmpty() {
		panic("The ST is empty")
	}

	keys := make([]HashSTKey, h.Size())

	for i, n := 0, 0; i < h.m; i++ {
		for _, v := range h.st[i].Keys() {
			keys[n] = v.(HashSTKey)
			n++
		}
	}

	return keys
}

// Does this BST contains the given key?
func (h *SeparateChainingHashST) Contains(key HashSTKey) bool {
	if key == nil {
		panic("argument to Contains() is nil")
	}
	return h.Get(key) != nil
}

// Returns the number of key-value pairs in this symbol table
func (h *SeparateChainingHashST) Size() int {
	return h.n
}

// Returns ture if this ST is empty
func (h *SeparateChainingHashST) IsEmpty() bool {
	return h.Size() == 0
}

// resize the hash table to have the given number of chains,
// rehashing all of the keys
func (h *SeparateChainingHashST) resize(chains int) {
	temp := NewSeparateChainingHashSTN(chains)

	for i := 0; i < h.m; i++ {
		for _, key := range h.st[i].Keys() {
			temp.Put(key.(HashSTKey), h.st[i].Get(key))
		}
	}

	h.m = temp.m
	h.n = temp.n
	h.st = temp.st
}

// hash function for keys
// returns value between 0 and m-1
func (h *SeparateChainingHashST) hash(key HashSTKey) int {
	return (key.hashCode() & 0x7fffffff) % h.m
}
