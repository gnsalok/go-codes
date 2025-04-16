package main

import "fmt"

func testDefer() {
	// Defer function to handle panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// Simulate a panic
	panic("Something went wrong!")
}

func main() {

	testDefer()
	fmt.Println("Program continues after panic")
	// Output: Error: Something went wrong!
	// Program continues after panic

}
