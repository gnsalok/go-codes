package main

import (
	"fmt"
)

func main() {

	// Compiler can perform implicit conversion on Untyped constant
	// Untyped Constant
	const PI = 3.14 * 3 // Constant of kind

	// Typed Constant
	const fixInt int = 3 // type : int

	fmt.Println(fixInt)
	fmt.Println(PI)
}
