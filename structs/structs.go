package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary float64
}

func main() {

	//Declare

	// var e1 Employee
	// or e1 := new(Employee) // reference to employee struct

	// Initialize
	// e1.id = 1
	// e1.name = "Alok"
	// e1.salary = 75000.000

	// OR

	//Declare and initialize in one go
	e1 := Employee{
		id:     1,
		name:   "Alok",
		salary: 100000,
	}

	fmt.Println(e1)
	fmt.Println("1st Employee is: ", e1.name)

}
