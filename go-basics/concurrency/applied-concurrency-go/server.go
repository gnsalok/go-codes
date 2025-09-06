package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm up and running...")
}

func main() {
	// configure path and handler function
	http.HandleFunc("/hello", Hello)

	// Listen on PORT and block main
	fmt.Println("Server in running...")
	http.ListenAndServe(":8080", nil)
}
