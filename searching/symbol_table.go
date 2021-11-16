package searching

// Symbol table (ST)
// The primary prupose of a symbol table is to associate a value with a key.
// The client can insert key-value pairs into the symbol table with the expection
// of later being able to search for the value associated with a given key.

type SymbolTable interface {
	Put(key STKey, value STValue)
	Get(key STKey) STValue
	Delete(key STKey)
	Contains(key STKey) bool
	Keys() []STKey
}

// The key in ST
type STKey interface {}

// The value in ST
type STValue interface{}