// Constructor
package main

// Lets say its order package then Name it like order.New(100)

// Right below your type, write constructor
type Order struct {
	Size float64
}

func NewOrder(size float64) *Order {
	return &Order{
		Size: size,
	}
}

func main() {

}
