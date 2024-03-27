package main

import "fmt"

func main() {
	ans := fact(5)
	fmt.Println("Factorial is :", ans)
}

// TRUST THE FUNCTION : RECURSIVE LEAP OF FAITH
func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
