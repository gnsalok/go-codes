package httpservices

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateGreeting(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/greetings", strings.NewReader(`{"name":"Gopher"}`))
	res := httptest.NewRecorder()

	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Fatalf("status = %d; want %d", res.Code, http.StatusCreated)
	}
	if body := res.Body.String(); !strings.Contains(body, "hello, Gopher") {
		t.Fatalf("body = %q; want greeting", body)
	}
}

func TestCreateGreetingRejectsInvalidRequests(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		body       string
		wantStatus int
	}{
		{name: "wrong method", method: http.MethodGet, wantStatus: http.StatusMethodNotAllowed},
		{name: "invalid JSON", method: http.MethodPost, body: `{`, wantStatus: http.StatusBadRequest},
		{name: "blank name", method: http.MethodPost, body: `{"name":"  "}`, wantStatus: http.StatusBadRequest},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.method, "/greetings", strings.NewReader(test.body))
			response := httptest.NewRecorder()

			NewHandler().ServeHTTP(response, request)

			if response.Code != test.wantStatus {
				t.Fatalf("status = %d; want %d", response.Code, test.wantStatus)
			}
		})
	}
}
