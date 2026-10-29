package main

import "fmt"

type Bot interface {
	greeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (EnglishBot) greeting() string {
	return "Hi there!"
}

func (SpanishBot) greeting() string {
	return "Hola!!"
}

func printMessage(bot Bot) {
	fmt.Println(bot.greeting())
}

func main() {
	english := EnglishBot{}
	spanish := SpanishBot{}

	var bot Bot
	fmt.Println(bot) // The zero value of an interface is nil.

	printMessage(english)
	printMessage(spanish)
}
