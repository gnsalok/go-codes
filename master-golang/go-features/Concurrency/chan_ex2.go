package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go count("employee", ch)

	for msg := range ch {

		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
