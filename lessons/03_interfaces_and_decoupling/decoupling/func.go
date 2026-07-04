package main

import "fmt"

type data struct {
	name string
}

// Value semantic, means it will create copy in stack
func (d data) displayName() {
	fmt.Println("Name is : ", d.name)
}

func main() {

	var d data

	d.name = "Alok"

	// f is pointing to the copy of data in stack
	f := d.displayName

	// call the method by variable
	// Indirection, it is pointer to copy of the data, not actual data
	f()

	// change the name
	d.name = "John"

	// does value change with f ? BIG NO... bcz we change copy of the data.
	// so where f is pointer remains unchanged.

	// But if we change the semantic, it will change the value.
	f()

}
