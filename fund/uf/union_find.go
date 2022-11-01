package uf

type UnionFind interface {
	// Union add connection between p and q
	Union(p, q int)

	// Find returns component identifier for p (o to n-1)
	Find(p int) int

	// Connected returns true if p and q are in the same component
	Connected(p, q int) bool

	// Count returns Number of components
	Count() int
}
