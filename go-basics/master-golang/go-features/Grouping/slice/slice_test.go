package main

import (
	"fmt"
)

func printAnyValue(a ...interface{}) {
	fmt.Println(a)
}

func main() {

	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	sl2 := make([]int, 5, 10)

	copy(sl2, sl)

	fmt.Println(sl2[0])

	sl2 = append(sl2, 0, -1, 3, 5)

	fmt.Println(sl2)

}