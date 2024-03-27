package main

import(
	"fmt"
)


func calculate(x int)(result int){
	result = x + 2
	return result
} 

func main(){
	fmt.Println("Go Tesing...")
	result := calculate(2)
	fmt.Println(result)
}