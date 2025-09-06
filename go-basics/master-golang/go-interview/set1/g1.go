package main

import "fmt"

func main() {

	m := make(map[string]int) //map[key]value

	m["answer"] = 100

	m["answer"] = 200

	delete(m, "answer")

	fmt.Println(m["answer"])

	v, ok := m["answer"]

	fmt.Println("The value :", v, "\nPresent? ", ok)

}
