package myapi_test

import (
	"context"
	"testing"

	"httpxclient/mocks"
	"httpxclient/pkg/myapi"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGreeter(t *testing.T) {
	mockAPI := mocks.NewAPI(t)
	mockAPI.
		EXPECT().
		GetUser(mock.Anything, "42").
		Return(myapi.User{ID: "42", Name: "Ada"}, nil)

	msg, err := myapi.Greeter(context.Background(), mockAPI, "42")
	require.NoError(t, err)
	require.Equal(t, "Hello, Ada", msg)
}
