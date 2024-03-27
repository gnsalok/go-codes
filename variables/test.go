package main

import "fmt"

func main() {

	course := "CS101"

	fmt.Println(course)
	fmt.Println("Address of course", &course)

	var ptr *string = &course
	fmt.Println("Value at that address", *ptr, *(&course))

}
