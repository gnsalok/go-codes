package main

import (
	"fmt"
	"time"
)

func main() {

	// Getting epoch time
	now := time.Now()
	epoch := now.Unix()

	// Converting epoch time to human redable format
	epochNew := int64(epoch)    // Example epoch timestamp
	t := time.Unix(epochNew, 0) // 0 for nanoseconds
	layout := "2006-01-02 15:04:05"
	formattedTime := t.Format(layout)
	fmt.Println("Formatted time:", formattedTime)

	// differnce since
	diff := time.Since(t)
	fmt.Println("difference since epoch", diff)

}
