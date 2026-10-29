package main

import (
	"fmt"
	"sync"
)

func Count() {
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}
}
func main() {

	var wg sync.WaitGroup

	wg.Go(func() {
		Count()
	})

	// OR

	// still old way to write go routine
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	Count()
	// }()

	wg.Wait()

}
