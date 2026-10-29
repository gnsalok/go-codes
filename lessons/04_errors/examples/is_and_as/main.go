package main

import (
	"errors"
	"fmt"
	"os"
)

func readFile() error {
	return &os.PathError{Op: "open", Path: "config.yaml", Err: os.ErrNotExist}
}

func serviceStart() error {
	err := readFile()
	if err != nil {
		return fmt.Errorf("service init failed: %w", err)
	}
	return nil
}

func main() {
	err := serviceStart()

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("config file missing")
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Printf("operation: %s, file: %s\n", pathErr.Op, pathErr.Path)
	}
}
