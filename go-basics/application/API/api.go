// Fetch data from API

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number  int                 `json:"number"`
	Message string              `json:"message" json:"status"`
	People  []map[string]string `json:"people"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(people1.Number, people1.Message, people1.People[0]["craft"])
	fmt.Println("----------------")
	fmt.Println(people1.Number, people1.Message, people1.People[0]["craft"])

	var p *people

	// convert json into object
	json.Unmarshal([]byte(`{"number":1, "status":"Success", "people":[{"craft": "ISS", "name": "Sergey Prokopyev"}]}`), &p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----Unmarshalling-----")
	fmt.Println(p.Number, p.Message, p.People)
}
