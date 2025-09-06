package main

import (
	"fmt"
	"reflect"
)

type Bot interface {
	Speak() string
}

type EnglishBot struct{}
type SpanishBot struct{}

//Method Receiver
func (EnglishBot) Speak() string {
	return "Hello"
}

func (SpanishBot) Speak() string {
	return "Hola!"
}

// Function
func greet(b Bot) {
	fmt.Println(b.Speak())
}

func main() {

	var eb EnglishBot
	var sb SpanishBot

	greet(eb)
	greet(sb)

	fmt.Println(reflect.TypeOf(eb))
	fmt.Println(reflect.TypeOf(sb))

}
