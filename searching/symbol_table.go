package searching

// SymbolTable (ST)
// The primary purpose of a symbol table is to associate a value with a key.
// The client can insert key-value pairs into the symbol table with the expectation
// of later being able to search for the value associated with a given key.
type SymbolTable interface {
	Put(key STKey, value STValue)
	Get(key STKey) STValue
	Delete(key STKey)
	Contains(key STKey) bool
	Keys() []STKey
}

// STKey the key in ST
type STKey interface{}

// STValue the value in ST
type STValue interface{}
