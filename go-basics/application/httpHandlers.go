package main

import (
	"net/http"
)

type fooHandler struct {
	Message string
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello bar"))
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))

}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Hello foo..."})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
