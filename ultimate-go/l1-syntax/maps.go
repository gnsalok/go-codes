package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	 * One Important thing  - Try to keep Key which can be hashed, and you check using IF statement, Otherwise
	 * It will become a key violation
	 */

	m := map[int]string{
		1: "Alok",
		2: "Rio",
		3: "Tokyo",
	}

	// Note that order never guarantee, it's random
	// Here we are working on copy of the value of the map, not actual value
	for k, v := range m {
		fmt.Printf("%v - %v\n", k, v)
	}

	fmt.Println("\n----------")

	delete(m, 2)

	// Note that, if you deleted some key from the map,  without checking `found`, you will always get Zero value
	v, found := m[2]

	fmt.Println("Rio ", found, v)

	// Sort they key - so that we can predict the order
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%v - %v\n", key, m[key])
	}

}
