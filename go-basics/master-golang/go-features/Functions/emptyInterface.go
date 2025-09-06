package main

import "fmt"

//Empty Interface
func show(a ...interface{}) {
	fmt.Printf("Value : %v\tType : %T\n", a, a)

}

func main() {

	show("Alok",21)
	show(21)

}
