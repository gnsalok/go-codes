package main

import (
	"fmt"
	"time"
)

var (
	ch = make(chan string)
)

func main() {

	start := time.Now()

	go doSomething()
	go doSomethingElse()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	elapsed := time.Since(start)
	fmt.Printf("Total time took : %v ", elapsed)

}

func doSomething() {
	time.Sleep(2 * time.Second)
	ch <- "do something finished...."
}

func doSomethingElse() {
	time.Sleep(2 * time.Second)
	ch <- "do something else is finished"
}
