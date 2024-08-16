package testutil

import "time"

// Stopwatch A utility to measure the running time (wall clock) of a program
type Stopwatch struct {
	start time.Time
}

func NewStopwatch() Stopwatch {
	return Stopwatch{time.Now()}
}

// ElapsedTime returns the elapsed time (in seconds) since the stopwatch was created
func (s Stopwatch) ElapsedTime() float64 {
	now := time.Now()
	dt := now.Sub(s.start)
	return float64(dt.Milliseconds()) / 1000
}
