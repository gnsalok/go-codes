package main

import (
	"fmt"
	"time"
)

/*
- epoc time is the best way to store the time as it store time as INT, it's good for range based query.
- In backend store everything as UTC and covert timezone at frontend for the end user.
*/

func main() {

	// utcTime := time.Now().UTC()
	// fmt.Println(utcTime)

	// // Define IST timezone (UTC+5:30)
	// istLocation, err := time.LoadLocation("Asia/Kolkata")
	// if err != nil {
	// 	fmt.Println("Error loading IST location:", err)
	// 	return
	// }

	// // Convert UTC to IST
	// istTime := utcTime.In(istLocation)
	// fmt.Println("IST Time: ", istTime)

	// println("-----------------------\n")

	epocTime := time.Now().Unix()
	fmt.Println(epocTime)

	utcTime := time.Unix(epocTime, 0).UTC()

	fmt.Println(utcTime)

	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading IST location:", err)
		return
	}

	// Convert UTC to IST
	istTime := utcTime.In(istLocation)
	fmt.Println("IST Time: ", istTime)
}
