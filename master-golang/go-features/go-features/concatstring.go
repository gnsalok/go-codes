package main

import (
	"fmt"
	"strings"
)

func main() {

	var mystring strings.Builder

	myString.WriteString("Hello ")
	// here we append to the end of our string
	myString.WriteString("World")

	// print out our concatenated string
	fmt.Println(myString.String())

}
