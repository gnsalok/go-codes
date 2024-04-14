// Sample program to show see if the class can find the bug.
package main

import "log"

// customError is just an empty struct.
type customError struct{}

// Error implements the error interface.
// TIPs : Never use value semantic with Custom Error
func (c *customError) Error() string {
	return "Find the bug."
}

// fail returns nil values for both return types.
// Solution : Return error type, not *CustomError
// Bcz it is a concrete type
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
