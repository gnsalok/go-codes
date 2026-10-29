package main

import "fmt"

type example struct {
	x float64
	f bool
}

func main() {
	// Every field starts at its type's zero value.
	var e1 example
	fmt.Printf("%+v\n", e1) // {x:0 f: false}

	e2 := example{
		x: 1.234,
		f: true,
	}
	fmt.Println(e2.f)
	fmt.Println(e2.x)
}
