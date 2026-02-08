package main

import (
	"fmt"
	"sync"
)

var (
	wg              sync.WaitGroup
	widgetInventory int32 = 1000
)

func main() {

	fmt.Println("Starting inventory count = ", widgetInventory)
	wg.Add(2)

	go makeSales()
	go newPurchase()

	wg.Wait()

	fmt.Println("Ending inventory count is : ", widgetInventory)
}

func makeSales() { // 1000000 widget sold
	defer wg.Done()
	// making sale 3000 times
	for i := 0; i < 3000; i++ {
		widgetInventory -= 100
	}
}

func newPurchase() { // 1000000 new purchased
	defer wg.Done()
	//making purchase 3000 times
	// ending inventory should be the same
	for i := 0; i < 3000; i++ {
		widgetInventory += 100
	}

}
