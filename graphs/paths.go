package graphs

import (
	"github.com/youngzhu/algs4-go/fund"
	"github.com/youngzhu/algs4-go/sorting/pq"
)

// source vertex s
type Paths interface {
	// Is there a path from s to v
	HasPathTo(v int) bool
	// Path from s to v
	PathTo(v int) fund.Iterator
}

type Distance float64

func (d Distance) CompareTo(x pq.Item) int {
	xx := x.(Distance)
	if d < xx {
		return -1
	} else if d > xx {
		return 1
	} else {
		return 0
	}
}

const (
	DistanceInfinity = 10000.0

	// the same
	// DistanceZero = Distance(0.0)
	DistanceZero Distance = 0.0
)