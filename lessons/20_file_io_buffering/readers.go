// Package fileiobuffering compares read strategies without owning the source's
// lifecycle. Callers decide whether the reader is a file, network connection, or
// in-memory value and remain responsible for closing it when necessary.
package fileiobuffering

import (
	"bufio"
	"errors"
	"io"
)

const chunkSize = 64 * 1024

type Stats struct {
	Bytes    int64
	Checksum uint64
}

// ReadOneByte requests one byte from the underlying reader on every call.
func ReadOneByte(reader io.Reader) (Stats, error) {
	return readWithBuffer(reader, 1)
}

// ReadBufferedByte requests one byte at a time from a buffered reader. The
// buffer turns most calls into memory access instead of calls to the source.
func ReadBufferedByte(reader io.Reader) (Stats, error) {
	buffered := bufio.NewReaderSize(reader, chunkSize)
	var stats Stats

	for {
		value, err := buffered.ReadByte()
		if err == nil {
			stats.Bytes++
			stats.Checksum += uint64(value)
			continue
		}
		if errors.Is(err, io.EOF) {
			return stats, nil
		}
		return stats, err
	}
}

// ReadChunks processes the source through a reusable 64 KiB buffer.
func ReadChunks(reader io.Reader) (Stats, error) {
	return readWithBuffer(reader, chunkSize)
}

func readWithBuffer(reader io.Reader, size int) (Stats, error) {
	buffer := make([]byte, size)
	var stats Stats

	for {
		read, err := reader.Read(buffer)
		for _, value := range buffer[:read] {
			stats.Checksum += uint64(value)
		}
		stats.Bytes += int64(read)

		if errors.Is(err, io.EOF) {
			return stats, nil
		}
		if err != nil {
			return stats, err
		}
	}
}
