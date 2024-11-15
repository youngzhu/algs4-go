package searching

// If keys are small integers, we can use an array to implement a symbol table,
// by interpreting the key as an array index so that we can store the value
// associated with key i in array position i.

// Search algorithms that use hashing consist of two separate parts. The first
// step is to compute a hash function that transforms the search key into an
// array index. Ideally, different keys would map to different indices. This
// ideal is generally beyond our reach, so we have to face the possibility that
// two or more different keys may hash to the same array index. Thus, the second
// part of a hasing search is a collision-resolution process that deals with
// this situation.

// Hash functions.
// If we have an array that can hold M key-value pairs, then we need a function
// that can transform any given key into an index into that array: an integer in
// the range [0, M-1]. We seek a hash function that is both easy to compute and
// uniformly distributes the keys.

// We have 3 primary requirements in implementing a good hash function for a given key
// 1. It should be deterministic--equal keys must produce the same hash value
// 2. It should be efficient to compute
// 3. It should uniformly distribute the keys

// must be a power of 2
const initHashCapacity = 4

type HashSymbolTable interface {
	Put(key HashSTKey, value STValue)
	Get(key HashSTKey) STValue
	Delete(key HashSTKey)
	Contains(key HashSTKey) bool
	Keys() []HashSTKey
}

type HashSTKey interface {
	hashCode() int
}

type StringHashKey string

func (s StringHashKey) hashCode() int {
	hash := 0

	for _, v := range s {
		hash = 31*hash + int(v)
	}

	return hash
}
