package main

import "fmt"

func main() {

	lx := []int{1, 2, 3, 4, 5}
	reveseArray := reverseList(lx)
	fmt.Println(reveseArray)

}

func reverseList(lx []int) []int {

	reverseSlice := make([]int, len(lx))
	size := len(lx) - 1
	count := 0

	for i := size; i >= 0; i-- {
		reverseSlice[count] = lx[i]
		count += 1
	}
	return reverseSlice
}
