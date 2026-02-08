package main

import "fmt"
import "errors"

// divide two number
func Divide(a int, b int) (int, error) {

	// can not divide by `0`
	if b == 0 {
		// throwing error from new method which belongs to package error
		return 0, errors.New("Devide by Zero Occured!")
	} else {
		return (a / b), nil
	}
}

func main() {

	// divide 4 by 0
	result, err := Divide(4, 0)

	// if error occures
	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("4/0 is", result)
	}
}
