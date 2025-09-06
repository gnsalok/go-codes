package main

import "fmt"

func main() {

	// new func will return ptr, which value in non-nil
	var x = new(int)
	fmt.Println(x == nil) // prints false
	fmt.Println(*x)       // prints 0

}
