package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Alok"
	course := "cs101"

	fmt.Println(converter(name, course))

}

func converter(name, course string) (n, c string) {
	name = strings.ToUpper(name)
	course = strings.Title(course)
	return name, course
}
