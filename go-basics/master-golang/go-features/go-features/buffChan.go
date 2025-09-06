package main

import (
	"fmt"
)

//channels are used to communicate with goroutines. Goroutine should be the to work with channels.
func sendValue(c chan int) {
	//sending value to the channel
	c <- 8

	fmt.Println("Execution Finished!")
}

func main() {

	fmt.Println("Go Channels!")
	//Buffered chan will execute crossed go routine.
	values := make(chan int, 2)

	//closing the channel
	defer close(values)

	go sendValue(values)
	go sendValue(values)
	//receiving value from the channel
	value := <-values

	fmt.Println(value)

}
