package searching

import "strings"

// Ordered symbol tables.
// In typical applications, keys are Comparable objects, so the option exists of
// using a.CompareTo(b) to compare two keys a and b. Several symbol-table
// implementations take advantage of order among the keys that is implied by
// Comparable to provide efficient implementations of Put() and Get() operations.
// More important, in such implementations, we can think of the symbol table as
// keeping the keys in order and consider a significantly expanded API that
// defines numerous and useful operations involving relative key order.

type OrderedSymbolTable interface {
	Put(key OSTKey, value STValue)
	Get(key OSTKey) STValue
	Delete(key OSTKey)
	Contains(key OSTKey) bool
	Keys() []OSTKey
}

// OSTKey The key in ordered symbol tables
type OSTKey interface {
	CompareTo(another OSTKey) int
	Equals(another OSTKey) bool
}

type StringKey string

func (k StringKey) CompareTo(x OSTKey) int {
	s1 := string(k)
	s2 := string(x.(StringKey))
	return strings.Compare(s1, s2)
}

func (k StringKey) Equals(x OSTKey) bool {
	s1 := string(k)
	s2 := string(x.(StringKey))
	return s1 == s2
}
