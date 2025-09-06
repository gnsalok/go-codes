package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("test.txt")

	if err != nil {
		fmt.Print("Error is :", err)
	}

}
