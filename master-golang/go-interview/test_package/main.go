package main

import "timer"

func main() {
	clock := new(timer.StopWatch)
	clock.Start()
	if clock.running { // ILLEGAL
		// …
	}
}
