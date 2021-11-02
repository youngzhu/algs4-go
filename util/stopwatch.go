package util

import (
	"time"
)

// A utility to measure the running time (wall clock) of a program
type Stopwatch struct {
	start time.Time
}

func NewStopwatch() Stopwatch {
	return Stopwatch{time.Now()}
}

func (s Stopwatch) ElapsedTime() float64 {
	now := time.Now()
	dt := s.start.Sub(now)
	return float64(dt.Milliseconds())/1000
}