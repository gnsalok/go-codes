package httpx

import "net/http"

// Doer is a tiny boundary you can mock if you want to unit-test client internals.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
