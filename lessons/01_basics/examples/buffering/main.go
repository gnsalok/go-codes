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
	_, _ = writer.WriteString("Hello, ")
	_, _ = writer.WriteString("Go!")

	fmt.Println("Before Flush:", destination.String()) // Empty
	fmt.Println("Bytes waiting in buffer:", writer.Buffered())

	// Flush sends the buffered data to the destination.
	_ = writer.Flush()

	fmt.Println("After Flush:", destination.String()) // Hello, Go!
}
