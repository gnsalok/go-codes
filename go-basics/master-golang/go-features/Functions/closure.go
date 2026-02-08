package main

import "fmt"

func main() {

	a := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := incrementor()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())


}


//writing closure
func incrementor() func() int{
	var x int
	return func() int {
		//this function can access the var x which defined outside the function.
		x++
		return x
	}
}