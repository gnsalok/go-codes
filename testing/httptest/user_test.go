package httptest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUserHandler(t *testing.T) {

	router := gin.Default()

	testCases := []struct {
		name           string
		payload        []byte
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Input",
			payload:        []byte(`{"id": "1", "username": "john_doe", "email": "john@example.com"}`),
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":"1","username":"john_doe","email":"john@example.com"}`,
		},
		{
			name:           "Invalid JSON",
			payload:        []byte(`invalid json`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid character 'i' looking for beginning of value"}`,
		},
	}

	for _, tc := range testCases {

		req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(tc.payload))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Check response status code
		assert.Equal(t, tc.expectedStatus, w.Code, "Test case: "+tc.name)

		// Check response body
		assert.JSONEq(t, tc.expectedBody, w.Body.String(), "Test case: "+tc.name)

	}
}
