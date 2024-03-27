package main

import (
	"fmt"
)

func addOne() func() int {
	var x int
	return func() int {
		x++
		return x + 1
	}
}

func main() {

	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5
}
