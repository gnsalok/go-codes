// main.go
package main

import "fmt"

// DataService is an interface representing a data service.
type DataService interface {
	FetchData() string
}

// RealDataService is a real implementation of the DataService interface.
type RealDataService struct{}

// FetchData fetches data from the real data service.
func (rds RealDataService) FetchData() string {
	// Imagine some actual implementation fetching data from a database or an API.
	return "Real Data"
}

// Consumer is a struct that consumes data from a DataService.
type Consumer struct {
	DataService DataService
}

// UseData consumes data from the DataService.
func (c Consumer) UseData() string {
	return c.DataService.FetchData()
}

func main() {
	// Using the real data service
	realDataService := RealDataService{}
	consumer := Consumer{DataService: realDataService}
	fmt.Println(consumer.UseData())

	// Now, let's create a mock for testing
	mockDataService := MockDataService{"Mocked Data"}
	consumerWithMock := Consumer{DataService: mockDataService}
	fmt.Println(consumerWithMock.UseData())
}
