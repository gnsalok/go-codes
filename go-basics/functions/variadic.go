package main

import "fmt"

func main() {
	min := findMin(3, 1, 4, 5, 6)
	fmt.Println("Min value :", min)

}

func findMin(min ...int) int { // it passes slice of array
	m := min[0]
	for i, v := range min {
		fmt.Println(i, v)
		if v < m {
			m = v
		}
	}
	return m
}
