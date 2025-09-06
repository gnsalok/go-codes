package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Success Function")
	}

	y := func() {
		fmt.Println("Failure Function")
	}

	doSomething(x, y)
}

func doSomething(success func(), failure func()) {
	worked := true
	if worked {
		success()
	} else {
		failure()
	}
}
