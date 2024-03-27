package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	Id        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
}

func main() {
	var emp *Employee

	err := json.Unmarshal([]byte(`{"id":1, "firstName":"Alok", "lastName":"Tripathi", "age":24 }`), &emp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(emp.Id, emp.FirstName, emp.LastName, emp.Age)
}
