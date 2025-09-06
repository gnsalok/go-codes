package main

import "fmt"

func printInterface(foo, bar any) {
	fmt.Println(foo, bar)
}

func printAny[T any](foo, bar T) {
	fmt.Println(foo, bar)
}

func main() {

	printInterface(12, 13)
	printInterface(1, "Alok")

	printAny(12, 13)
	// printAny(1, "Alok") // gives error; Both element should be of same Type

	//a workaround
	printAny[any](1, "Alok")

	printAny[any](nil, nil)

}
