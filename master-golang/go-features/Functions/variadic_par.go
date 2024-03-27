package main

import (
	"fmt"
)

func addNum(n ...int) int {
	sum := 0

	for i := 0; i < len(n); i++ {
		sum = sum + n[i]
	}

	return sum
}

func main() {

	sum := addNum(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum of numbers is ", sum)
}
