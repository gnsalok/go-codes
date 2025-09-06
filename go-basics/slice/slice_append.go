package main

import (
	"fmt"
)

func main() {

	// initialize a slice literal
	newSlice := []string{"a", "b", "c", "d"}
	fmt.Println("The original slice is:", newSlice)

	// add an element to the slice
	newSlice = append(newSlice, "e")
	fmt.Println("The updated slice is:", newSlice)

	// add multiple values to the slice
	newSlice = append(newSlice, "f", "g", "h")
	fmt.Println("The final slice is:", newSlice)

}
