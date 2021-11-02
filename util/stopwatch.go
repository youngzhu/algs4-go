package util

import "time"

// A utility to measure the running time (wall clock) of a program
type Stopwatch struct {
	start time.Time
}

func NewStopwatch() Stopwatch {
	return Stopwatch{time.Now()}
}

// Returns the elapsed time (in seconds) since the stopwatch was created
func (s Stopwatch) ElapsedTime() float64 {
	now := time.Now()
	dt := now.Sub(s.start)
	return float64(dt.Milliseconds())/1000
}