package main

import "fmt"
//a closure is a technique for implementing a lexically scoped name binding in a language--
// with first-class functions - Wikipedia


//In layman’s terms, a closure is a function value which is able to reference variables that lay outwith it’s body.

// Note - It’s important to note the distinct differences between both closures and anonymous functions which are 
// --commonly mistaken for closures. You can learn more about anonymous functions here: Go Anonymous Functions


func getLimit() func() int {
    limit := 10
    return func() int {
        limit -= 1
        return limit
    }
}

func main() {
    limit := getLimit()
    fmt.Println(limit()) // 9
    fmt.Println(limit()) // 8

    limit2 := getLimit()
    fmt.Println(limit2()) // 9
    fmt.Println(limit2()) // 8

    fmt.Println(limit()) // 7

}