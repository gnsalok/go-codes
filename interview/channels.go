package main

import "fmt"

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan string)

	str := [5]string{"a", "b", "c", "d", "e"}

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			// fmt.Print(i + 1)
			// <-ch2
		}
	}()

	for _, v := range str {
		fmt.Println(<-ch1)
		fmt.Println(v)
		// ch2 <- v
	}
}
