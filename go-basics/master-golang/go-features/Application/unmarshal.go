package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"firstName"`
	Last  string `json:"lastName"`
	Age   int    `json:"age"`
}

/*
Unmarshalling in GO
-----------------------------------------------------
Line 1: Creating json string into byte code.
Line 2: Create empty Response struct and assign res variable.
Line 3: Unmarshal by passing a pointer to an empty structs.
Line 3: Print the Struct Name value.
*/

func main() {
	s := `[{"firstName":null,"lastName":"Bond","age":32},{"firstName":"Miss","lastName":"Moneypenny","age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
