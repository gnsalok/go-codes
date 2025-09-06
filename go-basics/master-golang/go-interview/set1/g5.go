package main

import (
	"fmt"
)

//Creating struct in Go, it is way to create logical type in GO.
//This is composite data structure where we can store value of different type.
type Person struct {
	first string
	last  string
	age   int
}

func main() {

	//This is called creating value of type person.
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   21,
	}

	p2 := Person{
		first: "Tom",
		last:  "Cruise",
		age:   25,
	}

	// printing the value in different ways.
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
