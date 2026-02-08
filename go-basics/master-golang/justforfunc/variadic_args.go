package main

import(
	"fmt"
)

//Variadic function args with return
func printAnyValue(a ...interface{}){
	fmt.Println(a)
}

func main(){

	printAnyValue(12,"Alok", 13, 555)
}