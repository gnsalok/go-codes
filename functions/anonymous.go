package main

import "fmt"

func main() {

	func(n int) {
		fmt.Println("Hello world", n, "times")
	}(10)
}
