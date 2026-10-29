package main

import (
	"fmt"
	"sort"
)

func main() {
	// Map keys must be comparable. Iteration order is deliberately unspecified.
	cities := map[int]string{
		1: "Alok",
		2: "Rio",
		3: "Tokyo",
	}

	for id, city := range cities {
		fmt.Printf("%v - %v\n", id, city)
	}

	fmt.Println("\n----------")
	delete(cities, 2)

	city, found := cities[2]
	fmt.Println("Rio ", found, city)

	// Sort the keys when output order is part of the contract.
	keys := make([]int, 0, len(cities))
	for key := range cities {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%v - %v\n", key, cities[key])
	}
}
