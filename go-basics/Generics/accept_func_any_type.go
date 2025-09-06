package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

// helper method to turn any type into ptr type
func makePointer[T any](t T) *T {
	return &t

}

func main() {

	p := person{
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"), // This line won't compile, without using func makePointer
		LastName:   "Peterson",
	}

	fmt.Println(p)
}
