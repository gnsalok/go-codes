package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server runnig on Port 8080")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
