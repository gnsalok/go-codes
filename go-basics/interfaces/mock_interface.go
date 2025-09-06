package main

import "fmt"

type DataFetcher interface {
	FetchData() string
}

type MockDataFetcher struct{}

func (m MockDataFetcher) FetchData() string {
	return "Mock Data"
}

func main() {
	m := MockDataFetcher{}

	data := m.FetchData()

	fmt.Println(data)
}
