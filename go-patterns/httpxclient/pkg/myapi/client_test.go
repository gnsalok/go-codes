package myapi

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClient_GetUser_Success(t *testing.T) {
	m := mocks.NewDoer(t)

	// Expect a GET /users/42 and return a fake 200 JSON body
	m.EXPECT().
		Do(mock.MatchedBy(func(r *http.Request) bool {
			return r.Method == http.MethodGet && r.URL.Path == "/users/42"
		})).
		Return(&http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{"id":"42","name":"Ada"}`)),
			Header:     make(http.Header),
		}, nil)

	c := NewClient("https://example.test", WithHTTPClient(m))
	got, err := c.GetUser(context.Background(), "42")
	require.NoError(t, err)
	require.Equal(t, "Ada", got.Name)
}

func TestClient_CreateUser_Error(t *testing.T) {
	m := mocks.NewDoer(t)

	// Assert HTTP method and path, return 400
	m.EXPECT().
		Do(mock.MatchedBy(func(r *http.Request) bool {
			return r.Method == http.MethodPost && r.URL.Path == "/users"
		})).
		Return(&http.Response{
			StatusCode: 400,
			Body:       io.NopCloser(strings.NewReader(`{"error":"bad input"}`)),
			Header:     make(http.Header),
		}, nil)

	c := NewClient("https://example.test", WithHTTPClient(m))
	_, err := c.CreateUser(context.Background(), CreateUserRequest{Name: ""})
	require.Error(t, err)
	require.Contains(t, err.Error(), "create user: 400")
}
