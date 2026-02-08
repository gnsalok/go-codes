package main

import (
	"fmt"
)

func main() {

	//Anonymous Struct in Go
	person := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(person)
}
