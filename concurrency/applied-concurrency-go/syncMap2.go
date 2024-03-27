package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("example", []int{1, 2, 3})
	fmt.Println(m.Load("example")) // [1 2 3] true
}
