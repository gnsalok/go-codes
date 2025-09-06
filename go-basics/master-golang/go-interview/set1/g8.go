//Function and methods

/*
A method is a function that has a defined receiver.
*/

package main

import (
	"fmt"
)

type User struct {
	firstName, lastName string
}

// Receiver

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.firstName, u.lastName)
}

// Function

func Greeting() string {
	return "Returning User"
}

var i int

func main() {

	fmt.Println("Functoin VS Receiver")

	u := User{"Alok", "Tripathi"}
	user := u.Greeting()
	fmt.Println(user)
}
