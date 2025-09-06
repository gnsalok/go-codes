package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)

	m[1] = "Alok"
	m[2] = "Akhil"

	if val, ok := m[2]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Not exist")
	}
}
