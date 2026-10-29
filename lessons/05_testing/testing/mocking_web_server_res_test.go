// All material is licensed under the Apache License Version 2.0, January 2004.
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss><channel>
<title>Going Go Programming</title>
<description>Go articles</description>
<link>https://example.test/</link>
<item>
<title>Object-Oriented Programming Mechanics</title>
<description>Composition in Go.</description>
<link>https://example.test/composition</link>
</item>
</channel></rss>`

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

func feedHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, feed)
}

func TestDecodeFeedResponse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://example.test/feed", nil)
	recorder := httptest.NewRecorder()
	feedHandler(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}

	var document Document
	if err := xml.NewDecoder(response.Body).Decode(&document); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got := len(document.Channel.Items); got != 1 {
		t.Fatalf("item count = %d, want 1", got)
	}
}
