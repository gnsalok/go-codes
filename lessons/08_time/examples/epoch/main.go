package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	epochSeconds := now.Unix()
	instant := time.Unix(epochSeconds, 0)

	layout := "2006-01-02 15:04:05"
	formattedTime := instant.Format(layout)
	fmt.Println("Formatted time:", formattedTime)

	diff := time.Since(instant)
	fmt.Println("difference since epoch", diff)
}
