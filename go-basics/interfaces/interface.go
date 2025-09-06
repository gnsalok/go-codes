package main

import "fmt"

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
func printMessage(b Bot) {
	fmt.Println(b.greeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printMessage(eb)
	printMessage(sb)

}
