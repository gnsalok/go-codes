package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(".........Hello world Service.......")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world service is running...")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
