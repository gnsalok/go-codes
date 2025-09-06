package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
}

func main() {

	// Used to Encode Go Object into JSON
	data, _ := json.Marshal(&employee{1, "Alok", "Tripathi", 24})
	fmt.Println(string(data))

}
