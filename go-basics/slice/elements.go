package main

import "fmt"

func main() {
	lx := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//fmt.Println(lx[5:9])
	// fmt.Println(lx[5:])
	// fmt.Println(lx[:5])
	/*

		fmt.Println("Length", len(lx), "Capacity", cap(lx))

		lx = append(lx, 12)
		fmt.Println(lx)

		// it will double the size of the capacity just by adding one element to it.
		fmt.Println("Length", len(lx), "Capacity", cap(lx))
	*/

	// Appending slice into slice
	ly := []int{11, 12, 13, 14, 15}
	lx = append(lx, ly...)
	fmt.Println(lx)

}
