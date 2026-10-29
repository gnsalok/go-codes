package main

import "fmt"

func testDefer() {
	defer func() {
		if recovered := recover(); recovered != nil {
			fmt.Println("Error:", recovered)
		}
	}()

	panic("Something went wrong!")
}

func main() {
	testDefer()
	fmt.Println("Program continues after panic")
}
