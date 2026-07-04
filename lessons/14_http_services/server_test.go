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
