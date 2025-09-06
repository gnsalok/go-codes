package main

import "fmt"

type example struct {
	x float64
	f bool // by default, it will be false
}

func main() {

	var e1 example
	fmt.Printf("%+v\n", e1) // {x:0 f: false}

	e2 := example{
		1.234,
		true,
	}
	fmt.Printf("%+v\n", e2.f)
	fmt.Printf("%+v\n", e2.x)
}
