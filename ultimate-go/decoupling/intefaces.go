package main

import "fmt"

/*
 * Interface value are valueless.
 * When you attach any concrete type it will behave according to the type it passed.
 */

type Bot interface {
	greeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

// EnglishBot satisfy Bot interface
func (EnglishBot) greeting() string {
	return "Hi there!"
}

// SpanishBot satisfy Bot interface
func (SpanishBot) greeting() string {
	return "Hola!!"
}

// Any type which satisfy Bot interface can leverage this function
// Polymorphic function
func printMessage(b Bot) {
	fmt.Println(b.greeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	var b Bot
	// This is <nil>
	fmt.Println(b)

	printMessage(eb)
	printMessage(sb)

}
