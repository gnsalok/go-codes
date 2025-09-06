package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(5)

	for i := 0; i < 5; i++ {

		go func(j int) {
			m.Store(j, fmt.Sprintf("test %v", j))
			wg.Done()
		}(i)

	}

	wg.Wait()
	fmt.Println("Done.")

	for i := 0; i < 5; i++ {
		t, _ := m.Load(i)
		fmt.Println("for loop: ", t)
	}

	m.Range(func(k, v interface{}) bool {
		fmt.Println("range (): ", v)
		return true
	})
}
