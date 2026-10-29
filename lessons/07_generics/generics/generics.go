package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Add[T Number](x, y T) T {
	return x + y
}

// Map applies function f to each element of s and returns a new slice.
func Map[T, U any](values []T, transform func(T) U) []U {
	mapped := make([]U, 0, len(values))
	for _, value := range values {
		mapped = append(mapped, transform(value))
	}
	return mapped
}

func main() {
	fmt.Println(Add[int](2, 3))         // Output: 5
	fmt.Println(Add[float64](2.5, 3.7)) // Output: 6.2

	nums := []int{1, 2, 3}
	squares := Map(nums, func(v int) int { return v * v })
	fmt.Println(squares) // Output: [1 4 9]

	floats := []float64{1.5, 2.5, 3.5}
	strs := Map(floats, func(v float64) string { return fmt.Sprintf("%.1f", v) })
	fmt.Println(strs) // Output: [1.5 2.5 3.5]
}
