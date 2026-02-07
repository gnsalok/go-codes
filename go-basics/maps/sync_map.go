package main

import (
	"fmt"
	"sync"
)

func main() {

	// sync.Map is concurrent safe map
	// sync.Map does not have any type parameters, so it can store any type of key and value.
	// sync.Map does not have a built-in way to specify the initial size of the map, so it may be less efficient than a regular map if you know the number of elements in advance.

	/*
		Best for: Scenarios where keys are mostly read and rarely updated,
		or when multiple goroutines read/write for disjoint sets of keys.
	*/

	var emp sync.Map

	emp.Store(1, "Tokyo")
	emp.Store(2, "Professor")
	emp.Store(3, "Rio")

	//updating map
	emp.Store(3, "Alok")

	//delete element in map
	emp.Delete(1)

	emp.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
