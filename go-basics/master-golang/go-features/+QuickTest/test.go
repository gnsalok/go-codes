package main

import "fmt"

func foo(i string) (string, string) {
	fmt.Println(i)
	return "Hi", "There"

}

func main() {
	// fmt.Println("Hi there")
	// fmt.Println(foo("Hi"))

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

}
