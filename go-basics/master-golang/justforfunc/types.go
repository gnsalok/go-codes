//types in GO

package main

import (
	"fmt"
)

type hotdog int

func main() {
	// var varibale_name type
	var a int
	var b hotdog

	a = 10
	b = 20

	// a = b  you can't do this
	//  a = int(b)   This will work - This is called Conversion in Go not Casting
	//	var var_name type

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)

}

