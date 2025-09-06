package main

import "fmt"

func main() {

	// Slice will give mechanical sympathy to HW
	// Slice provides predictable access patter ; Row Major contiguous memory allocation and access

	friends := []string{"Ani", "Donnie", "tonie", "sonie"}

	/*
		Very Important:
		- When we are ranging over slice i.e. friends and getting value "v" - we are working on copy (value semantic) of the data.
		- OTOH, when we do "for i := range friends" - we are working on pointer semantic, there modifying friend[:2] inside
		  loop, modify the actual array.
	*/

	// This will work fine
	// Value semantic
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("%v\n", v)
	}

	// out of range as L#25 modifies the actual slice
	friends = []string{"Ani", "Donnie", "tonie", "sonie"}

	for i := range friends {
		friends = friends[:2]
		fmt.Printf("%v\n", friends[i])
	}

}
