package main

import (
	"fmt"
)

func main() {

	x := make([]int, 10, 100)
	x[0] = 12
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 112)
	fmt.Println(x)

	fmt.Println(len(x))
	fmt.Println(cap(x))

}
