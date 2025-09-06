package main

import (
	"encoding/json"
	"fmt"
)

type employee_struct struct {
	Number int `json:"number"`
}

func main() {
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &employee_struct)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)
}
