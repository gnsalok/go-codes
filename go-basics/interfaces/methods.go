package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
}

// Receiver
func (e Employee) greeting() string {
	return fmt.Sprintf("Dear %s %s", e.firstName, e.lastName)

}

func main() {

	e1 := Employee{
		firstName: "Alok",
		lastName:  "Tripathi",
	}

	fmt.Println(e1.greeting())

}
