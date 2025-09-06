package main

import (
	"fmt"
)
//return sum of array element 
func sum(arr *[3]int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}



func main() {

	// var arr [5]int
	array := [3]int{2, 3, 4} //slice in go
	sumOfArray := sum(&array)
	fmt.Printf("Sum of array = %d ", sumOfArray)

}
