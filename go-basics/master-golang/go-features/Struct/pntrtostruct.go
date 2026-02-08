package main

import (
	"fmt"
)

// ContactInfo struct
type ContactInfo struct {
	Email   string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

func main() {

	p1 := Person{
		FirstName: "James",
		LastName:  "Bond",
		ContactInfo: ContactInfo{
			Email:   "alok123@gmail.com",
			ZipCode: 273402,
		},
	}
	fmt.Println(p1)

}
