/*
- Struct is value type.
-

*/
package main

import (
	"fmt"
)

type ContactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// Embedded struct
		contact: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 233333,
		},
	}

	jim.updateName("jimmy")

	// calling the receiver function
	jim.print()
}

// if you remove the * from person then changes will not reflect in main.
// Struct is value type
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// print the struct
func (p person) print() {
	fmt.Printf("%+v", p)
}
