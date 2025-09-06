package main

import "fmt"

type Number interface {
	int | float64
}

func Add[T Number](x, y T) T {
	return x + y
}

func main() {
	fmt.Println(Add[int](2, 3))         // Output: 5
	fmt.Println(Add[float64](2.5, 3.7)) // Output: 6.2
}
