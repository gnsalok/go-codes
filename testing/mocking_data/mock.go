// mocks_test.go
package main

// MockDataService is a mock implementation of the DataService interface for testing.
type MockDataService struct {
	MockedData string
}

// FetchData fetches mocked data from the mock data service.
func (mds MockDataService) FetchData() string {
	return mds.MockedData
}
