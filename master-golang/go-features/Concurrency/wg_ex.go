package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		showNum()
		wg.Done()
	}()
	showLett()
}

func showNum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func showLett() {
	for i := 65; i <= 90; i++ {
		fmt.Println(string(i))
	}
}
