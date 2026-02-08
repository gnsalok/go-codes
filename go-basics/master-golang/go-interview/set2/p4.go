package main

import (
	"fmt"
)

func table(n int) {

	for i := 0; i < 10; i++ {
		fmt.Printf("%v * %v = %v\n", n, i+1, n*(i+1))
	}

}

func main() {
	var num int
	fmt.Println("Enter a num to find a table?")
	fmt.Scanf("%v", &num)

	table(num)

}
