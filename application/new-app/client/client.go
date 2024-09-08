package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
