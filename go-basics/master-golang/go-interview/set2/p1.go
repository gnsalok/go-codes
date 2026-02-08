package main

import (
	"fmt"
)

func main() {

	//define array of size 5
	var arr [5]int
	fmt.Println(arr)

	arr[4] = 100
	fmt.Println(arr)
	fmt.Println("get:", arr[4])

	fmt.Println("lenght of arr :", len(arr))

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D ", twoD)

}
