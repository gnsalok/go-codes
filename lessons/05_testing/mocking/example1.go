// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how you can personally mock concrete types when
// you need to for your own packages or tests.
package main

import "github.com/gnsalok/go-codes/lessons/05_testing/mocking/pubsub"

// publisher is an interface to allow this package to mock the pubsub
// package support.
type publisher interface {
	Publish(key string, value any) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, value any) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {

	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	return nil
}

func main() {
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}

	// The consumer owns the interface; pubsub does not need to declare it.
	for _, publisher := range pubs {
		_ = publisher.Publish("key", "value")
		_ = publisher.Subscribe("key")
	}
}
