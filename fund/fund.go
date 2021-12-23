package fund

// fundamental
// common structs and interfaces

type Iterator []interface{}

type Iterable interface {
	Iterator() Iterator
}