package main

import (
	"fmt"
)

// Generic function to process any type
func sortItems[T any](items []T, sortFunc func(T)) {
	for _, item := range items {
		sortFunc(item)
	}
}

// Example usage
func main() {
	// Example with integers
	intArray := []int{1, 2, 3, 4, 5}
	sortItems(intArray, func(item int) {
		fmt.Println("Processing int:", item)
	})

	// Example with custom type
	type Node struct {
		Value int
		Next  *Node
	}
	nodes := []*Node{
		{Value: 10},
		{Value: 20},
		{Value: 30},
	}
	sortItems(nodes, func(item *Node) {
		fmt.Println("Processing Node:", item.Value)
	})
}
