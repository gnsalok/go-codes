package main

import (
	"fmt"
	"time"
)

func main() {
	epochSeconds := time.Now().Unix()
	fmt.Println(epochSeconds)

	utcTime := time.Unix(epochSeconds, 0).UTC()
	fmt.Println(utcTime)

	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading IST location:", err)
		return
	}

	istTime := utcTime.In(istLocation)
	fmt.Println("IST Time: ", istTime)
}
