package main

import (
	"fmt"
)

type Calculator struct {
	num1 int
	num2 int
}

//Extending the method of Calculator struct 
func (c *Calculator) add(n1 int, n2 int) int {
	c.num1 = n1
	c.num2 = n2
	return (c.num1 + c.num2)
}

func (c *Calculator) sub(n1 int, n2 int) int {
	c.num1 = n1
	c.num2 = n2
	return (c.num1 - c.num2)
}

func (c *Calculator) mul(n1 int, n2 int) int {
	c.num1 = n1
	c.num2 = n2
	return (c.num1 * c.num2)
}

func (c *Calculator) div(n1 int, n2 int) int {
	c.num1 = n1
	c.num2 = n2
	return (c.num1 / c.num2)
}

func main() {

	var C Calculator
	var result int
	var o1 int
	var o2 int

	fmt.Printf("Please enter you choice?\n")
	var ch int

	fmt.Printf("1.Add\n2.Sub\n3.Mul\n4.Div\n")
	fmt.Scanf("%d", &ch)

	switch ch {
	case 1:
		fmt.Printf("Enter you two operand to perform the operation?")
		fmt.Scanf("%d %d", &o1, &o2)
		result = C.add(o1, o2)
		fmt.Printf("The result is :%v", result)
		break


		/*
			add more methods to perform calculation....
		*/
	default:
		fmt.Println("Incorrect option!!")
	}

}
