package main

import "fmt"

//Empty interfaces are used by code that handles values of unknown type.
// type i interface{}

var i interface{}

func main() {

	describe(i)

	i := 42
	describe(i)

	j := "hello"
	describe(j)
}

func describe(i interface{}) {
	// var a string
	fmt.Printf("(%v, %T)\n", i, i)

	// a = i.(string)
	// fmt.Printf("(%v, %T)\n", a, a)
}
