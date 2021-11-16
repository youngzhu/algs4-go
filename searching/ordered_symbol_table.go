package searching

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