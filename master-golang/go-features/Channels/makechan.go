package main 

import(
	"fmt"
)

func main(){

	c := make(chan int)
/*
if you try to do this without go routine it will block the channal cause sending and receiving happen at the same time.
	c <- 44
	For this perticular problem we can use buffered channel.
*/

	go func(){
	c <- 44
	}()

	fmt.Println(<-c)
}