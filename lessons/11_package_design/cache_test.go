package packagedesign

import "testing"

func TestCacheStoresValuesBehindSmallAPI(t *testing.T) {
	cache := NewCache()
	cache.Set("language", "go")

	got, ok := cache.Get("language")
	if !ok || got != "go" {
		t.Fatalf("Get() = %q, %v; want %q, true", got, ok, "go")
	}
}
