package main

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	testFileName = "large_test_file.bin"
	testFileSize = 200 * 1024 * 1024 // 200 MB
)

type result struct {
	name     string
	duration time.Duration
	bytes    int64
	checksum uint64
}

func main() {
	if err := ensureTestFile(testFileName, testFileSize); err != nil {
		fmt.Printf("failed to create test file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Testing file: %s\n", testFileName)
	fmt.Printf("File size: %.2f MB\n\n", float64(testFileSize)/(1024*1024))

	tests := []struct {
		name string
		fn   func(string) (int64, uint64, error)
	}{
		{
			name: "1. Unbuffered: 1 byte per Read",
			fn:   processOneByteAtATime,
		},
		{
			name: "2. bufio.Reader: 64 KB buffer",
			fn:   processWithBufio,
		},
		{
			name: "3. Chunked Read: 64 KB slice",
			fn:   processWithChunk,
		},
	}

	var results []result

	for _, test := range tests {
		start := time.Now()

		bytesProcessed, checksum, err := test.fn(testFileName)
		if err != nil {
			fmt.Printf("%s failed: %v\n", test.name, err)
			continue
		}

		results = append(results, result{
			name:     test.name,
			duration: time.Since(start),
			bytes:    bytesProcessed,
			checksum: checksum,
		})
	}

	printResults(results)
}

func ensureTestFile(filename string, expectedSize int64) error {
	info, err := os.Stat(filename)

	if err == nil && info.Size() == expectedSize {
		fmt.Println("Using existing test file.")
		return nil
	}

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	fmt.Printf("Creating %.2f MB test file...\n",
		float64(expectedSize)/(1024*1024))

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 1024*1024)
	buffer := make([]byte, 1024*1024)

	var written int64

	for written < expectedSize {
		remaining := expectedSize - written
		writeSize := int64(len(buffer))

		if remaining < writeSize {
			writeSize = remaining
		}

		if _, err := rand.Read(buffer[:writeSize]); err != nil {
			return err
		}

		n, err := writer.Write(buffer[:writeSize])
		if err != nil {
			return err
		}

		written += int64(n)
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Println("Test file created.")
	return nil
}

// Method 1:
//
// This asks the operating system for only one byte on every Read call.
// For a 200 MB file, it can result in roughly 200 million Read calls.
func processOneByteAtATime(filename string) (int64, uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	buffer := make([]byte, 1)

	var (
		total    int64
		checksum uint64
	)

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			checksum += uint64(buffer[0])
			total += int64(n)
		}

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return total, checksum, err
		}
	}

	return total, checksum, nil
}

// Method 2:
//
// bufio.Reader internally reads a larger block from the file.
// The application still requests one byte at a time, but most requests
// are served from memory instead of calling the operating system.
func processWithBufio(filename string) (int64, uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 64*1024)

	var (
		total    int64
		checksum uint64
	)

	for {
		value, err := reader.ReadByte()

		if err == nil {
			checksum += uint64(value)
			total++
			continue
		}

		if errors.Is(err, io.EOF) {
			break
		}

		return total, checksum, err
	}

	return total, checksum, nil
}

// Method 3:
//
// The application directly requests 64 KB chunks.
// This avoids one-byte-at-a-time function call overhead and reduces
// operating-system read calls.
func processWithChunk(filename string) (int64, uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	buffer := make([]byte, 64*1024)

	var (
		total    int64
		checksum uint64
	)

	for {
		n, err := file.Read(buffer)

		for _, value := range buffer[:n] {
			checksum += uint64(value)
		}

		total += int64(n)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return total, checksum, err
		}
	}

	return total, checksum, nil
}

func printResults(results []result) {
	fmt.Println("\nResults")
	fmt.Println("------------------------------------------------------------")

	var fastest time.Duration

	for _, r := range results {
		throughput := float64(r.bytes) /
			(1024 * 1024) /
			r.duration.Seconds()

		fmt.Printf("%s\n", r.name)
		fmt.Printf("  Duration   : %v\n", r.duration)
		fmt.Printf("  Throughput : %.2f MB/s\n", throughput)
		fmt.Printf("  Checksum   : %d\n\n", r.checksum)

		if fastest == 0 || r.duration < fastest {
			fastest = r.duration
		}
	}

	fmt.Println("Relative performance")
	fmt.Println("------------------------------------------------------------")

	for _, r := range results {
		speedDifference := float64(r.duration) / float64(fastest)

		fmt.Printf("%-38s %.2fx\n", r.name, speedDifference)
	}
}
