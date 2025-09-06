package main

import (
	"fmt"
	"time"
)

var (
	result = 0
	value  = 3
)

func main() {
	goChan := make(chan int)
	mainChan := make(chan string)

	go calculateSquare(value, goChan)
	go reportResult(goChan, mainChan)

	<-mainChan // block until it can read something from mainChan

}

func calculateSquare(value int, goChan chan int) {
	fmt.Println("Waiting for 3 seconds...")
	time.Sleep(3 * time.Second)
	result = value * value
	goChan <- result
}

func reportResult(goChan chan int, mainChan chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("The result of ", value, "squareed is ", <-goChan)
	mainChan <- "you can quit now.."

}
