package subtest

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newStatusClient() *http.Client {
	return &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			status := http.StatusNotFound
			if req.URL.Path == "/feed" {
				status = http.StatusOK
			}

			return &http.Response{
				StatusCode: status,
				Body:       io.NopCloser(strings.NewReader("test response")),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}
}

func TestDownload(t *testing.T) {
	client := newStatusClient()
	tests := []struct {
		name       string
		url        string
		wantStatus int
	}{
		{name: "existing feed", url: "https://example.test/feed", wantStatus: http.StatusOK},
		{name: "missing feed", url: "https://example.test/missing", wantStatus: http.StatusNotFound},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := client.Get(test.url)
			if err != nil {
				t.Fatalf("GET %q: %v", test.url, err)
			}
			defer response.Body.Close()

			if response.StatusCode != test.wantStatus {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.wantStatus)
			}
		})
	}
}

func TestDownloadInParallel(t *testing.T) {
	client := newStatusClient()
	tests := []struct {
		name       string
		url        string
		wantStatus int
	}{
		{name: "existing feed", url: "https://example.test/feed", wantStatus: http.StatusOK},
		{name: "missing feed", url: "https://example.test/missing", wantStatus: http.StatusNotFound},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			response, err := client.Get(test.url)
			if err != nil {
				t.Fatalf("GET %q: %v", test.url, err)
			}
			defer response.Body.Close()

			if response.StatusCode != test.wantStatus {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.wantStatus)
			}
		})
	}
}
