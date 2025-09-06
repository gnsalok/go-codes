package main

import "time"

// IMP : Need to organize you rGo Codebase

// A StopWatch is a simple clock utility.
// Its zero value is an idle clock with 0 total time.
// If it start is caps S means your can export in other package.
type StopWatch struct {
	start   time.Time
	total   time.Duration
	running bool
}

// Start turns the clock on.
func (s *StopWatch) Start() {
	if !s.running {
		s.start = time.Now()
		s.running = true
	}
}
