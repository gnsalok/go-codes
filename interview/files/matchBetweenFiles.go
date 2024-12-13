/*
Problems:

There two files of distint days. If the same user made request for same input,
what is the best way to compare from two files that user is the same?

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file1 := "file1.txt" // Path to the first file
	file2 := "file2.txt" // Path to the second file

	// Read and index the first file
	userInputMap := make(map[string]struct{})
	readFile(file1, func(user, input string) {
		key := user + "|" + input // Create a unique key for user-input
		userInputMap[key] = struct{}{}
	})

	// Compare with the second file
	fmt.Println("Matching User-Input Combinations:")
	readFile(file2, func(user, input string) {
		key := user + "|" + input
		if _, exists := userInputMap[key]; exists {
			fmt.Println("Match found:", user, input)
		}
	})
}

// Helper function to read a file line-by-line and process user-input combinations
func readFile(filename string, processLine func(user, input string)) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",") // Assuming CSV format: user,input
		if len(parts) == 2 {
			user := strings.TrimSpace(parts[0])
			input := strings.TrimSpace(parts[1])
			processLine(user, input)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
