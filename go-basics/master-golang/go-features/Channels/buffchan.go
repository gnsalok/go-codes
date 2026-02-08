package main

import (
	"fmt"
)

func main() {

	//making buffered channal
	c := make(chan int, 1)
	c <- 44
	fmt.Println(<-c)
}
