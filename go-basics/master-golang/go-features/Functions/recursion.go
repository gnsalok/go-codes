package main

import "fmt"

func main() {

  v := fact(5)
  fmt.Println("Fact is : "v)

}

func fact(n int) int{
	
	if n == 0{
		return 1
	}

	return n * fact(n-1)
	
}