package main

import "fmt"

func main() {

	f(5)

}

func f(i int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover at", r)
		}
	}()

	if i > 2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
}
