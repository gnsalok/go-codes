package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

// Car embedding Engine Interface
// Car can have its own implementation for Start and Stop

type Car struct {
	Engine
	Model string
}

func (c Car) Drive() {
	fmt.Println("Driving", c.Model)
}

func main() {

	c := Car{}
	c.Drive()

}
