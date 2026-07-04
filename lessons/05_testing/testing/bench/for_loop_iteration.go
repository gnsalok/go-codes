package main

import "testing"

/*
**command**
go test -bench=. -benchmem
*/
type Data struct {
	data [1201201]byte
}

func BenchmarkWithCopy(b *testing.B) {
	items := make([]Data, 1000)

	b.ResetTimer()

	for b.Loop() {

		for _, v := range items {
			_ = v
		}
	}
}

func BenchmarkWithoutCopy(b *testing.B) {
	items := make([]Data, 1000)

	b.ResetTimer()

	for b.Loop() {

		for i := range items {
			_ = items[i]
		}
	}
}
