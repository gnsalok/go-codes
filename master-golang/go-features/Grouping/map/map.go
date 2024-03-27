package main

import (
	"fmt"
)

func main() {
	//map declaration and def
	m := map[string]int{
		"James": 21,
		"Wrack": 22,
	}
	fmt.Println(m)

	//This is , ok idiom
	if v, ok := m["James"]; ok {
		fmt.Println("This is the if print", v)
	}

}
