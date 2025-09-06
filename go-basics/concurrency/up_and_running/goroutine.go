package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {

	wg.Add(2)

	start := time.Now()

	go doSomething()
	go doSomethingElse()

	elapsed := time.Since(start)

	wg.Wait() // Blocking code

	fmt.Printf("Total time took : %v ", elapsed)

}

func doSomething() {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("I have done something")
}

func doSomethingElse() {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("I have done something else...")
}
