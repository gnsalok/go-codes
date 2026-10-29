package main

import "fmt"

func main() {
	// An untyped constant is given a concrete type only when its context
	// requires one. A typed constant always keeps its declared type.
	const circumferenceRatio = 3.14 * 3
	const fixedInteger int = 3

	fmt.Println(fixedInteger)
	fmt.Println(circumferenceRatio)
}
