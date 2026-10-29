package benchmark

import "testing"

type MyStruct struct {
	Value int
}

func (ms MyStruct) GetValue() int {
	return ms.Value
}

type MyInterface interface {
	GetValue() int
}

var benchmarkSink int

func BenchmarkInterface(b *testing.B) {
	data := make([]MyInterface, 10_000)
	for i := range data {
		data[i] = MyStruct{Value: i}
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for _, value := range data {
			sum += value.GetValue()
		}
		benchmarkSink = sum
	}
}

func BenchmarkConcrete(b *testing.B) {
	data := make([]MyStruct, 10_000)
	for i := range data {
		data[i] = MyStruct{Value: i}
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for _, value := range data {
			sum += value.Value
		}
		benchmarkSink = sum
	}
}
