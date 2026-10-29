package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	fileio "github.com/gnsalok/go-codes/lessons/20_file_io_buffering"
)

const sampleSize = 8 * 1024 * 1024

type zeroReader struct{}

func (zeroReader) Read(buffer []byte) (int, error) {
	clear(buffer)
	return len(buffer), nil
}

type result struct {
	name     string
	duration time.Duration
	stats    fileio.Stats
}

func main() {
	path, err := createSampleFile(sampleSize)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create sample file:", err)
		os.Exit(1)
	}
	defer os.Remove(path)

	strategies := []struct {
		name string
		read func(io.Reader) (fileio.Stats, error)
	}{
		{name: "one-byte reads", read: fileio.ReadOneByte},
		{name: "buffered byte reads", read: fileio.ReadBufferedByte},
		{name: "64 KiB chunk reads", read: fileio.ReadChunks},
	}

	for _, strategy := range strategies {
		measured, err := measure(path, strategy.name, strategy.read)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", strategy.name, err)
			continue
		}
		fmt.Printf("%-20s duration=%-10s bytes=%d checksum=%d\n",
			measured.name, measured.duration, measured.stats.Bytes, measured.stats.Checksum)
	}
}

func createSampleFile(size int64) (string, error) {
	file, err := os.CreateTemp("", "go-file-io-*.bin")
	if err != nil {
		return "", err
	}
	path := file.Name()

	_, writeErr := io.CopyN(file, zeroReader{}, size)
	closeErr := file.Close()
	if err := errors.Join(writeErr, closeErr); err != nil {
		_ = os.Remove(path)
		return "", err
	}
	return path, nil
}

func measure(path, name string, read func(io.Reader) (fileio.Stats, error)) (result, error) {
	file, err := os.Open(path)
	if err != nil {
		return result{}, err
	}

	start := time.Now()
	stats, readErr := read(file)
	duration := time.Since(start)
	closeErr := file.Close()

	return result{name: name, duration: duration, stats: stats}, errors.Join(readErr, closeErr)
}
