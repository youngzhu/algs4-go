package uf

// QuickUnion is a Quick-union algorithm.
// It's based on the same data structure—the site-indexed id[] array—but
// it uses a different interpretation of the values that leads to more
// complicated structures. Specifically, the id[] entry for each site will
// be the name of another site in the same component (possibly itself).
// To implement Find() we start at the given site, follow its link to
// another site, follow that sites link to yet another site, and so forth,
// following links until reaching a root, a site that has a link to itself.
// Two sites are in the same component if and only if this process leads
// them to the same root. To validate this process, we need Union() to
// maintain this invariant, which is easily arranged: we follow links to
// find the roots associated with each of the given sites, then rename one
// of the components by linking one of these roots to the other.
type QuickUnion struct {
	parent []int // parent[i]: parent of i
	count  int   // number of components
}

func NewQuickUnion(n int) *QuickUnion {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &QuickUnion{parent, n}
}

func (qu *QuickUnion) Union(p, q int) {
	rootP := qu.Find(p)
	rootQ := qu.Find(q)
	if rootP == rootQ {
		return
	}
	qu.parent[rootP] = rootQ
	qu.count--
}

func (qu *QuickUnion) Find(p int) int {
	qu.validate(p)
	for p != qu.parent[p] {
		p = qu.parent[p]
	}
	return p
}

func (qu *QuickUnion) Count() int {
	return qu.count
}

func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

func (qu *QuickUnion) validate(p int) {
	n := len(qu.parent)
	if p < 0 || p >= n {
		panic("invalid index")
	}
}
