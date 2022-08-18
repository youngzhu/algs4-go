package testutil

import (
	"math/rand"
	"time"
)

type Random struct {
	*rand.Rand
}

func NewRandom() *Random {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return &Random{r}
}

// Returns a random integer uniformly in [a, b)
func (r Random) UniformIntRange(a, b int) int {
	if b <= a {
		panic("invalid range")
	}

	return a + r.UniformIntN(b-a)
}

// Returns a random integer uniformly in [0, n)
func (r Random) UniformIntN(n int) int {
	return r.Intn(n)
}
