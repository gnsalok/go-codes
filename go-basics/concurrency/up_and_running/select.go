package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			ch1 <- "sending every second"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "sending every 2 seconds"
		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			ch3 <- "We are done!"
		}
	}()

	for {
		select { // Listening / monitor to the all the channel and in case of any message (based on channel case) it will log the message
		case msg := <-ch1:
			fmt.Println(msg)

		case msg := <-ch2:
			fmt.Println(msg + " something cool happend")

		case msg := <-ch3:
			fmt.Println(msg)
			os.Exit(0)
		}

	}

}
