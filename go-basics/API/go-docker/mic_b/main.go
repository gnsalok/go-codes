package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://mic_a:8080/hello")
	if err != nil {
		log.Fatalf("Error calling Microservice A: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response from Microservice A: %v", err)
	}

	fmt.Println(string(body))
}
