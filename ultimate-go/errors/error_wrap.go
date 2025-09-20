package main

import (
	"errors"
	"fmt"
)

func process() error {
	println("Processing...")

	origErr := fmt.Errorf("An error occurred")
	return fmt.Errorf("process failed: %w", origErr)
}

func main() {
	if err := process(); err != nil {
		fmt.Println("Error:", err)
		// Unwrap the error
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Println("Unwrapped error:", unwrapped)
		}
	}
}
