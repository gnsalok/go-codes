package main

import "fmt"

type data struct {
	name string
}

func (d data) displayName() {
	fmt.Println("Name is : ", d.name)
}

func main() {
	d := data{name: "Alok"}

	// Creating a method value evaluates and stores a copy of the value receiver.
	// Later mutations to d therefore do not change what display sees.
	display := d.displayName
	display()

	d.name = "John"
	display()
}
