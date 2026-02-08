package main

import "fmt"

// Constant Declaration
// variable grouping in to function

// const pi = 3.14
// or
const (
	pi = 3.14
	g  = 9.18
)

func main() {
	fmt.Println(pi, g)

	TestFunc()
}

func TestFunc() {
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println(a, b, c)

}
