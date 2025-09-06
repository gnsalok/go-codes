package main

import "fmt"

func main() {
	ch := 5

	// Go have implicit break

	switch ch {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wed")
	case 4:
		fmt.Println("Thu")
		break // it will break the checks and exit the program
	case 5:
		fmt.Println("Friday")
		fallthrough // it will execute next case anyway but its not recommended to use
	case 6:
		fmt.Println("Saturday")
		fallthrough
	default:
		fmt.Println("Invalid Choice")
	}

}
