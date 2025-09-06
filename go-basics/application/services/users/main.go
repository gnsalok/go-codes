package main

import (
	"flag"
	"net/http"

	"github.com/gnsalok/ps-go/application/services/users/api"
)

func main() {
	listenAdd := flag.String("listenaddr", ":4999", "todo")
	flag.Parse()

	http.HandleFunc("/", api.HandlerLandingPage)
	http.HandleFunc("/user", api.HandleGetUser)
	http.HandleFunc("/account", api.HandleGetAccount)

	http.ListenAndServe(*listenAdd, nil)
}
