/*

Slice is reference type.

*/

package main

import (
	"fmt"
)

func main() {

	// slc := []datatype{ values }  -- composite literal

	slice := []int{2, 3, 4, 5, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	//slicing a slice
	fmt.Println(slice[1:5])

	//appending a slice

	slice = append(slice, 11, 12, 13)
	fmt.Println(slice)

	//appending slice to a slice
	//append(where, what)

	appSlice := []int{15, 16, 17}
	slice = append(slice, appSlice...)
	fmt.Println(slice)

	//Deleting from the slice
	slice = append(slice[:2], slice[4:]...)
	fmt.Println(slice)

	//creating slice using make

	x := make([]int, 10, 100)
	x[0] = 12
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	/*
	   //Index out of range error
	   	x[10]=112
	   	fmt.Println(x)
	   	fmt.Println(len(x))
	   	fmt.Println(cap(x))
	*/

	x = append(x, 112)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
