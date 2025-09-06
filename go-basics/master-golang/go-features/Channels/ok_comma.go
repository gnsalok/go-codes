package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 33
	}()

	v, ok := <-ch
	fmt.Println(v, ok)
}
