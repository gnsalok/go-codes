package main

import "fmt"

func main() {
	func(value int) {
		fmt.Println("Value is : ", value)
	}(42)

}
