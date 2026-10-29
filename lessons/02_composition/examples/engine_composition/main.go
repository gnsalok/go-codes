package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

// Car embeds the behavior it delegates to an Engine implementation.
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
