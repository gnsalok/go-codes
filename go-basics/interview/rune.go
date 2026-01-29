package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
UNICODE & RUNE in GO :	https://golangbyexample.com/understanding-rune-in-golang/
When to use :
- rune is an alias for int32, and it represents a Unicode code point.
- You should use a rune when you intend to save Unicode Code Point in the rune value.
- A rune array should be used when all values in the array are meant to be a Unicode Code Point.
- Rune is also used to represent a character.
*/
func main() {
	r := 'a'

	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print Code Point
	fmt.Printf("Unicode CodePoint: %U\n", r)

	//Print Character
	fmt.Printf("Character: %c\n", r)
	s := "0b£"

	//This will print the Unicode Points
	fmt.Printf("%U\n", []rune(s))

	//This will the decimal value of Unicode Code Point
	fmt.Println([]rune(s))

}
