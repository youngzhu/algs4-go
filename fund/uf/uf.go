package uf

// Weighted quick-union by rank with path compression by halving.
type UF struct {
	parent []int  // parent[i]: parent of i
	rank   []byte // rank[i]: rank of subtree rooted at i (never more than 31)
	count  int    // number of components
}

// New an empty union-find data strcture with n elements (0...n-1)
// Initially, each elements is in its own set.
func NewUF(n int) *UF {
	if n < 0 {
		panic("n must be positive")
	}

	parent := make([]int, n)
	rank := make([]byte, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UF{parent, rank, n}
}

// Merges the set containing element p with the set containg element q
func (u *UF) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}

	// make root of smaller rank point to root of larger rank
	if u.rank[rootP] < u.rank[rootQ] {
		u.parent[rootP] = rootQ
	} else if u.rank[rootP] > u.rank[rootQ] {
		u.parent[rootQ] = rootP
	} else {
		u.parent[rootQ] = rootP
		u.rank[rootP]++
	}

	u.count--
}

// Returns the canonical element of the set containing element p
func (u *UF) Find(p int) int {
	u.validate(p)
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]] // path compression by halving
		p = u.parent[p]
	}
	return p
}

// Returns the number of sets
func (u *UF) Count() int {
	return u.count
}

// Returns true if the two elements are in the same set
func (u *UF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

// validate that p is a valid index
func (u UF) validate(p int) {
	n := len(u.parent)
	if p < 0 || p >= n {
		panic("invalid index")
	}
}
