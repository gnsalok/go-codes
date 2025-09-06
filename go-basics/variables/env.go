package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")
	fmt.Println("Logged User:", user, "\n")

	//Loop through all the env variable
	for i, env := range os.Environ() {
		fmt.Println(i, env)
	}
}
