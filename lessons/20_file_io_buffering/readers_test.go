package fileiobuffering

import (
	"bytes"
	"testing"
)

func TestReadStrategiesProduceSameStats(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 255}
	want := Stats{Bytes: int64(len(data)), Checksum: 270}

	strategies := map[string]func(*bytes.Reader) (Stats, error){
		"one byte": func(reader *bytes.Reader) (Stats, error) {
			return ReadOneByte(reader)
		},
		"buffered byte": func(reader *bytes.Reader) (Stats, error) {
			return ReadBufferedByte(reader)
		},
		"chunks": func(reader *bytes.Reader) (Stats, error) {
			return ReadChunks(reader)
		},
	}

	for name, read := range strategies {
		t.Run(name, func(t *testing.T) {
			got, err := read(bytes.NewReader(data))
			if err != nil {
				t.Fatalf("read: %v", err)
			}
			if got != want {
				t.Fatalf("stats = %+v, want %+v", got, want)
			}
		})
	}
}
