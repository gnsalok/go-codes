package main 

import(
	_ "fmt"
)

func main(){

	// Normal channel declaration 
	// ch := make(chan int)


	//making channel direction that can store values in the channel
	//without goroutine you need to run as buffered 


	//this send only type channel,you can't extrac the value from it.
	ch := make(chan <- int)

	// receive channel, you can't send value to the channel.
	// ch := make(<- chan int)

	// goroutine because sending and receiving happend so fast that's why it blocks the channel
	go func(){
	ch <- 155
	}()

	// fmt.Println(<-ch)

	


}