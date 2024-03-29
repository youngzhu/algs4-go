package uf

// WeightedQuickUnion Weighted quick-union algorithm (without path compression).
// Rather than arbitrarily connecting the second tree to the first for Union()
// in the quick-union algorithm, we keep track of the size of each tree and
// always connect the smaller tree to the larger.
type WeightedQuickUnion struct {
	parent []int // parent[i]: parent of i
	count  int   // number of components
	size   []int // size[i]: number of elements in subtree rooted at i
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &WeightedQuickUnion{parent, n, size}
}

func (qu *WeightedQuickUnion) Union(p, q int) {
	rootP := qu.Find(p)
	rootQ := qu.Find(q)
	if rootP == rootQ {
		return
	}

	// make smaller root point to larger one
	if qu.size[rootP] < qu.size[rootQ] {
		qu.parent[rootP] = rootQ
		qu.size[rootQ] += qu.size[rootP]
	} else {
		qu.parent[rootQ] = rootP
		qu.size[rootP] += qu.size[rootQ]
	}

	qu.count--
}

func (qu *WeightedQuickUnion) Find(p int) int {
	qu.validate(p)
	for p != qu.parent[p] {
		p = qu.parent[p]
	}
	return p
}

func (qu *WeightedQuickUnion) Count() int {
	return qu.count
}

func (qu *WeightedQuickUnion) Connected(p, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

func (qu *WeightedQuickUnion) validate(p int) {
	n := len(qu.parent)
	if p < 0 || p >= n {
		panic("invalid index")
	}
}
