package main

import "fmt"

func produce(out chan<- string) {
	for _, t := range []string{"news1", "news2", "news3"} {
		out <- t
	}
	close(out)
}

func consume(in <-chan string) {
	for t := range in {
		fmt.Println(t)
	}
	// NOTE : you can't close a channel from the receiver side, it will panic. Only the sender can close the channel.
	// close(in) // This will panic if uncommented
	// out <-in // This will panic if uncommented, you can't send to a receive-only channel
}

func main() {
	ch := make(chan string)
	go produce(ch)
	consume(ch)
}
