package main

import (
	"fmt"
)

//Function to return sum of float array value.
//dealing with call by Ref
func sumArr(arr *[5]float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func main() {

	arr := [5]float64{1.33, 2, 3.56, 4.5, 5}
	result := sumArr(&arr)
	fmt.Println(result)
}
