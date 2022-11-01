package uf

// QuickFindUF Quick-find algorithm.
// It maintains the invariant that p and q are connected if and only if
// id[p] = id[q]. In other words, all sites in a component must have the
// same value in id[]
type QuickFindUF struct {
	id    []int // id[i]: component identifier of i
	count int   // number of components
}

// NewQuickFindUF returns an empty union-find data structure with n elements (0...n-1)
// Initially, each element is in its own set.
func NewQuickFindUF(n int) *QuickFindUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &QuickFindUF{id, n}
}

func (qf *QuickFindUF) Find(p int) int {
	qf.validate(p)
	return qf.id[p]
}

func (qf *QuickFindUF) Union(p, q int) {
	qf.validate(p)
	qf.validate(q)

	// needed for correctness to reduce the number of array accesses
	pID := qf.id[p]
	qID := qf.id[q]

	// p and q are already in the same component
	if pID == qID {
		return
	}

	n := len(qf.id)
	for i := 0; i < n; i++ {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}

	qf.count--
}

func (qf *QuickFindUF) Count() int {
	return qf.count
}

func (qf *QuickFindUF) Connected(p, q int) bool {
	qf.validate(p)
	qf.validate(q)
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFindUF) validate(p int) {
	n := len(qf.id)
	if p < 0 || p >= n {
		panic("invalid index")
	}
}
