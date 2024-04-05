package main

import "fmt"

type notifierInt interface {
	notify()
}
type duration int

func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func main() {

	/*
		Why you can't call notify method on type int? Bcz you can’t always take the address of a value in Go.
		* First, you can’t guarantee that every value is addressable. If you can’t take a value’s address, it can’t be shared and therefore a pointer receiver method can’t be used.
		* This next example shows how you can’t always take the address of a value in Go.
	*/
	duration(42).notify()
}
