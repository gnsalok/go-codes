// struct initialization

package main

import "fmt"

type Vector struct {
	x int
	y int
}

func main() {

	// use name of the variable while initializing
	vector := Vector{
		x: 20,
		y: 10,
	}
	fmt.Println(vector)
	fmt.Println(vector.x, vector.y)
}
