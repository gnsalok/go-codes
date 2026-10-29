// This program fixes the typed-nil interface bug in the companion example.
package main

import "log"

// customError is just an empty struct.
type customError struct{}

// Error implements the error interface.
func (c *customError) Error() string {
	return "Find the bug."
}

// Returning the error interface directly lets the function return a genuinely
// nil interface when no error occurred.
func fail() ([]byte, error) {
	return nil, nil
}

func main() {
	var err error

	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?", err)
	}

	log.Println("No Error")
}
