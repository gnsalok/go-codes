package main

import "fmt"

type Bot interface {
	greeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (*EnglishBot) greeting() string {
	return "Hi, Good Morning!!"
}

func (*SpanishBot) greeting() string {
	return "Hola!!"
}

func greet(b Bot) {
	fmt.Println(b.greeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	greet(&eb)
	greet(&sb)

}
