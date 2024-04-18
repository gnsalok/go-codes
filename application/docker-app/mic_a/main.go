package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Microservice A!")
	})

	fmt.Println("Microservice A listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
