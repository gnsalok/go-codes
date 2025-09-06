package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter you choice to select day of the week?")
	var ch int
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Print("Monday")
		break
	case 2:
		fmt.Print("Tuesday")
		break
	case 3:
		fmt.Print("Wednesday")
		break
	case 4:
		fmt.Print("Thursday")
		break
	case 5:
		fmt.Print("Friday")
		break
	case 6:
		fmt.Print("Saturday")
		break
	case 7:
		fmt.Print("Sunday")
		break
	default:
		fmt.Print("Pls enter valid choice")
	break
	}

}
