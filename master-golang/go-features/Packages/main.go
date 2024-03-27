package main

import(
	"fmt"
	"github.com/gnsalok/gocodebase/practice-go/Packages/calculator"
)

func main(){
	x := []float64{1,2,3,4,5,6,7,8,9,}
	//unfirling the slice
	sum := calculator.Add(x...)
	fmt.Println("Sum of numbers : ",sum)

	
}