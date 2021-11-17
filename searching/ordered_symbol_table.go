package searching

// Ordered symbol tables.
// In typical applications, keys are Comparable objects, so the option exists of
// using a.CompareTo(b) to compare two keys a and b. Several symbol-table
// implementations take advantage of order among the keys thai is implied by
// Comparable to provide efficient implementations of Put() and Get() operations.
// More important, in such implementations, we can think of the symbol table as
// keeping the keys in order and consider a significantly expanded API that 
// difines numerous and useful operations involving relative key order.

type OrderedSymbolTable interface {
	Put(key OrderedSTKey, value STValue)
	Get(key OrderedSTKey) STValue
	Delete(key OrderedSTKey)
	Contains(key OrderedSTKey) bool
	Keys() []OrderedSTKey
}

// The key in ST
type OrderedSTKey interface {
	CompareTo(another STKey) int
	Equals(another STKey) bool
}

// key type
type (
	StringKey string
)

func (k StringKey) CompareTo(x STKey) int {
	kk := x.(StringKey)
	if k < kk {
		return -1
	} else if k > kk {
		return 1
	} else {
		return 0
	}
}

func (k StringKey) Equals(x STKey) bool {
	return k.CompareTo(x) == 0
}