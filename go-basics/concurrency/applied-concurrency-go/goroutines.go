package main

import (
	"fmt"
	"time"
)

// Never use sleep() in production, its not guarantee

func main() {
	go hello()
	time.Sleep(1000 * time.Millisecond)
	goodbye()
}

func hello() {
	fmt.Println("Hello")
}

func goodbye() {
	fmt.Println("Goodbye")
}
