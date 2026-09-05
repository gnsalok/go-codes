package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// destination represents a file or network connection.
	var destination bytes.Buffer

	// The buffered writer temporarily holds small writes in memory.
	writer := bufio.NewWriter(&destination)
	writer.WriteString("Hello, ")
	writer.WriteString("Go!")

	fmt.Println("Before Flush:", destination.String()) // Empty
	fmt.Println("Bytes waiting in buffer:", writer.Buffered())

	// Flush sends the buffered data to the destination.
	writer.Flush()

	fmt.Println("After Flush:", destination.String()) // Hello, Go!
}
