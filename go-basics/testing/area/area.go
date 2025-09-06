package area

import "math"

/*
link : https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces
*/

// Interface
type Shape interface {
	Area() float64
}

// Types
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Methods
func (rectangle Rectangle) Area() float64 {
	area := (rectangle.Width * rectangle.Height)
	return area

}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
