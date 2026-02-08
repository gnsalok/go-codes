package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
	}
	ch := make(chan string)

	// Looping links slice
	for _, link := range links {
		go checklink(link, ch)

	}

	//looping channels
	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checklink(link, ch)
		}(l)
	}

}

// Check either link is up or down and put back in channel.
func checklink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might down!")
		c <- link
		return
	} else {
		fmt.Println(link, "is up")
		c <- link
	}

}
