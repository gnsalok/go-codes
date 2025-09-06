package benchamarking

import (
	"testing"
)

type MyStruct struct {
	Value int
}

func (ms MyStruct) GetValue() int {
	return ms.Value
}

type MyInterface interface {
	GetValue() int
}

func BenchmarkInterface(b *testing.B) {
	var data []MyInterface
	for i := 0; i < 1000000; i++ {
		data = append(data, MyStruct{Value: i})
	}
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			_ = v.GetValue()
		}
	}
}

func BenchmarkConcrete(b *testing.B) {
	var data []MyStruct
	for i := 0; i < 1000000; i++ {
		data = append(data, MyStruct{Value: i})
	}
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			_ = v.Value
		}
	}
}
