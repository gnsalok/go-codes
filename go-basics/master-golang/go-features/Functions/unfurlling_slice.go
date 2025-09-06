package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 6, 5, 7, 8, 9}

	//unfurling the slice
	sum := sum(xi...)
	fmt.Println("The total is : ", sum)

}

func sum(xi ...int) int {
	fmt.Printf("Type of xi %T\n", xi)
	total := 0

	for _, v := range xi {
		fmt.Println(v)
		total += v
	}
	return total

}
