package main

import(
	"fmt"
)


func main(){

	defer foo()
	bar()
	
}

func foo(){
	fmt.Println("Foo Called")
}


func bar(){
	fmt.Println("bar Called")
	
}