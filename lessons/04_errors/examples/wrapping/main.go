package main

import (
	"errors"
	"fmt"
)

func process() error {
	println("Processing...")

	origErr := errors.New("An error occurred")
	return fmt.Errorf("process failed: %w", origErr)
}

func main() {
	if err := process(); err != nil {
		fmt.Println("Error:", err)
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Println("Unwrapped error:", unwrapped)
		}
	}
}
