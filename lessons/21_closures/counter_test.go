package main

import "testing"

func TestCounterRemembersState(t *testing.T) {
	counter := NewCounter()

	for want := 1; want <= 3; want++ {
		if got := counter(); got != want {
			t.Fatalf("counter() = %d, want %d", got, want)
		}
	}
}

func TestCountersHaveIndependentState(t *testing.T) {
	first := NewCounter()
	second := NewCounter()

	first()
	first()

	if got := second(); got != 1 {
		t.Fatalf("new counter returned %d, want 1", got)
	}
}
