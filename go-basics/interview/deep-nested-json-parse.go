package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	Zip    string `json:"zip"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	City    string  `json:"city"`
	Address Address `json:"address"`
}

func main() {

	// write a nested json
	jsonData := `{"name": "John", "age": 30, "city": "New York", "address": {"street": "123 Main St", "zip": "10001"}}`
	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		panic(err)
	}
	// Output the person struct
	fmt.Printf("Name: %s, Age: %d, City: %s, Street: %s, Zip: %s\n", person.Name, person.Age, person.City, person.Address.Street, person.Address.Zip)

}
