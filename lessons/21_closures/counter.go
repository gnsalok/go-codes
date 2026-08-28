// Package main demonstrates how a closure can retain private state.
package main

import "fmt"

// NewCounter returns a closure that increments and returns its count each time
// it is called. The count starts at zero and belongs only to this closure.
func NewCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	first := NewCounter()

	fmt.Println("First counter:", first())
	fmt.Println("First counter:", first())
	fmt.Println("First counter:", first())

	// A second closure captures a different count variable.
	second := NewCounter()
	fmt.Println("Second counter:", second())
}
