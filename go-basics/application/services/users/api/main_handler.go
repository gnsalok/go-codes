package api

import (
	"fmt"
	"net/http"
)

func HandlerLandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world service is running...")
}
