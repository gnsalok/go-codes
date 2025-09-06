package main

import "fmt"

func main() {

	age := 61

	if age >= 18 && age <= 30 {
		fmt.Println("above 18 ; below 30")
	} else if age > 30 && age <= 60 {
		fmt.Println("above 30 ; below 60")
	} else {
		fmt.Println("above 60")
	}

}
