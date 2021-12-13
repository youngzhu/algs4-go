package uf

type UnionFind interface {
	// Add connection between p and q
	Union(p, q int)

	// Component identifier for p (o to n-1)
	Find(p int) int

	// Return true if p and q are in the same component
	Connected(p, q int) bool

	// Number of components
	Count() int
}
