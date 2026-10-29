package main

import "fmt"

func main() {
	friends := []string{"Ani", "Donnie", "tonie", "sonie"}

	// range copies the original slice header before the loop starts. The values
	// therefore remain available even though friends is shortened in the body.
	for _, friend := range friends {
		friends = friends[:2]
		fmt.Println(friend)
	}

	// This loop intentionally panics when i reaches 2: the range still produces
	// indexes from the original length, but indexing uses the shortened slice.
	friends = []string{"Ani", "Donnie", "tonie", "sonie"}

	for i := range friends {
		friends = friends[:2]
		fmt.Println(friends[i])
	}
}
