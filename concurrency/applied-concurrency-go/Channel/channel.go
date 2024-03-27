package main

import (
	"fmt"
	"time"
)

/*
ch <- data  : seding data into channel

data := <- ch : receiving data from a channel

*/

func main() {
	ch := make(chan string)

	// buffered channel, works in async way.
	// its not wait for receiver, but hold the value and take value from channel when it's need in program.
	// ch := make(chan string, 1)

	go greet(ch)

	time.Sleep(5 * time.Second)

	fmt.Println("Main Ready.")

	message := <-ch

	close(ch)

	fmt.Println("Message received.")

	fmt.Println("Message : ", message)

}

// ch chan<- string : send only channel
// ch <-chan string : receive only
func greet(ch chan<- string) {
	fmt.Println("Greeter is Ready.")
	ch <- "hello world"
	fmt.Println("Greeter completed.")
}
