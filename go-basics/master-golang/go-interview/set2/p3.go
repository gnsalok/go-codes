package main

import (
	"fmt"
)

func factorial(n int) int {
	if n < 1 {
		return 1
	} else {
		n = n * factorial(n-1)
		return n
	}
}

func main() {
	var i int
	i = 5
	result := factorial(i)
	fmt.Println(result)
}
