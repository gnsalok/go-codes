package main

import (
	"fmt"
)

func main() {

	demoPanic()

}

func demoPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")
}
