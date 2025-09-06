package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Package level var sdeclaration
var (
	// name, course string
	// module, clip int

	// Assignment
	name, course = "Alok", "CS101"
	module       = "4"
	clip         = 4
)

func main() {
	isTrue := true
	if isTrue {
		fmt.Println("Yes it is True")
	}

	// Short declaration
	test := "test"
	fmt.Println(test)

	fmt.Println("Name and Course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")

	fmt.Print("\n")

	// Runtime reflection to check type
	fmt.Println("Name and Course are set to", reflect.TypeOf(name), "and", reflect.TypeOf(course), ".")
	fmt.Println("Module and clip are set to", reflect.TypeOf(module), "and", reflect.TypeOf(clip), ".")

	iModule, err := strconv.Atoi(module)

	// if ther is any error
	if err != nil {
		fmt.Println("Error :", err)
	}
	total := clip + iModule
	fmt.Println("Total is ", total)

	fmt.Print("\n")

	// Go is Pass By Value ; Pointers variable
	fmt.Println("Course is", &course)

	var ptr *string = &course // To store ref we need ptr variable

	fmt.Println("Pointer to course", ptr)  // will give memory address ; ptr will hold reference
	fmt.Println("which holds value", *ptr) // will give value to this memory address ; dereference

}
