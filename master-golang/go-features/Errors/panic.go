package main

import (
	_ "fmt"
	"os"
)

func main() {

	_, err := os.Open("log.txt")

	if err != nil {
		panic("Some error happend!!")
	}

}
