package main

import "fmt"

func main() {

	const (
		CategoryBooks = iota
		CategoryHealth
		CategoryClothing
	)
	fmt.Println(CategoryHealth)
}
